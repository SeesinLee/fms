package moddle

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type FaultList struct {
	ID	uuid.UUID	`json:"id" gorm:"type:char(36);primary_key"`
	FaultSketch	string	`json:"faultSketch" gorm:"type:longtext"`
	FaultLevel	string	`json:"faultLevel"`
	FaultType string	`json:"faultType"`
	Influence string	`json:"influence"`
	Duration	int	`json:"duration"`
	ActualTime string	`json:"actualTime"`
	Status string	`json:"status"`
	Creator	string	`json:"creator"`
	Cause   string    `json:"cause" gorm:"type:longtext"`
	Solution string	`json:"solution" gorm:"type:longtext"`
	CreatedAt Time `json:"createdAt" gorm:"type:timestamp"`
	StartAt time.Time `json:"startAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	EndedAt time.Time `json:"endAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (l *FaultList) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
