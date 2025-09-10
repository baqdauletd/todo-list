package models

import "time"

type Task struct{
	ID 				uint				`gorm:"primaryKey;autoIncrement"`
	Title       	string     			`gorm:"not null"`
	Description   	string				`gorm:"type:text"`
	Completed 		bool				`gorm:"default:false"`
	DueDate			*time.Time			`gorm:"type:timestamp"`
	CreateAT  		time.Time			`gorm:"autoCreateTime"`
	Priority 		string				`gorm:"not null"`
}




