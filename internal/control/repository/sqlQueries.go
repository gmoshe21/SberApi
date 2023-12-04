package repository

const (
	queryAddNewTask = `INSERT INTO task (
			id,
			title,
			description,
			data,
			checking_completion)
		VALUES($1,$2,$3,$4,'false');` 

	queryGetInfoTask = `SELECT TO_JSON (rows) FROM (
			SELECT *
			FROM task
			WHERE id = $1
		) AS rows;`

	queryUpdateTask = `UPDATE task SET checking_completion = NOT checking_completion 
		WHERE id = $1;`

	queryDeleteTask = `DELETE FROM task WHERE id = $1;`

	queryGetAllTask = `SELECT JSON_AGG (rows) FROM (
			SELECT id, title, checking_completion
			FROM task
		) AS rows;`
)