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

func (s *Storage) GetListByName(ctx context.Context, name string) (*[]Company, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE fullName LIKE ?", tableName)
	var companies []Company

	rows, err := s.client.Db.QueryContext(ctx, query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var company Company
		err = rows.Scan(
			&company.Id,
			&company.FullName,
			&company.ShortName,
			&company.Description,
		)
		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}
	return &companies, nil
}
