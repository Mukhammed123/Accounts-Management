package requester

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"

	"apulse.ai/tzuchi-upmp/server/utils/slices"
)

func MakeRequestToHDSS(method string, url string, requestData datatypes.JSONMap,
	accessToken *string, queries map[string]string, allowedStatusCode []int) (interface{}, error) {
	var requestBody io.Reader
	if requestData != nil {
		if marshaledRequestData, err := json.Marshal(requestData); err != nil {
			panic(err)
		} else {
			requestBody = bytes.NewBuffer(marshaledRequestData)
		}
	}
	if request, err := http.NewRequest(method, apiBaseURLOfHDSS+url, requestBody); err != nil {
		panic(err)
	} else {
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Access-Control-Allow-Origin", "*")
		// request.Header.Set("Content-Type", "application/json")
		if accessToken != nil {
			request.Header.Add("Authorization", "Bearer "+*accessToken)
		}
		if len(queries) > 0 {
			queryOfRequest := request.URL.Query()
			for key, value := range queries {
				queryOfRequest.Add(key, value)
			}
			request.URL.RawQuery = queryOfRequest.Encode()
		}
		var responseData interface{}
		if response, err := http.DefaultClient.Do(request); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else if !slices.Contain(allowedStatusCode, response.StatusCode) {
			defer response.Body.Close()
			if responseMap, ok := responseData.(map[string]interface{}); ok {
				if errorMessage, ok := responseMap["error_message"]; ok {
					return nil, echo.NewHTTPError(response.StatusCode, errorMessage)
				}
			}
			if marshaledResponseData, err := json.Marshal(responseData); err != nil {
				panic(err)
			} else {
				return nil, echo.NewHTTPError(response.StatusCode, string(marshaledResponseData))
			}
		}
		return responseData, nil
	}
}
