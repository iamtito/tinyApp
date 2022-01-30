package users

import (
	"github.com/iamtito/tinyApp/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id"`
	common.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}
// [7.340ms] [rows:0] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`first_name`,`last_name`,`email`,`password`,`role_id`) VALUES ('2022-01-16 16:44:43.85','2022-01-16 16:44:43.85',NULL,'kabir','bolatito','tbola45@gmail.com','$2a$14$N3KHbvF.aV5h7TD7wc8MDuTHQ2Ui/Rr1m/qIHxFAVTQ6e7qTU7Boa',1)
type Role struct {
	// common.Model
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Permissions []Permissions `json:"permissions" gorm:"many2many:role_permissions"`
}

func (user *User) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	return users
}

type PasswordReset struct {
	ID    uint
	Email string
	Token string
}
