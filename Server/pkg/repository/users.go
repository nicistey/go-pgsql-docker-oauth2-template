package repository

import (
	"Server/pkg/models" 
	"context"              
)


func (repo *PGRepo) GetUsers() ([]models.User, error) {//curl http://localhost:8090/api/users
	rows, err := repo.pool.Query(context.Background(), 
		`SELECT IDus, IDgoogle, name, email
  		FROM users;
 `)

	if err != nil {
		return nil, err 
	}
	defer rows.Close() 

	var data []models.User 
	for rows.Next() {      
		var item models.User 
		err = rows.Scan(    
			&item.IDus,
			&item.GoogleID,
			&item.Name,
			&item.Email,
		)
		if err != nil {
			return nil, err 
		}
		data = append(data, item) 
	}
	return data, nil 
}
func (repo *PGRepo) NewUser(item models.User) (id int, err error) {//curl -X POST -H "Content-Type: application/json" -d "{\"login\": \"Test9\", \"password\": \"Tes1\"}" localhost:8090/api/users
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
	err = repo.pool.QueryRow(context.Background(), `
	INSERT INTO users ( IDgoogle, name,email )
	VALUES ( $1, $2, $3)
	RETURNING IDus;`,
		&item.GoogleID,
		&item.Name,
		&item.Email,
	).Scan(&id)
	return id, err
}


func (repo *PGRepo) UpdateUser(IDus int,item models.User) (id int, err error) {//curl -X POST -H "Content-Type: application/json" -d "{\"IDus\": 1, \"Event_name\": \"Lol\", \"Event_time\": \"2024-01-26T10:30:00Z\", \"Description\": \"Ex\", \"Location\":\"Tes1222\", \"Is_public\": false}" localhost:8090/api/events/7
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
	err = repo.pool.QueryRow(context.Background(), `
		UPDATE users
		SET name = $2, email = $3
		WHERE IDus = $1
		RETURNING IDus;`,
		&IDus,
		&item.Name,
		&item.Email,
	).Scan(&id)
	return id, err
}

func (repo *PGRepo) GetUserByID(id int) (models.User, error) {//curl http://localhost:8090/api/users
	var user models.User
	err := repo.pool.QueryRow(context.Background(), 
		`SELECT IDus, IDgoogle,name,email
		FROM users
		WHERE IDus =$1;
 `,id).Scan(
	&user.IDus,
	&user.GoogleID,
	&user.Name,
	&user.Email,
 )

	if err != nil {
		return models.User{}, err 
	}

	return user, err 
}

func (repo *PGRepo) GetUserByGoogleID(id string) (models.User, error) {//curl http://localhost:8090/api/users
	var user models.User
	err := repo.pool.QueryRow(context.Background(), 
		`SELECT IDus, IDgoogle,name,email
		FROM users
		WHERE IDgoogle =$1;
 `,id).Scan(
	&user.IDus,
	&user.GoogleID,
	&user.Name,
	&user.Email,
 )

	if err != nil {
		return models.User{}, err 
	}

	return user, err 
}

func (repo *PGRepo) DeleteUser(id int) ( err error) {//curl http://localhost:8090/api/users
	// repo.mu.Lock()
	// defer repo.mu.Unlock()
  	_,err = repo.pool.Exec(context.Background(), `
	DELETE FROM users WHERE IDus =$1;`,
		id,
	)
	return err
}
func (repo *PGRepo) CheckGoogleIDExists(googleID string) (bool, error) {
    var exists bool
    err := repo.pool.QueryRow(context.Background(),`SELECT EXISTS (SELECT 1 FROM users WHERE IDGoogle = $1)`, 
		googleID).Scan(
		&exists)
    if err != nil {
        return false, err 
    }
    return exists, nil
}