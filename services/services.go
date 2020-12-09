package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DBConnection This service returns the connection to the db
func DBConnection(dataSource string) (*gorm.DB, error) {

	gormDB, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
