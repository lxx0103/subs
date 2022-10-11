package subscription

import "time"

type Subscription struct {
	ID                  int64     `db:"id" json:"id"`
	OpenID              string    `db:"open_id" json:"open_id"`
	CatID               int       `db:"cat_id" json:"cat_id"`
	SubCatID            int       `db:"sub_cat_id" json:"sub_cat_id"`
	SubTypeID           int       `db:"sub_type_id" json:"sub_type_id"`
	SourceID            int       `db:"source_id" json:"source_id"`
	Driver              string    `db:"driver" json:"driver"`
	DriverContactTypeID int       `db:"driver_contact_type_id" json:"driver_contact_type_id"`
	DriverContact       string    `db:"driver_contact" json:"driver_contact"`
	PayTypeID           int       `db:"pay_type_id" json:"pay_type_id"`
	PaymentAmount       float64   `db:"payment_amount" json:"payment_amount"`
	SubStartDate        string    `db:"sub_start_date" json:"sub_start_date"`
	SubEndDate          string    `db:"sub_end_date" json:"sub_end_date"`
	NeedAlert           int       `db:"need_alert" json:"need_alert"`
	AlertBefore         int       `db:"alert_before" json:"alert_before"`
	Status              int       `db:"status" json:"status"`
	Created             time.Time `db:"created" json:"created"`
	CreatedBy           string    `db:"created_by" json:"created_by"`
	Updated             time.Time `db:"updated" json:"updated"`
	UpdatedBy           string    `db:"updated_by" json:"updated_by"`
}
