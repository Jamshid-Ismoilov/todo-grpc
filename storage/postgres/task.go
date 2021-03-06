package postgres

import (
	"database/sql"
	"time"
	"fmt"

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
	var id string
	err := r.db.QueryRow(`
        INSERT INTO tasks(id, assignee, title, summary, deadline, status, created_at)
        VALUES ($6, $1,$2,$3, $4, $5, $7) returning id`, task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, task.Id, task.CreatedAt).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}
	task, err = r.Get(id)

	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task
	var updated sql.NullString

	err := r.db.QueryRow(`
        SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks
        WHERE id=$1 and deleted_at is null`, id).Scan(
			&task.Id, 
			&task.Assignee, 
			&task.Title, 
			&task.Summary, 
			&task.Deadline, 
			&task.Status, 
			&task.CreatedAt, 
			&updated,
		)
	task.UpdatedAt = updated.String

	if err != nil {
		return pb.Task{}, err
	}
	fmt.Println("GET function")

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks LIMIT $1 OFFSET $2`,
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
		var updated sql.NullString
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt, &updated)
		task.UpdatedAt = updated.String
	if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	fmt.Println(tasks)
	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`UPDATE tasks SET assignee=$1, title=$2, summary=$3, deadline=$4, status=$5, updated_at=$7 WHERE id=$6`,
		task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, task.Id, time.Now())
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE tasks SET deleted_at = $2 WHERE id=$1`, id, time.Now())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) ListOverdue(page, limit int64, now string) ([]*pb.Task, int64, error) {
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status FROM tasks where deadline <= $3 LIMIT $1 OFFSET $2;`,
		page,
		limit,
		now,
	)
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

	err = r.db.QueryRow(`SELECT count(*) FROM tasks where deadline >= $1`, now).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
