package user

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"apulse.ai/tzuchi-upmp/server/model"
)

type Store struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
}

func NewStore(db *gorm.DB, enforcer *casbin.Enforcer) *Store {
	return &Store{
		db:       db.Model(new(model.User)).Session(&gorm.Session{}),
		enforcer: enforcer,
	}
}

func (s *Store) Init() error {
	if err := s.addDefaultAdministratorModel(); err != nil {
		return err
	}
	return nil
}

func (s *Store) addDefaultAdministratorModel() error {
	datum := model.User{
		Username: "admin",
		IDNumber: "Y800000004", // FIXME: 暫時措施
	}
	var count int64
	if err := s.db.Where(&datum).Count(&count).Error; err != nil {
		return err
	} else if count > 0 {
		return nil
	}
	datum.FullName = "系統管理員"
	if err := s.db.Create(&datum).Error; err != nil {
		return err
	}
	roles := []string{"admin"}
	if _, err := s.enforcer.AddRolesForUser(datum.ID.String(), roles); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetData() ([]model.User, error) {
	data := new([]model.User)
	if err := s.db.Find(data).Error; err != nil {
		return nil, err
	}
	return *data, nil
}

func (s *Store) AddData(data []model.User) error {
	return s.db.Create(&data).Error
}

func (s *Store) GetDatumByID(id uuid.UUID) (*model.User, error) {
	datum := new(model.User)
	tx := s.db.Omit("password")
	tx = tx.Where("id = ?", id)
	if err := tx.First(datum).Error; err != nil {
		return nil, err
	}
	return datum, nil
}

func (s *Store) UpdateDatumByID(id uuid.UUID, patch map[string]interface{}) error {
	data := make([]model.User, 0, 1)
	tx := s.db.Model(&data).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}})
	tx = tx.Where("id = ?", id)
	if err := tx.Updates(patch).Error; err != nil {
		return err
	} else if len(data) == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *Store) DeleteDatumByID(id uuid.UUID) error {
	data := make([]model.User, 0, 1)
	tx := s.db.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}})
	tx = tx.Where("id = ?", id)
	if err := tx.Delete(&data).Error; err != nil {
		return err
	} else if len(data) == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *Store) GetPasswordOfDatumByID(id uuid.UUID) ([]byte, error) {
	datum := new(model.User)
	tx := s.db.Select("password")
	tx = tx.Where("id = ?", id)
	if err := tx.First(datum).Error; err != nil {
		return nil, err
	}
	return datum.Password, nil
}

func (s *Store) GetDatumByUsername(username string) (*model.User, error) {
	datum := new(model.User)
	tx := s.db.Omit("password")
	tx = tx.Where("username = ?", username)
	if err := tx.First(datum).Error; err != nil {
		return nil, err
	}
	return datum, nil
}

func (s *Store) SyncData(data []model.User, withPassword bool) error {
	if len(data) == 0 {
		return nil
	}
	columns := []string{"username", "full_name", "email", "role"}
	if withPassword {
		columns = append(columns, "password")
	}
	return s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id_number"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Create(&data).Error
}
