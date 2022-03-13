package user

import (
	"back/pkg/mysqlClient"
	"context"
)

type Service struct {
	storage *mysqlClient.MySQLClient
}

func NewUserService(mysql *mysqlClient.MySQLClient) *Service {
	return &Service{
		storage: mysql,
	}
}

func (us *Service) Find(ctx context.Context, query map[string]interface{}) ([]*User, error) {
	return nil, nil
}

func (us *Service) FindById(ctx context.Context, id string) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	var user User

	row := us.storage.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.CountryCode,
		&user.RegionCode,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *Service) Create(ctx context.Context, dto CreateUserDTO) (*User, error) {
	return nil, nil
}
func (us *Service) Update(ctx context.Context, dto CreateUserDTO) (*User, error) {
	return nil, nil
}
