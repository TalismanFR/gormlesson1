package model

import uuid "github.com/satori/go.uuid"

type PayRequest struct {
	InvoiceId   uuid.UUID
	Amount      int64
	Description string `gorm:"size:2000"`
}
