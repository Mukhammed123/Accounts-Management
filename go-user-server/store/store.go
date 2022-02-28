package store

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"apulse.ai/tzuchi-upmp/server/store/user"
)

type Store struct {
	User     *user.Store
	Enforcer *casbin.Enforcer
}

func setupDB() (*gorm.DB, error) {
	if db, err := openDB(); err != nil {
		return nil, err
	} else if err := migrateDB(db); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func setupEnforcer(db *gorm.DB) (*casbin.Enforcer, error) {
	if adapter, err := gormadapter.NewAdapterByDB(db); err != nil {
		return nil, err
	} else if enforcer, err := casbin.NewEnforcer("model/auth.conf", adapter); err != nil {
		return nil, err
	} else {
		return enforcer, nil
	}
}

// NewStore 設置資料庫、權限管理器以及操作資料庫所需的儲存框架，並讓儲存框架對資料庫進行預設資料的初始化
func NewStore() (*Store, error) {
	if db, err := setupDB(); err != nil {
		return nil, err
	} else if enforcer, err := setupEnforcer(db); err != nil {
		return nil, err
	} else {
		store := &Store{
			User:     user.NewStore(db, enforcer),
			Enforcer: enforcer,
		}
		if err := store.User.Init(); err != nil {
			return nil, err
		}
		return store, nil
	}
}
