package sql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the containers table in the database
type ContainerService struct{}

// Returns a container service for interacting with containers in the database
func GetContainerService() ContainerService {
	return ContainerService{}
}

// Structure representing a row in the containers table
type Container struct {
	ID      uuid.UUID
	Meta    Meta `json:"meta"`
	Created time.Time
}

// Creates a new container
func (s ContainerService) NewContainer(e Container) (*Container, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//insert container
	sql := `
				INSERT INTO public.containers(
					id, meta, created
				)
				VALUES ($1, $2, $3);
			`

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, e.Meta, createdAt)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	//update id in model
	e.ID = uid
	e.Created = createdAt

	//return pointer to inserted container model
	return &e, nil
}

// Finds and returns a container by its id
func (s ContainerService) GetContainerByID(containerID uuid.UUID) (*Container, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//find container with this id
	container := Container{}

	sqlStatement :=
		`
			SELECT id, meta, created
			FROM public.containers
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, containerID)

	switch err := row.Scan(&container.ID, &container.Meta, &container.Created); err {
	case sql.ErrNoRows:
		return nil, errors.New("Container not found")
	case nil:
		return &container, nil
	default:
		return nil, err
	}
}

// Updates a container in the database
func (s ContainerService) UpdateContainer(container Container) error {
	//get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	//update container in database
	sqlStatement :=
		`
			UPDATE public.containers
			SET meta=$2, created=$3
			WHERE id = $1
		`

	_, err = db.Exec(sqlStatement, container.ID, container.Meta, container.Created)

	return err
}
