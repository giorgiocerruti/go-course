package main

import "fmt"

func getActor(actorID int32) ([]Actor, error) {
	var actors []Actor

	query := "SELECT id, firstname, lastname FROM actor where id  = ?"

	result, err := db.Query(query, actorID)
	if err != nil {
		return nil, fmt.Errorf("Erro fatching data: %v", err)
	}
	defer db.Close()

	for result.Next() {
		var actor Actor

		if err := result.Scan(&actor.Id, &actor.Firstname, &actor.Lastname); err != nil {
			return nil, fmt.Errorf("Error scanning result %v", err)
		}

		actors = append(actors, actor)

		if err = result.Err(); err != nil {
			return nil, fmt.Errorf("Scan error: %v", err)
		}
	}

	return actors, nil

}
