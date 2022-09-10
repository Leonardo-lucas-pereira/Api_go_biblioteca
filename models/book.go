package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Name        string    `json:"name"`
	Description string    `json:"Description"`
	MediumPrice float32   `json:"medium_price"`
	Author      string    `json:"author"`
	ImageURl    string    `json:"img_url"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeleteddAt  time.Time `json:"deleted" gorm:"index"`
}
