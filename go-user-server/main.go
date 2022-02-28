package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"

	"apulse.ai/tzuchi-upmp/server/handler"
	"apulse.ai/tzuchi-upmp/server/scheduler"
	"apulse.ai/tzuchi-upmp/server/service"
	"apulse.ai/tzuchi-upmp/server/service/requester"
	"apulse.ai/tzuchi-upmp/server/store"
	"apulse.ai/tzuchi-upmp/server/utils/validator"

	_ "apulse.ai/tzuchi-upmp/server/docs"
)

const (
	Version = "0.1.0-alpha"

	baseURLOfAssets  = "/assets"
	baseURLOfAPI     = "/api"
	baseURLOfSwagger = "/swagger"
)

func init() {
	// 讀取寫在.env檔案中的環境變數
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf("failed to load .env file, reason: %v", err))
	}
	requester.Init()
}

// setupMiddleware 設置中介軟體，包含紀錄器、網址預處理以及跨域設定
func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(context echo.Context) bool {
			return !strings.HasPrefix(context.Path(), baseURLOfAPI+"/")
		},
	}))

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderAcceptEncoding,
			echo.HeaderAuthorization,
			echo.HeaderContentType,
			echo.HeaderOrigin,
		},
		AllowOrigins: []string{"http://localhost:3000"},
	}))
}

func main() {
	isInDevelopment, _ := strconv.ParseBool(os.Getenv("DEVELOPMENT"))

	// 設置伺服器框架echo、中介軟體和驗證器，如果是在開發模式則設置Swagger文件
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	if isInDevelopment {
		e.GET(baseURLOfSwagger+"/*", echoSwagger.WrapHandler)
	}
	setupMiddleware(e)
	e.Validator = validator.NewValidator()

	// 設置資料庫儲存、處理器，並讓處理器中的路由註冊在echo的/api中
	if st, err := store.NewStore(); err != nil {
		e.Logger.Fatal(err)
	} else {
		sv := service.NewService(st)

		if err := sv.HDSS.SetAccessToken(); err != nil {
			e.Logger.Fatal(err)
		}
		if err := sv.HDSS.SyncDataOfUser(); err != nil {
			e.Logger.Fatal(err)
		}

		h := handler.NewHandler(st, sv) // api
		h.RegisterAPI(e.Group(baseURLOfAPI), st.Enforcer)
		h.RegisterAssets(e.Group(baseURLOfAssets))

		sd := scheduler.NewScheduler(sv, e.Logger) // run funcs in an interval background
		sd.SetupJobs()
		sd.Start()
	}

	// 啟動伺服器
	if isInDevelopment {
		e.Logger.Fatal(e.Start(":8000"))
	} else {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "../web/dist",
			Index: "index.html",
			HTML5: true,
		}))
		certFile := os.Getenv("TLS_CERT_FILE_PATH")
		keyFile := os.Getenv("TLS_KEY_FILE_PATH")
		if certFile != "" && keyFile != "" {
			e.Logger.Fatal(e.StartTLS(":443", certFile, keyFile))
		} else {
			e.Logger.Fatal(e.Start(":80"))
		}
	}
}
