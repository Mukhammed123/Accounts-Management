package store

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"apulse.ai/tzuchi-upmp/server/model"
)

/* delete all tables:
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;
*/

const createTypes = `
do $$ begin
create extension if not exists "uuid-ossp";
if not exists(select * from pg_type where typname = 'sex') then
	create type sex as enum('male', 'female');
end if;
if not exists(select * from pg_type where typname = 'weekday') then
	create type weekday as enum('sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday');
end if;
end $$
`

func openDB() (*gorm.DB, error) {
	// Reference: https://www.postgresql.org/docs/10/libpq-connect.html#LIBPQ-CONNSTRING
	url := os.Getenv("DATABASE_URL")
	if db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		return nil, err
	} else if err := db.Exec(createTypes).Error; err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		new(model.User),
	)
}
