package config

import (
	"errors"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	return errors.New("PQP")
}

func GetLogger() *Logger {
	//Initialize Logger
	return NewLogger()
}
