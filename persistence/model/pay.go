package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type Pay struct {
	Uuid           uuid.UUID      `gorm:"primary_key;type:uuid;"`
	CreatedAt      datatypes.Date `gorm:"type:timestamp"`
	PayRequestJson postgres.Jsonb `gorm:"type:jsonb;column:pay_request;"`
	PayRequest     PayRequest     `gorm:"-"`
}

func (a *Pay) MarshallPayRequest() error {
	var err error
	a.PayRequestJson.RawMessage, err = json.Marshal(a.PayRequest)

	return err
}

func (a *Pay) UnmarshalPayRequest() error {
	return json.Unmarshal(a.PayRequestJson.RawMessage, &a.PayRequest)
}
