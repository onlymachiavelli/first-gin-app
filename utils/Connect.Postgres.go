

package utils
import (
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

func Connect () (*gorm.DB) {
	dsn := "postgres:root@tcp(localhost:4269)/testgo?charset=utf8mb4&parseTime=True&loc=Local" 
	db, err  := gorm.Open(postgres.Open(dsn) , &gorm.Config{}) 
	if err != nil {
		return nil
	}
	return db
}