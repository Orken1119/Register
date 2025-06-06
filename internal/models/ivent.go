package models

import (
	"context"
)

type Ivent struct {
	ID                    int    `json:"id"`
	Name                  string `json:"ivent_name"`
	Information           string `json:"information"`
	Organization          string `json:"organization"`
	PosterUrl             string `json:"poster_url"`
	PreviewUrl            string `json:"preview_url"`
	SkillsDirection       string `json:"skill_direction"`
	Address               string `json:"address"`
	Time                  string `json:"time"`
	CountOfPeople         int    `json:"people_count"`
	HowManyPeopleAccepted int    `json:"members_count"`
	EstimatedWorkHours    int    `json:"estimated_work_hours"`
}

type MainIvent struct {
	Name                  string `json:"ivent_name"`
	Information           string `json:"information"`
	Organization          string `json:"organization"`
	PosterUrl             string `json:"poster_url"`
	PreviewUrl            string `json:"preview_url"`
	SkillsDirection       string `json:"skill_direction"`
	Address               string `json:"address"`
	Time                  string `json:"time"`
	CountOfPeople         int    `json:"people_count"`
	HowManyPeopleAccepted int    `json:"members_count"`
	EstimatedWorkHours    int    `json:"estimated_work_hours"`
}

type IventRepository interface {
	CreateIvent(c context.Context, ivent *MainIvent) (*Ivent, error)
	DeleteIvent(c context.Context, id int) error
	UpdateIvent(c context.Context, ivent *Ivent) error
	GetAllIvent(c context.Context) (*[]Ivent, error)
	GetIventById(c context.Context, id int) (*Ivent, error)
}
