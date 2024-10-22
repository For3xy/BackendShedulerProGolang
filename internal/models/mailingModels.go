package models

type Mailing struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	TypeMailing    int    `json:"type" gorm:"not null"`
	Filter         string `json:"filter" gorm:"not null"`
	GroupFlag      bool   `json:"group_flag" gorm:"not null"`
	SingleFileFlag bool   `json:"single_file_flag" gorm:"not null"`
	Address        string `json:"address" gorm:"not null"`
	ChatId         string `json:"chat_id" gorm:"not null"`
}

type CreateMailing struct {
	TypeMailing    int    `json:"type" gorm:"not null"`
	Filter         string `json:"filter" gorm:"not null"`
	GroupFlag      bool   `json:"group_flag" gorm:"not null"`
	SingleFileFlag bool   `json:"single_file_flag" gorm:"not null"`
	Address        string `json:"address" gorm:"not null"`
	ChatId         string `json:"chat_id" gorm:"not null"`
}

type UpdateMailing struct {
	TypeMailing    int    `json:"type" gorm:"not null"`
	Filter         string `json:"filter" gorm:"not null"`
	GroupFlag      bool   `json:"group_flag" gorm:"not null"`
	SingleFileFlag bool   `json:"single_file_flag" gorm:"not null"`
	Address        string `json:"address" gorm:"not null"`
	ChatId         string `json:"chat_id" gorm:"not null"`
}
