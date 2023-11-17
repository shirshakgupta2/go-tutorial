package constants

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB
var Err error

const CLOUD_SQL_CONNECTION_URI = "postgresql://postgres:12345678@15.207.118.139:5432/postgres"

func InitDataBase() {

	DataBase, _ = gorm.Open(postgres.Open(CLOUD_SQL_CONNECTION_URI))
	
}