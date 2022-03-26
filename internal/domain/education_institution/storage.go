package education_institution

import (
	"back/pkg/mysqlClient"
	"context"
	"fmt"
)

const (
	tableName = `edu_institution`
)

type Storage struct {
	client *mysqlClient.MySQLClient
}

func NewEducationInstitutionStorage(mysql *mysqlClient.MySQLClient) *Storage {
	return &Storage{
		client: mysql,
	}
}

func (s *Storage) Create(ctx context.Context, dto CreateEducationInstitutionDTO) (id int64, err error) {
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

func (s *Storage) GetListByName(ctx context.Context, name string) (*[]EducationInstitution, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE fullName LIKE ?", tableName)
	var eduList []EducationInstitution

	rows, err := s.client.Db.QueryContext(ctx, query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var eduInst EducationInstitution
		err = rows.Scan(
			&eduInst.Id,
			&eduInst.FullName,
			&eduInst.ShortName,
			&eduInst.Description,
		)
		if err != nil {
			return nil, err
		}

		eduList = append(eduList, eduInst)
	}
	return &eduList, nil
}
