package subscription

import (
	"errors"
	"subs/core/database"
	"time"
)

type subscriptionService struct {
}

func NewSubscriptionService() *subscriptionService {
	return &subscriptionService{}
}

func (s *subscriptionService) NewSubscription(info SubscriptionNew) error {
	db := database.WDB()
	tx, err := db.Begin()
	if err != nil {
		msg := "开启事务失败"
		return errors.New(msg)
	}
	defer tx.Rollback()
	repo := NewSubscriptionRepository(tx)
	var newSubscription Subscription
	newSubscription.OpenID = info.OpenID
	newSubscription.CatID = info.CatID
	newSubscription.SubCatID = info.SubCatID
	newSubscription.SubTypeID = info.SubTypeID
	newSubscription.SourceID = info.SourceID
	newSubscription.Driver = info.Driver
	newSubscription.DriverContactTypeID = info.DriverContactTypeID
	newSubscription.DriverContact = info.DriverContact
	newSubscription.PayTypeID = info.PayTypeID
	newSubscription.PaymentAmount = info.PaymentAmount
	newSubscription.SubStartDate = info.SubStartDate
	newSubscription.SubEndDate = info.SubEndDate
	newSubscription.NeedAlert = info.NeedAlert
	newSubscription.AlertBefore = info.AlertBefore
	newSubscription.Status = 1
	newSubscription.Created = time.Now()
	newSubscription.CreatedBy = info.OpenID
	newSubscription.Updated = time.Now()
	newSubscription.UpdatedBy = info.OpenID
	err = repo.CreateSubscription(newSubscription)
	if err != nil {
		msg := "创建订阅失败"
		return errors.New(msg)
	}
	tx.Commit()
	return nil
}

func (s *subscriptionService) UpdateSubscription(id int64, info SubscriptionUpdate) error {
	db := database.WDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	repo := NewSubscriptionRepository(tx)
	subscription, err := repo.GetSubscriptionByID(id, info.OpenID)
	if err != nil {
		msg := "订阅不存在"
		return errors.New(msg)
	}
	if info.CatID != 0 {
		subscription.CatID = info.CatID
	}
	if info.SubCatID != 0 {
		subscription.SubCatID = info.SubCatID
	}
	if info.SubTypeID != 0 {
		subscription.SubTypeID = info.SubTypeID
	}
	if info.SourceID != 0 {
		subscription.SourceID = info.SourceID
	}
	if info.Driver != "" {
		subscription.Driver = info.Driver
	}
	if info.DriverContactTypeID != 0 {
		subscription.DriverContactTypeID = info.DriverContactTypeID
	}
	if info.DriverContact != "" {
		subscription.DriverContact = info.DriverContact
	}
	if info.PayTypeID != 0 {
		subscription.PayTypeID = info.PayTypeID
	}
	if info.PaymentAmount != 0 {
		subscription.PaymentAmount = info.PaymentAmount
	}
	if info.SubStartDate != "" {
		subscription.SubStartDate = info.SubStartDate
	}
	if info.SubEndDate != "" {
		subscription.SubEndDate = info.SubEndDate
	}
	if info.NeedAlert != 0 {
		subscription.NeedAlert = info.NeedAlert
	}
	if info.AlertBefore != 0 {
		subscription.AlertBefore = info.AlertBefore
	}
	subscription.Updated = time.Now()
	subscription.UpdatedBy = info.OpenID
	err = repo.UpdateSubscription(id, *subscription)
	if err != nil {
		msg := "创建扫码历史失败"
		return errors.New(msg)
	}
	tx.Commit()
	return err
}

func (s *subscriptionService) GetSubscriptionList(filter SubscriptionFilter) (int, *[]SubscriptionResponse, error) {
	db := database.RDB()
	query := NewSubscriptionQuery(db)
	count, err := query.GetSubscriptionCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := query.GetSubscriptionList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}
