package models

type Rating struct {
	Rating_id	int `json:"rating_id" gorm:"primaryKey;unique;notNull"`
	Name		string `json:"name" gorm:"notNull;size:255"`
	Entry_type	string `json:"entry_type" gorm:"notNull;size:255"`
	Rating		int `json:"rating" gorm:"notNull"`
	Consumed	bool `json:"consumed" gorm:"notNull"`
	User_id		int `json:"user_id" gorm:"notNull;foreignKey:user_id;references:user_id;constraint:OnUpdate,OnDelete"`
}

type SearchRating struct {
	Rating_id	int `json:"rating_id" gorm:"primaryKey;unique;notNull"`
	Name		string `json:"name" gorm:"notNull;size:255"`
	Entry_type	string `json:"entry_type" gorm:"notNull;size:255"`
	Rating		int `json:"rating" gorm:"notNull"`
	Consumed	int `json:"consumed" gorm:"notNull"`
	User_id		int `json:"user_id" gorm:"notNull;foreignKey:user_id;references:user_id;constraint:OnUpdate,OnDelete"`
}

