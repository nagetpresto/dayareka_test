package models

type User struct {
	ID       		int    		`json:"id" gorm:"primary_key:auto_increment"`
	Name     		string 		`json:"name" gorm:"type: varchar(255)"`
}

type UserProfileResponse struct {
	ID    		int    	`json:"id"`
	Name  		string 	`json:"name"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}