

import (
	"fmt" 
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

func Connect () {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" 
	db, err  := gorm.Open(postgres.Open(dsn) , &gorm.Config{}) 
}