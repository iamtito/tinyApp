package users

type Permissions struct {
	// common.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}