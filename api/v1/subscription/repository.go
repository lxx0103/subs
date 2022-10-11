package subscription

import (
	"database/sql"
	"time"
)

type subscriptionRepository struct {
	tx *sql.Tx
}

func NewSubscriptionRepository(transaction *sql.Tx) *subscriptionRepository {
	return &subscriptionRepository{
		tx: transaction,
	}
}

func (r *subscriptionRepository) CreateSubscription(info Subscription) error {
	_, err := r.tx.Exec(`
		INSERT INTO s_subscriptions
		(
			open_id,
			cat_id,
			sub_cat_id,
			sub_type_id,
			source_id,
			driver,
			driver_contact_type_id,
			driver_contact,
			pay_type_id,
			payment_amount,
			sub_start_date,
			sub_end_date,
			need_alert,
			alert_before,
			status,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, info.OpenID, info.CatID, info.SubCatID, info.SubTypeID, info.SourceID, info.Driver, info.DriverContactTypeID, info.DriverContact, info.PayTypeID, info.PaymentAmount, info.SubStartDate, info.SubEndDate, info.NeedAlert, info.Status, info.Created, info.CreatedBy, info.Updated, info.UpdatedBy)
	return err
}

func (r *subscriptionRepository) GetSubscriptionByID(id int64, openID string) (*Subscription, error) {
	var res Subscription
	row := r.tx.QueryRow(`
	SELECT 
	id,
	open_id,
	cat_id, 
	sub_cat_id, 
	sub_type_id,
	source_id,
	driver,
	driver_contact_type_id,
	driver_contact,
	pay_type_id,
	payment_amount,
	sub_start_date,
	sub_end_date,
	need_alert,
	alert_before
	FROM s_subscriptions 
	WHERE id = ? AND open_id = ? AND s.status > 0 LIMIT 1`, id, openID)
	err := row.Scan(&res.ID, &res.OpenID, &res.CatID, &res.SubCatID, &res.SubTypeID, &res.SourceID, &res.Driver, &res.DriverContactTypeID, &res.DriverContact, &res.PayTypeID, &res.PaymentAmount, &res.SubStartDate, &res.SubEndDate, &res.NeedAlert, &res.AlertBefore)
	return &res, err
}

func (r *subscriptionRepository) UpdateSubscription(id int64, info Subscription) error {
	_, err := r.tx.Exec(`
		UPDATE s_subscriptions SET		
		cat_id = ?,
		sub_cat_id = ?,
		sub_type_id = ?,
		source_id = ?,
		driver = ?,
		driver_contact_type_id = ?,
		driver_contact = ?,
		pay_type_id = ?,
		payment_amount = ?,
		sub_start_date = ?,
		sub_end_date = ?,
		need_alert = ?,
		alert_before = ?,
		status = ?,
		updated = ?,
		updated_by = ?
		WHERE id = ?
	`, info.CatID, info.SubCatID, info.SubTypeID, info.SourceID, info.Driver, info.DriverContactTypeID, info.DriverContact, info.PayTypeID, info.PaymentAmount, info.SubStartDate, info.SubEndDate, info.NeedAlert, 1, info.Updated, info.UpdatedBy, id)
	return err
}

func (r *subscriptionRepository) DeleteSubscription(id int64, byUser string) error {
	_, err := r.tx.Exec(`
		UPDATE s_subscriptions 
		SET status = -1,
		updated = ?,
		updated_by = ?,
		WHERE id = ?
	`, time.Now(), byUser, id)
	return err
}
