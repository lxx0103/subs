package setting

import "subs/core/request"

type CatFilter struct {
	Name string `form:"name" binding:"omitempty,max=64,min=1"`
	request.PageInfo
}

type CatResponse struct {
	ID     string `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Status int    `db:"status" json:"status"`
}
