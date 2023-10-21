package database

import (
	"go-clean-architecture/src/domain"
	"log"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		log.Fatalf("Execute error: %v", err)
	}

	id64, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("LastInsertId error: %v", err)
	}
	id = int(id64)
	return
}

// identifier 識別子 アイデンティファイア
func (repo *UserRepository) FindById(identifier int) (domain.User, error) {
	row, err := repo.Query(
		"SELECT id, first_name, last_name FROM users WHERE id = ?", identifier,
	)
	defer row.Close()
	if err != nil {
		log.Fatalf("Query error: %v", err)
	}

	var id int
	var firstName string
	var lastName string

	row.Next()

	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		log.Fatalf("Query error: %v", err)
	}

	return domain.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (repo *UserRepository) FindAll() (domain.Users, error) {
	row, err := repo.Query(
		"SELECT id, first_name, last_name FROM users",
	)
	defer row.Close()
	if err != nil {
		log.Fatalf("Query error: %v", err)
	}

	var users domain.Users

	for row.Next() {
		var id int
		var firstName string
		var lastName string

		if err = row.Scan(&id, &firstName, &lastName); err != nil {
			log.Printf("Query error: %v", err)
			continue
		}

		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return users, nil
}
