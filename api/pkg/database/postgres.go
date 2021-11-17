package database

import (
	"fmt"

	"github.com/Spuxy/CleanArchitecture/api/pkg/secrets"
	"github.com/jinzhu/gorm"
)

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
	ds, _ := secrets.NewDotSecret("/home/god/Programming/Go/CleanProject/api/pkg/secrets/secrets/")
	dbuser, _ := ds.Get("user")
	dbpass, _ := ds.Get("pass")
	dbhost, _ := ds.Get("host")
	dbname, _ := ds.Get("name")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open("postgres", dbURL)
	return db, err
}
