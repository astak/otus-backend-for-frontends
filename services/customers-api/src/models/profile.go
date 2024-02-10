package models

type Profile struct {
	ID        int64  `db:"id"`
	AccountId string `db:"account_id"`
	UserName  string `db:"username"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
}
