package model

import (
	"github.com/damingerdai/blog-service/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state`
}

func (tag Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	*Model
	List  []*Tag
	Pager *app.Pager
}
