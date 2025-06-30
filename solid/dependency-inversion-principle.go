package main

import (
	"fmt"
)

/**
# Dependency Inversion Principle (DIP)
- High levels of modules should not be dependent on low levels of modules. Both should follow abstractions (e.g., interfaces).
- One of the popular principles that has been widely used in object-oriented programming is the Dependency Inversion Principle which states that abstractions should not depend on details but rather details should depend on abstractions.
- This principle is in charge of limiting the dependency between two distinct parties or instead between a superior and inferior module, by presenting an abstraction.
- This increases modularity as the system achieves flexibility and can easily be modified to suit change.
**/

type MySQL struct {
}

func (db MySQL) Query() interface{} {
	return []string{"alex", "john", "mike"}
}

type PostgreSQL struct {
}

func (db PostgreSQL) Query() interface{} {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

type DBConn interface {
	Query() interface{}
}

type UsersRepository struct {
	db DBConn
}

func (r UsersRepository) GetUsers() []string {
	var users []string
	res := r.db.Query()

	switch res.(type) {
	case map[string]string:
		for _, u := range res.(map[string]string) {
			users = append(users, u)
		}
		return users
	case []string:
		return res.([]string)
	}

	return []string{}
}

func main() {
	mysqlDB := MySQL{}
	postgreSQLDB := PostgreSQL{}
	repo1 := UsersRepository{db: mysqlDB}
	repo2 := UsersRepository{db: postgreSQLDB}
	fmt.Println(repo1.GetUsers())
	fmt.Println(repo2.GetUsers())
}
