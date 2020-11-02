package models

type User struct {
	Id        uint64
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Status    string
}

func NewUser(user User) (bool, error) {
	con := Connect()
	defer con.Close()

	sql := "insert into users (firstname, lastname, email, password) values($1, $2, $3, $4)"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}
