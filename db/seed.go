package db

import (
	"context"

	"git.bode.fun/orders/db/entity"
	"github.com/uptrace/bun"
)

func Seed(db *bun.DB) error {
	if err := seedUsers(db); err != nil {
		return err
	}

	if err := seedOrders(db); err != nil {
		return err
	}

	if err := seedArticles(db); err != nil {
		return err
	}

	return nil
}

func seedUsers(db *bun.DB) error {
	user := &entity.User{
		// FIXME: Replace with actual GUID Format
		GUID:     "1",
		Username: "admin",
		UserDN:   "CN=Admin,OU=Users,DC=example,DC=com",
	}

	_, err := db.NewInsert().Model(user).Exec(context.TODO())
	if err != nil {
		return err
	}

	user = &entity.User{
		GUID:     "2",
		Username: "user",
		UserDN:   "CN=User,OU=Users,DC=example,DC=com",
	}

	_, err = db.NewInsert().Model(user).Exec(context.TODO())

	return err
}

func seedOrders(db *bun.DB) error {
	order := &entity.Order{
		AuthorID: 1,
	}

	_, err := db.NewInsert().Model(order).Exec(context.TODO())

	return err
}

func seedArticles(db *bun.DB) error {
	article := &entity.Article{
		Name:       "Test",
		Amount:     2,
		URL:        "https://example.com",
		Price:      3.5,
		CostCentre: "IT",
		OrderID:    1,
		AuthorID:   1,
		ForUserID:  2,
	}

	_, err := db.NewInsert().Model(article).Exec(context.TODO())

	return err
}
