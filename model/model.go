package model

type Todo struct {
	Id   uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func CreateTodo(todo string) err {

	statement := `insert into todos (todo, done) values ($1, $2);`

	_, err := db.Query(statement, todo, false)

	return err

}

func GetAllTodos() ([]Todo, error) {

	statement := `select * from todos;`

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		var title string
		var done bool
		var id uint64

		err := db.Scan(&id, &title, &done)

		if err != nil {
			return todos, err
		}

		todo := Todo{
			Id:   id,
			Todo: title,
			Done: done,
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func GetTodo(id uint64) (Todo, error) {

	statement := `select * from todos where id = $1;`

	todo := Todo{}
	todo.Id = id

	row, err := db.QueryRow(statement, id)
	if err != nil {
		return todo, err
	}

	for row.Next() {
		var title string
		var done bool

		err := row.Scan(&title, &done)
		if err != nil {
			return todo, err
		}
		todo.Todo = title
		todo.Done = done
	}
	return todo, err
}

func MarkDone(id uint64) error {

	todo, err := GetTodo(id)
	if err != nil {
		return err
	}

	statement := `update todos set done = $2 where id = $1;`

	_, err = db.Query(statement, id, !todo.Done)

	return err
}

func Delete(id uint64) error {

	statement := `delete from todos where id = $1;`

	_, err := db.Exec(statement, id)

	return err
}
