package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	pb "github.com/Jamshid-Ismoilov/todo-grpc/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	var id int64
	err := r.db.QueryRow(`
        INSERT INTO tasks(assignee, title, summary, deadline, status)
        VALUES ($1,$2,$3, $4, $5) returning id`, task.Assignee, task.Title, task.Summary, task.Deadline, task.Status).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id int64) (pb.Task, error) {
	var task pb.Task
	err := r.db.QueryRow(`
        SELECT id, assignee, title, summary, deadline, status FROM tasks
        WHERE id=$1`, id).Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status FROM tasks LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		tasks []*pb.Task
		task  pb.Task
		count int64
	)
	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
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
