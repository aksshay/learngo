package services

import (
	"crypto/md5"
	"encoding/hex"

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

//GenerateMD5 This service is used to generate MD5 of a string
func GenerateMD5(text string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", nil
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
