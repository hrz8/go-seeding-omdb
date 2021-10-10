package models

import (
	"time"
)

type (
	// LogRequest represents LogRequest object for DB
	LogRequest struct {
		ID         uint      `gorm:"column:id;primaryKey" json:"id"`
		Type       string    `gorm:"column:type" json:"type" validate:"required"`
		RequestURI string    `gorm:"column:request_uri" json:"requestUri"`
		Header     string    `gorm:"column:header" json:"header"`
		Method     string    `gorm:"column:method" json:"method"`
		CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`
	}
)
