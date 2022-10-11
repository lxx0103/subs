package subscription

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type subscriptionQuery struct {
	conn *sqlx.DB
}

func NewSubscriptionQuery(connection *sqlx.DB) *subscriptionQuery {
	return &subscriptionQuery{
		conn: connection,
	}
}

func (r *subscriptionQuery) GetSubscriptionCount(filter SubscriptionFilter) (int, error) {
	where, args := []string{"status > 0"}, []interface{}{}
	if v := filter.CatID; v != 0 {
		where, args = append(where, "cat_id = ?"), append(args, v)
	}
	if v := filter.SubCatID; v != 0 {
		where, args = append(where, "sub_cat_id = ?"), append(args, v)
	}
	if v := filter.SubTypeID; v != 0 {
		where, args = append(where, "sub_type_id = ?"), append(args, v)
	}
	if v := filter.SourceID; v != 0 {
		where, args = append(where, "source_id = ?"), append(args, v)
	}
	if v := filter.Driver; v != "" {
		where, args = append(where, "driver like ?"), append(args, "%"+v+"%")
	}
	if v := filter.PayTypeID; v != 0 {
		where, args = append(where, "pay_type_id = ?"), append(args, v)
	}
	if v := filter.OpenID; v != "" {
		where, args = append(where, "open_id = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count
		FROM s_subscriptions
		WHERE `+strings.Join(where, " AND "), args...)
	return count, err
}

func (r *subscriptionQuery) GetSubscriptionList(filter SubscriptionFilter) (*[]SubscriptionResponse, error) {
	where, args := []string{"s.status > 0"}, []interface{}{}
	if v := filter.CatID; v != 0 {
		where, args = append(where, "s.cat_id = ?"), append(args, v)
	}
	if v := filter.SubCatID; v != 0 {
		where, args = append(where, "s.sub_cat_id = ?"), append(args, v)
	}
	if v := filter.SubTypeID; v != 0 {
		where, args = append(where, "s.sub_type_id = ?"), append(args, v)
	}
	if v := filter.SourceID; v != 0 {
		where, args = append(where, "s.source_id = ?"), append(args, v)
	}
	if v := filter.Driver; v != "" {
		where, args = append(where, "s.driver like ?"), append(args, "%"+v+"%")
	}
	if v := filter.PayTypeID; v != 0 {
		where, args = append(where, "s.pay_type_id = ?"), append(args, v)
	}
	if v := filter.OpenID; v != "" {
		where, args = append(where, "s.open_id = ?"), append(args, v)
	}
	args = append(args, filter.PageID*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var subscriptions []SubscriptionResponse
	err := r.conn.Select(&subscriptions, `
		SELECT 
		s.id,
		s.cat_id, 
		c.name as cat_name, 
		s.sub_cat_id, 
		sc.name as sub_cat_name, 
		s.sub_type_id,
		st.name as sub_type_name,
		s.source_id,
		ss.name as source_name,
		s.driver,
		s.driver_contact_type_id,
		dct.name as driver_contact_type_name,
		s.driver_contact,
		s.pay_type_id,
		p.name as pay_type_name,
		s.payment_amount,
		s.sub_start_date,
		s.sub_end_date,
		s.need_alert,
		s.alert_before
		FROM s_subscriptions s
		LEFT JOIN c_categories c
		ON s.cat_id = c.id
		LEFT JOIN c_sub_categories sc
		ON s.sub_cat_id = sc.id
		LEFT JOIN c_sub_types st
		ON s.sub_type_id = st.id
		LEFT JOIN c_sources ss
		ON s.source_id = ss.id
		LEFT JOIN c_contact_types dct
		ON s.driver_contact_type_id = dct.id
		LEFT JOIN c_payment_types p
		ON s.pay_type_id = p.id
		WHERE `+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		LIMIT ?, ?
	`, args...)
	return &subscriptions, err
}
