package user_company

import (
	"back/pkg/mysqlClient"
	"context"
	"fmt"
)

const (
	tableName = `user_company`
)

type Storage struct {
	client *mysqlClient.MySQLClient
}

func NewUserCompanyStorage(mysql *mysqlClient.MySQLClient) *Storage {
	return &Storage{
		client: mysql,
	}
}

func (s *Storage) Create(ctx context.Context, dto CreateUserJobExperienceDTO) (int64, error) {
	var query = fmt.Sprintf("INSERT INTO %s (userId,companyId,employmentTypeId,jobTitle,inProgress,startDate,endDate) VALUES(?,?,?,?,?,?,?)", tableName)

	res, err := s.client.Db.ExecContext(
		ctx,
		query,
		&dto.UserId,
		&dto.CompanyId,
		&dto.EmploymentTypeId,
		&dto.JobTitle,
		&dto.InProgress,
		&dto.StartDate,
		&dto.EndDate,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
