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
	query := "SELECT id, firstName, lastName, email, createdAt, updatedAt FROM users WHERE id = ?"
	var user User

	row := s.client.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) CollectUserInfoById(ctx context.Context, id string) (*FullUserInfoDTO, error) {
	query := `SELECT 
       u.id, 
       u.firstName, 
       u.lastName, 
       u.dataOfBirth, 
       u.email, 
       u.shortInfo, 
       r.role, 
       u.createdAt, 
       u.updatedAt,
       
       ue.id,
       ei.id,
       ei.fullName, 
       ue.faculty, 
       ue.inProgress, 
       ue.startDate, 
       ue.endDate,
       
       uc.id,
       c.id,
       c.fullName,
       empt.type, 
       uc.jobTitle, 
       uc.inProgress, 
       uc.startDate, 
       uc.endDate 
FROM users u 
    LEFT JOIN user_company uc ON (u.id = uc.userId)
    LEFT JOIN company c ON (uc.companyId = c.id)
    LEFT JOIN user_education ue ON (u.id = ue.userId)
    LEFT JOIN edu_institution ei ON (ue.eduInstitutionId = ei.id)
    INNER JOIN role r ON (u.roleId = r.id) 
    LEFT JOIN employment_type empt ON (uc.employmentTypeId = empt.id) 
WHERE u.id = ?`

	var user FullUserInfoDTO

	rows, err := s.client.Db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var eduInstitution eduInstitution
		var company company
		var userJob jobUserInfo
		var userEdu educationUserInfo

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.DateOfBirth,
			&user.Email,
			&user.ShortInfo,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,

			&userEdu.Id,
			&eduInstitution.Id,
			&eduInstitution.Name,
			&userEdu.Faculty,
			&userEdu.InProgress,
			&userEdu.StartDate,
			&userEdu.EndDate,

			&userJob.Id,
			&company.Id,
			&company.Name,
			&userJob.EmploymentType,
			&userJob.JobTitle,
			&userJob.InProgress,
			&userJob.StartDate,
			&userJob.EndDate,
		)
		if err != nil {
			return nil, err
		}

		if eduInstitution.Id.Valid {
			userEdu.EduInstitution = &eduInstitution
			user.Education = append(user.Education, userEdu)
		}

		if company.Id.Valid {
			userJob.Company = &company
			user.JobExperience = append(user.JobExperience, userJob)
		}
	}

	if user.Education == nil {
		user.Education = []educationUserInfo{}
	}
	if user.JobExperience == nil {
		user.JobExperience = []jobUserInfo{}
	}

	return &user, nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := "SELECT id,firstName, lastName, email, password FROM users WHERE email = ?"
	var user User

	row := s.client.Db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) Create(ctx context.Context, dto CreateUserDTO) (string, error) {
	var id = uuid.New().String()
	query := fmt.Sprintf(
		"INSERT INTO %s (id, firstName, lastName, email, password, roleId) VALUES (?,?,?,?,?,?)", tableName,
	)

	passHash, err := generatePassHash(dto.Password)
	if err != nil {
		return "", err
	}

	_, err = s.client.Db.ExecContext(ctx, query, id, &dto.FirstName, &dto.LastName, &dto.Email, passHash, 1)
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
