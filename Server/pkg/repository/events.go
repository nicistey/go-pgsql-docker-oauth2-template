package repository

import (
	"Server/pkg/models" 
	"context"              
)
func (repo *PGRepo) GetEventsByID(IDus int) ([]models.Event, error) {//curl http://localhost:8090/api/events
	rows, err := repo.pool.Query(context.Background(), 
		`SELECT  IDev, IDus, event_name, event_time, description,location,is_public
		FROM events
		WHERE IDev =$1;
 `,IDus)

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
// 

func (repo *PGRepo) GetEvents() ([]models.Event, error) {//curl http://localhost:8090/api/events
	rows, err := repo.pool.Query(context.Background(), 
		`SELECT IDev, IDus, event_name, event_time, description,location,is_public
  		FROM events;
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
func (repo *PGRepo) NewEvent(item models.Event) (id int, err error) {//curl -X POST -H "Content-Type: application/json" -d "{\"IDus\": 1, \"Event_name\": \"Tes1\", \"Event_time\": \"2024-01-26T10:30:00Z\", \"Description\": \"Testik\", \"Location\": \"Tes1\", \"Is_public\": true}" localhost:8090/api/events
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
	err = repo.pool.QueryRow(context.Background(), `
	INSERT INTO events (IDus, event_name, event_time, description, location, is_public )
	VALUES ( $1, $2, $3 , $4 , $5 , $6)
	RETURNING IDev;`,
		&item.IDus,
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

func (repo *PGRepo) GetEventByID(id int) (models.Event, error) {//curl http://localhost:8090/api/events/1
	var event models.Event
	err := repo.pool.QueryRow(context.Background(), // Выполняем запрос к базе данных. context.Background() используется в качестве контекста.
		`SELECT  IDev, IDus, event_name, event_time, description,location,is_public
		FROM events
		WHERE IDev =$1;
 `,id).Scan(
		&event.IDev,
		&event.IDus,
		&event.Event_name,
		&event.Event_time,
		&event.Description,
		&event.Location,
		&event.Is_public,
 )

	if err != nil {
		return models.Event{}, err // Возвращаем ошибку, если запрос к базе данных завершился с ошибкой.
	}

	return event, err // Возвращаем массив книг и nil, если ошибок не было.
}

func (repo *PGRepo) DeleteEvent(id int) ( err error) {//curl -X DELETE localhost:8090/api/events/6
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
  	_,err = repo.pool.Exec(context.Background(), `
	DELETE FROM events WHERE IDev =$1;`,
		id,
	)
	return err
}