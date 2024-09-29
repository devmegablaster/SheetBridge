package models

import (
	"github.com/devmegablaster/SheetBridge/internal/constants"
	"github.com/google/uuid"
)

type Synk struct {
	Id            uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ConnectionId  uuid.UUID  `json:"connectionId" gorm:"type:uuid;not null" validate:"required"`
	Connection    Connection `json:"connection" gorm:"foreignKey:ConnectionId;references:Id" validate:"required"`
	UserId        uuid.UUID  `json:"userId" gorm:"type:uuid;not null" validate:"required"`
	User          User       `json:"user" gorm:"foreignKey:UserId;references:Id" validate:"required"`
	SpreadsheetId string     `json:"spreadsheetId" gorm:"type:varchar(255);not null" validate:"required"`
	SheetId       string     `json:"sheetId" gorm:"type:varchar(255);not null" validate:"required"`
	Table         string     `json:"table" gorm:"type:varchar(255);not null" validate:"required"`
	Status        string     `json:"status" gorm:"type:varchar(255);not null" validate:"required"`
	Message       string     `json:"message" gorm:"type:varchar(255);default:''"`
	SchemaId      uuid.UUID  `json:"schemaId" gorm:"type:uuid;not null"`
	Schema        Schema     `json:"schema" gorm:"foreignKey:SynkId;references:Id"`
}

type Schema struct {
	Id     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Col    []string  `json:"cols" gorm:"type:varchar(255)[]"`
	Type   []string  `json:"types" gorm:"type:varchar(255)[]"`
	SynkId uuid.UUID `json:"synkId" gorm:"type:uuid;not null" validate:"required"`
}

type SynkRequest struct {
	ConnectionId  string `json:"connectionId" validate:"required"`
	SpreadsheetId string `json:"spreadsheetId" validate:"required"`
	SheetId       string `json:"sheetId" validate:"required"`
	Table         string `json:"table" validate:"required"`
}

func (sr *SynkRequest) ToSynk(userId uuid.UUID) *Synk {
	return &Synk{
		ConnectionId:  uuid.MustParse(sr.ConnectionId),
		UserId:        userId,
		SpreadsheetId: sr.SpreadsheetId,
		SheetId:       sr.SheetId,
		Table:         sr.Table,
		Status:        constants.Synk.STATUS_INIT,
	}
}

type SynkResponse struct {
	Id            string `json:"id"`
	ConnectionId  string `json:"connectionId"`
	SpreadsheetId string `json:"spreadsheetId"`
	SheetId       string `json:"sheetId"`
	Table         string `json:"table"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

func (s *Synk) ToResponse() *SynkResponse {
	return &SynkResponse{
		Id:            s.Id.String(),
		ConnectionId:  s.ConnectionId.String(),
		SpreadsheetId: s.SpreadsheetId,
		SheetId:       s.SheetId,
		Table:         s.Table,
		Status:        s.Status,
		Message:       s.Message,
	}
}
