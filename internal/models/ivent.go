package models

import (
	"context"
)

type Ivent struct {
	ID              int    `json:"id"`
	Name            string `json:"ivent_name"`
	Information     string `json:"information"`
	Organization    string `json:"organization"`
	PosterUrl       string `json:"poster_url"`
	PreviewUrl      string `json:"preview_url"`
	SkillsDirection string `json:"skill_direction"`
}

type MainIvent struct {
	Name            string `json:"ivent_name"`
	Information     string `json:"information"`
	Organization    string `json:"organization"`
	PosterUrl       string `json:"poster_url"`
	PreviewUrl      string `json:"preview_url"`
	SkillsDirection string `json:"skill_direction"`
}

type IventRepository interface {
	CreateIvent(c context.Context, ivent *MainIvent) (*Ivent, error)
	DeleteIvent(c context.Context, id int) error
	UpdateIvent(c context.Context, ivent *Ivent) error
	GetAllIvent(c context.Context) (*[]Ivent, error)
	GetIventById(c context.Context, id int) (*Ivent, error)
}
