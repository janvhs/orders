package db

import "github.com/jmoiron/sqlx"

func Seed(db *sqlx.DB) error {
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

func seedUsers(db *sqlx.DB) error {
	_, err := db.Exec(
		`INSERT INTO users (guid, username, user_dn)
VALUES ("1", "admin", "CN=Admin,OU=Users,DC=example,DC=com")`)

	return err
}

func seedOrders(db *sqlx.DB) error {
	_, err := db.Exec(`INSERT INTO orders (author_id)
VALUES (1)`)

	return err
}

func seedArticles(db *sqlx.DB) error {
	_, err := db.Exec(`INSERT INTO articles
(name, amount, url, price, cost_centre, order_id, author_id, for_user_id)
VALUES ("Test", 1, "https://example.com", 1.0, "Test", 1, 1, 1)`)

	return err
}
