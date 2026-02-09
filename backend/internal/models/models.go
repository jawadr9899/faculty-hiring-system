package models

import "time"

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Id          string        `gorm:"primaryKey;" json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
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
	UserId           string `gorm:"not null"`
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

	Position []Position `gorm:"many2many:analytics_positions;" json:"-"`

	Timestamps
}

type Position struct {
	Id          string        `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"not null;" json:"name"`
	Description string        `gorm:"not null;" json:"description"`
	Application []Application `gorm:"foriegnKey:PositionId;references:Id;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"-"`

	Timestamps
}

type Application struct {
	Id         string `gorm:"primarykey;"`
	UserId     string `gorm:"not null;"`
	PositionId string `gorm:"not null"`
	Timestamps
}
