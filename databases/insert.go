package main

import "fmt"

func addActor(firstname, lastname string) (int64, error) {
	result, err := db.Exec("INSERT INTO actor (firstname, lastname) VALUES (?, ?)", firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor Insert: %v", err)
	}

	return id, nil
}
