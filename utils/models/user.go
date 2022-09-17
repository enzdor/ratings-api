package models

type User struct {
	User_id		int 	`gorm:"primaryKey;unique;notNull" json:"user_id"`
	Email		string 	`json:"email" gorm:"unique;notNull;size:100"`
	Password	string	`json:"password" gorm:"notNull; size:255"`
}
