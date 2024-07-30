// // models/user.go
// package models

// import "time"

// type User struct {
// 	Id              int       `json:"id"`
// 	Username        string    `json:"username"`
// 	Password        string    `json:"password"`
// 	Firstname       string    `json:"firstname"`
// 	Lastname        string    `json:"lastname"`
// 	Confirmpassword string    `json:"confirmpassword"`
// 	LoginAttempts   int       `json:"loginattempts" gorm:"default:0"`
// 	BlockedUntil    time.Time `json:"blocked_until"`
// }

// type ViewUser struct {
// 	Id              int       `json:"id"`
// 	Username        string    `json:"username"`
// 	Firstname       string    `json:"firstname"`
// 	Lastname        string    `json:"lastname"`
// 	Password        string    `json:"password"`
// 	Confirmpassword string    `json:"confirmpassword"`
// 	DeletedAt       time.Time `gorm:"index" json:"deletedAt,omitempty"`
// }
