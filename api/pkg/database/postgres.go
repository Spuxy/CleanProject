package database

import (
	"fmt"

	"github.com/Spuxy/CleanArchitecture/pkg/secrets"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const secret_path string = "/api/pkg/secrets/secrets"

type Postgres struct {
	Db *gorm.DB
}

func NewPostgres() (*Postgres, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("We could not connect into database due to: %s", err)
	}
	return &Postgres{db}, nil
}

func Connect() (*gorm.DB, error) {
	ds, _ := secrets.NewDotSecret(secret_path)
	dbuser, _ := ds.Get("postgres_user")
	dbpass, _ := ds.Get("postgres_passwd")
	dbhost, _ := ds.Get("postgres_db")
	dbname, _ := ds.Get("postgres_name")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	return db, err
}
