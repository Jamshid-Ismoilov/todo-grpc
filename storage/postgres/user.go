package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	pb "github.com/rustagram/template-service/genproto"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user pb.User) (pb.User, error) {
	var id int64
	err := r.db.QueryRow(`
        INSERT INTO users(first_name, last_name)
        VALUES ($1,$2) returning id`, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return pb.User{}, err
	}

	user, err = r.Get(id)
	if err != nil {
		return pb.User{}, err
	}

	return user, nil
}

func (r *userRepo) Get(id int64) (pb.User, error) {
	var user pb.User
	err := r.db.QueryRow(`
        SELECT id, first_name, last_name FROM users
        WHERE id=$1`, id).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return pb.User{}, err
	}

	return user, nil
}

func (r *userRepo) List(page, limit int64) ([]*pb.User, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, first_name, last_name FROM users LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		users []*pb.User
		user  pb.User
		count int64
	)
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM users`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *userRepo) Update(user pb.User) (pb.User, error) {
	result, err := r.db.Exec(`UPDATE users SET first_name=$1, last_name=$2 WHERE id=$3`,
		user.FirstName, user.LastName, user.Id)
	if err != nil {
		return pb.User{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.User{}, sql.ErrNoRows
	}

	user, err = r.Get(user.Id)
	if err != nil {
		return pb.User{}, err
	}

	return user, nil
}

func (r *userRepo) Delete(id int64) error {
	result, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
