package database

type User struct {
	ID       int    `gorm:"primary,default:autoincrement"`
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}
