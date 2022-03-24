package company

import (
	"back/pkg/mysqlClient"
	"context"
	"fmt"
)

const (
	tableName = `company`
)

type Storage struct {
	client *mysqlClient.MySQLClient
}

func NewCompanyStorage(mysql *mysqlClient.MySQLClient) *Storage {
	return &Storage{
		client: mysql,
	}
}

func (s *Storage) Create(ctx context.Context, dto CreateCompanyDTO) (id int64, err error) {
	query := fmt.Sprintf("INSERT INTO %s (fullName,shortName,description) VALUES(?,?,?)", tableName)

	res, err := s.client.Db.ExecContext(
		ctx,
		query,
		&dto.FullName,
		&dto.ShortName,
		&dto.Description,
	)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
