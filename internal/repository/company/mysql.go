package company

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/S-L-T/go-assessment/internal/adapter/mysql"
	"github.com/S-L-T/go-assessment/internal/core/domain"
	"os"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQL() (MySQLRepository, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/db?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
	)
	if err != nil {
		return MySQLRepository{}, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return MySQLRepository{
		db: db,
	}, nil

}

func (r MySQLRepository) IsAlive(ctx context.Context) bool {
	err := r.db.PingContext(ctx)
	return err == nil
}

func (r MySQLRepository) Get(ctx context.Context, id uuid.UUID) (domain.Company, error) {
	stmt, err := r.db.Prepare("SELECT id, name, description, total_employees, is_registered, type_id FROM company WHERE id=UUID_TO_BIN(?) LIMIT 1;")
	if err != nil {
		return domain.Company{}, err
	}
	defer stmt.Close()

	res, err := stmt.Query(id.String())
	if err != nil {
		return domain.Company{}, err
	}

	var a mysql.Adapter
	if res.Next() {
		err = res.Scan(&a.ID, &a.Name, &a.Description, &a.TotalEmployees, &a.IsRegistered, &a.Type)
		if err != nil {
			return domain.Company{}, err
		}
	}

	company, err := a.ToDomain()
	if err != nil {
		return domain.Company{}, err
	}

	return company, nil
}

func (r MySQLRepository) Insert(ctx context.Context, company domain.Company) (uuid.UUID, error) {
	stmt, err := r.db.Prepare(
		"INSERT INTO company(name, description, total_employees, is_registered, type_id) VALUES (?,?,?,?,?);")
	if err != nil {
		return uuid.UUID{}, err
	}
	defer stmt.Close()

	a := mysql.NewAdapter(company)
	res, err := stmt.Query(
		a.Name,
		a.Description,
		a.TotalEmployees,
		a.IsRegistered,
		a.Type,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	var id string
	if res.Next() {
		err = res.Scan(&id)
		if err != nil {
			return uuid.UUID{}, err
		}
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedID, nil
}

func (r MySQLRepository) Update(ctx context.Context, company domain.Company) error {
	stmt, err := r.db.Prepare(
		"UPDATE company SET name=?, description=?, total_employees=?, is_registered=?, type_id=? WHERE id=UUID_TO_BIN(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	a := mysql.NewAdapter(company)
	res, err := stmt.Exec(
		a.Name,
		a.Description,
		a.TotalEmployees,
		a.IsRegistered,
		a.Type,
		a.ID,
	)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if ra != 1 {
		return errors.New("update failed; no rows affected")
	}

	return nil
}

func (r MySQLRepository) Delete(ctx context.Context, id uuid.UUID) error {
	stmt, err := r.db.Prepare("DELETE FROM company WHERE id=UUID_TO_BIN(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id.String())
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return errors.New("delete failed; no rows affected")
	}

	return nil
}
