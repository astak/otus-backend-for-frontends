package data

import (
	"context"
	"errors"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type ProfileRepository interface {
	GetProfile(accountId string) (*models.Profile, error)
	UpdateProfile(user *models.Profile) (*models.Profile, error)
}

type PostgresProfileRepository struct {
	db *Database
}

func NewProfileRepository(db *Database) (*PostgresProfileRepository, error) {
	return &PostgresProfileRepository{db: db}, nil
}

func (repo PostgresProfileRepository) GetProfile(accountId string) (*models.Profile, error) {
	sql := `
	SELECT id, account_id, username, firstname, lastname, email, phone
	FROM profiles
	WHERE account_id = $1
	`
	var user models.Profile
	rows, err := repo.db.Conn.Query(context.Background(), sql, accountId)
	if err != nil {
		panic(err)
	}
	err = pgxscan.ScanOne(&user, rows)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		panic(err)
	}
	return &user, nil
}

func (repo PostgresProfileRepository) UpdateProfile(user *models.Profile) (*models.Profile, error) {
	sql := `
	UPDATE profiles SET
		username = $2,
		firstname = $3,
		lastname = $4,
		email = $5,
		phone = $6
	WHERE account_id = $1
	`
	res, err := repo.db.Conn.Exec(context.Background(), sql, user.AccountId, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		panic(err)
	}
	count := res.RowsAffected()
	if count > 0 {
		return user, nil
	}
	sql = `
	INSERT INTO profiles (account_id, username, firstname, lastname, email, phone)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = repo.db.Conn.Exec(context.Background(), sql, user.AccountId, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		panic(err)
	}
	return user, nil
}
