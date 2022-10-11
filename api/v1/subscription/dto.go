package subscription

import "subs/core/request"

type SubscriptionResponse struct {
	ID                    int64   `db:"id" json:"id"`
	CatID                 int     `db:"cat_id" json:"cat_id"`
	CatName               string  `db:"cat_name" json:"cat_name"`
	SubCatID              int     `db:"sub_cat_id" json:"sub_cat_id"`
	SubCatName            string  `db:"sub_cat_name" json:"sub_cat_name"`
	SubTypeID             int     `db:"sub_type_id" json:"sub_type_id"`
	SubTypeName           string  `db:"sub_type_name" json:"sub_type_name"`
	SourceID              int     `db:"source_id" json:"source_id"`
	SourceName            string  `db:"source_name" json:"source_name"`
	Driver                string  `db:"driver" json:"driver"`
	DriverContactTypeID   int     `db:"driver_contact_type_id" json:"driver_contact_type_id"`
	DriverContactTypeName string  `db:"driver_contact_type_name" json:"driver_contact_type_name"`
	DriverContact         string  `db:"driver_contact" json:"driver_contact"`
	PayTypeID             int     `db:"pay_type_id" json:"pay_type_id"`
	PayTypeName           string  `db:"pay_type_name" json:"pay_type_name"`
	PaymentAmount         float64 `db:"payment_amount" json:"payment_amount"`
	SubStartDate          string  `db:"sub_start_date" json:"sub_start_date"`
	SubEndDate            string  `db:"sub_end_date" json:"sub_end_date"`
	NeedAlert             int     `db:"need_alert" json:"need_alert"`
	AlertBefore           int     `db:"alert_before" json:"alert_before"`
}

type SubscriptionNew struct {
	CatID               int     `json:"cat_id" binding:"required"`
	SubCatID            int     `json:"sub_cat_id" binding:"required"`
	SubTypeID           int     `json:"sub_type_id" binding:"required"`
	SourceID            int     `json:"source_id" binding:"required"`
	Driver              string  `json:"driver" binding:"omitempty,max=128"`
	DriverContactTypeID int     `json:"driver_contact_type_id" binding:"omitempty"`
	DriverContact       string  `json:"driver_contact" binding:"omitempty,max=128"`
	PayTypeID           int     `json:"pay_type_id" binding:"omitempty,max=128"`
	PaymentAmount       float64 `json:"payment_amount" binding:"omitempty,max=128"`
	SubStartDate        string  `json:"sub_start_date" binding:"required,max=128"`
	SubEndDate          string  `json:"sub_end_date" binding:"required,max=128"`
	NeedAlert           int     `json:"need_alert" binding:"required,oneof=1 2"`
	AlertBefore         int     `json:"alert_before" binding:"omitempty,max=30"`
	OpenID              string  `json:"open_id" swaggerignore:"true"`
}

type SubscriptionFilter struct {
	CatID     int    `form:"cat_id" binding:"omitempty"`
	SubCatID  int    `form:"sub_cat_id" binding:"omitempty"`
	SubTypeID int    `form:"sub_type_id" binding:"omitempty"`
	SourceID  int    `form:"source_id" binding:"omitempty"`
	Driver    string `form:"driver" binding:"omitempty,max=128"`
	PayTypeID int    `form:"pay_type_id" binding:"omitempty,max=128"`
	OpenID    string `form:"open_id" swaggerignore:"true"`
	request.PageInfo
}

type SubscriptionUpdate struct {
	CatID               int     `json:"cat_id" binding:"omitempty"`
	SubCatID            int     `json:"sub_cat_id" binding:"omitempty"`
	SubTypeID           int     `json:"sub_type_id" binding:"omitempty"`
	SourceID            int     `json:"source_id" binding:"omitempty"`
	Driver              string  `json:"driver" binding:"omitempty,max=128"`
	DriverContactTypeID int     `json:"driver_contact_type_id" binding:"omitempty"`
	DriverContact       string  `json:"driver_contact" binding:"omitempty,max=128"`
	PayTypeID           int     `json:"pay_type_id" binding:"omitempty,max=128"`
	PaymentAmount       float64 `json:"payment_amount" binding:"omitempty,max=128"`
	SubStartDate        string  `json:"sub_start_date" binding:"omitempty,max=128"`
	SubEndDate          string  `json:"sub_end_date" binding:"omitempty,max=128"`
	NeedAlert           int     `json:"need_alert" binding:"omitempty,oneof=1 2"`
	AlertBefore         int     `json:"alert_before" binding:"omitempty,max=30"`
	OpenID              string  `json:"open_id" swaggerignore:"true"`
}

type SubscriptionID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
