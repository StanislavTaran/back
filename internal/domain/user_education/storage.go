package user_education

import (
	"back/pkg/mysqlClient"
	"context"
	"fmt"
)

const (
	tableName = `user_education`
)

type Storage struct {
	client *mysqlClient.MySQLClient
}

func NewUserEducationStorage(mysql *mysqlClient.MySQLClient) *Storage {
	return &Storage{
		client: mysql,
	}
}

func (s *Storage) Create(ctx context.Context, dto CreateUserEducationDTO) (int64, error) {
	var query = fmt.Sprintf("INSERT INTO %s (userId,eduInstitutionId,faculty,inProgress,startDate,endDate) VALUES(?,?,?,?,?,?)", tableName)

	res, err := s.client.Db.ExecContext(
		ctx,
		query,
		&dto.UserId,
		&dto.EduInstitutionId,
		&dto.Faculty,
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
