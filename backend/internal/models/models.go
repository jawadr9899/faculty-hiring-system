package models

import (
	"time"
	"uhs/internal/types/common"
)

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Id          string        `gorm:"primaryKey;" json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Role        common.Role   `json:"-"`
	Password    string        `json:"password"`
	Analytics   Analytics     `gorm:"foreignKey:UserId;references:Id;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"-"`
	Application []Application `gorm:"foriegnKey:UserId;references:Id;" json:"-"`
	Cv          Cv            `gorm:"foriegnKey:UserId;references:Id;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"-"`
	Timestamps
}

type Cv struct {
	Id       string `gorm:"primaryKey" json:"id"`
	UserId   string `gorm:"uniqueIndex;not null;" json:"userId"`
	FileName string `json:"fileName"`
	FileUrl  string `json:"fileUrl"`
	Timestamps
}

type Analytics struct {
	Id               string `gorm:"primaryKey;" json:"id"`
	UserId           string `gorm:"not null;"    json:"useId"`
	PositionId       string `gorm:"uniqueIndex;not null;"   json:"positionId"`
	Dept             string `json:"dept"`
	DegreeLevel      string `json:"degreeLevel"`
	PublicationCount uint   `json:"publicationCount"`
	Experience       uint   `json:"experience"`
	AcademicScore    uint   `json:"academicScore"`
	ResearchScore    uint   `json:"researchScore"`
	TeachingScore    uint   `json:"teachingScore"`
	IndustrialScore  uint   `json:"industrialScore"`
	SalaryScore      uint   `json:"salaryScore"`
	AdminScore       uint   `json:"adminScore"`
	CompositeRank    uint   `json:"compositeRank"`
	Summary          string `json:"summary"`

	Timestamps
}

type Position struct {
	Id          string        `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"not null;" json:"name"`
	Description string        `gorm:"not null;" json:"description"`
	Application []Application `gorm:"foriegnKey:PositionId;references:Id;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"-"`
	Analytics   []Analytics   `gorm:"foriegnKey:PositionId;references:Id;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"-"`

	Timestamps
}

type Application struct {
	Id         string `gorm:"primarykey;"`
	UserId     string `gorm:"not null;"`
	PositionId string `gorm:"uniqueIndex;not null"`
	Timestamps
}
