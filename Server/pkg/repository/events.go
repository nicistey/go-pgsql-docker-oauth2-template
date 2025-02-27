package repository

import (
	"Server/pkg/models" 
	"context"              
)

func (repo *PGRepo) GetEvents() ([]models.Event, error) {//curl http://localhost:8090/api/events
	rows, err := repo.pool.Query(context.Background(), 
		`SELECT IDev, IDus, event_name, event_time, description,location,is_public
  		FROM events
		WHERE is_public = true;
 `)

	if err != nil {
		return nil, err 
	}
	defer rows.Close() 

	var data []models.Event 
	for rows.Next() {      
		var item models.Event 
		err = rows.Scan(     
			&item.IDev,
			&item.IDus,
			&item.Event_name,
			&item.Event_time,
			&item.Description,
			&item.Location,
			&item.Is_public,
		)
		if err != nil {
			return nil, err 
		}
		data = append(data, item) 
	}
	return data, nil 
}

func (repo *PGRepo) GetEventsByID(id int) ([]models.Event, error) {//curl http://localhost:8090/api/events/1
	rows, err := repo.pool.Query(context.Background(), 
		`SELECT IDev, IDus, event_name, event_time, description,location,is_public
  		FROM events
		WHERE IDus = $1;
 `, id)

	if err != nil {
		return nil, err 
	}
	defer rows.Close() 

	var data []models.Event 
	for rows.Next() {      
		var item models.Event 
		err = rows.Scan(     
			&item.IDev,
			&item.IDus,
			&item.Event_name,
			&item.Event_time,
			&item.Description,
			&item.Location,
			&item.Is_public,
		)
		if err != nil {
			return nil, err 
		}
		data = append(data, item) 
	}
	return data, nil 
}

func (repo *PGRepo) NewEvent(item models.Event, userID int) (id int, err error) {//curl -X POST -H "Content-Type: application/json" -d "{\"IDus\": 1, \"Event_name\": \"Tes1\", \"Event_time\": \"2024-01-26T10:30:00Z\", \"Description\": \"Testik\", \"Location\": \"Tes1\", \"Is_public\": true}" localhost:8090/api/events
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
	err = repo.pool.QueryRow(context.Background(), `
	INSERT INTO events (IDus, event_name, event_time, description, location, is_public )
	VALUES ( $1, $2, $3 , $4 , $5 , $6)
	RETURNING IDev;`,
		userID,
		&item.Event_name,
		&item.Event_time,
		&item.Description,
		&item.Location,
		&item.Is_public,
	).Scan(&id)
	return id, err
}

func (repo *PGRepo) UpdateEvent(IDev int,item models.Event) (id int, err error) {//curl -X POST -H "Content-Type: application/json" -d "{\"IDus\": 1, \"Event_name\": \"Lol\", \"Event_time\": \"2024-01-26T10:30:00Z\", \"Description\": \"Ex\", \"Location\":\"Tes1222\", \"Is_public\": false}" localhost:8090/api/events/7
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
	err = repo.pool.QueryRow(context.Background(), `
		UPDATE events
		SET IDus = $2,
		event_name = $3,
		event_time = $4,
		description = $5,
		location = $6,
		is_public = $7
		WHERE IDev = $1
		RETURNING IDev;`,
		&IDev,
		&item.IDus, // мб это убрать
		&item.Event_name,
		&item.Event_time,
		&item.Description,
		&item.Location,
		&item.Is_public,
	).Scan(&id)
	return id, err
}

func (repo *PGRepo) DeleteEvent(id int) (userID int, err error) {
    err = repo.pool.QueryRow(context.Background(), `
        DELETE FROM events 
        WHERE IDev = $1
        RETURNING IDus;
    `, id).Scan(&userID)
    return userID, err
}