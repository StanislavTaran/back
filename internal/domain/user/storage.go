package user

import (
	"back/pkg/mysqlClient"
	"context"
	"fmt"
	"github.com/google/uuid"
)

const (
	tableName = `users`
)

type Storage struct {
	client *mysqlClient.MySQLClient
}

func NewUserStorage(mysql *mysqlClient.MySQLClient) *Storage {
	return &Storage{
		client: mysql,
	}
}

func (s *Storage) FindById(ctx context.Context, id string) (*User, error) {
	query := "SELECT id, firstName, lastName, email,isActive, createdAt, updatedAt FROM users WHERE id = ?"
	var user User

	row := s.client.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) GetUserPassword(ctx context.Context, email string) (string, error) {
	query := "SELECT password FROM users WHERE email = ?"
	var password string

	row := s.client.Db.QueryRowContext(ctx, query, email)
	err := row.Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (s *Storage) Create(ctx context.Context, dto CreateUserDTO) (string, error) {
	var id = uuid.New().String()
	query := fmt.Sprintf(
		"INSERT INTO %s (id, firstName, lastName, email, password) VALUES (?,?,?,?,?)", tableName,
	)

	passHash, err := generatePassHash(dto.Password)
	if err != nil {
		return "", err
	}

	_, err = s.client.Db.ExecContext(ctx, query, id, &dto.FirstName, &dto.LastName, &dto.Email, passHash)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Storage) ActivateUser(ctx context.Context, id string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET isActive = 1 WHERE id=?", tableName,
	)

	_, err := s.client.Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
