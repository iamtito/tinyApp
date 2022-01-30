package database

import (
	"fmt"

	"github.com/iamtito/tinyApp/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	username = "root"
	password = "admin"
	hostname = "db"
	database = "test"
	// database = "go_basics"
)

func dsn(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, database)
}
func Connect() {
	// connect, err := gorm.Open(mysql.Open("root:admin@tcp(db:3306)/go_basics?parseTime=true"), &gorm.Config{})
	connect, err := gorm.Open(mysql.Open(dsn(database)), &gorm.Config{})

	if err != nil {
		panic("Database Connection Failed")
	}
	// TODO: Create a step to check if the database exist or not and try creating it
	// // check if db exists
	// stmt := fmt.Sprintf("use '%s';", database)
	// rs := connect.Raw(stmt)
	// if rs.Error != nil {
	// 	fmt.Sprintln("db not exist")
	// 	// return rs
	// }else{
	// 	cr := connect.Raw("CREATE DATABASE %s", database)
	// 	if cr.Error != nil {
	// 		fmt.Sprint("Error creating db")
	// 		// return cr.Error
	// 	}
	// }
	// defer Connect.Close()
	DB = connect

	connect.AutoMigrate(&users.User{}, &users.PasswordReset{}, &users.Role{}, &users.Permissions{}, &users.Product{}, &users.Order{}, &users.OrderItem{})
}
