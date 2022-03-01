package todolist

const (
	CreateToDoListQuery = `
	INSERT INTO to_do_list (
		id ,
		name ,
		start_on ,
		qty ,
		status
	) VALUES (?, ?, ?, ?, ?);`
)