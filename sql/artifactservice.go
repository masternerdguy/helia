package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_artifacts table
type ArtifactService struct{}

// Gets a artifact service for interacting with artifacts in the database
func GetArtifactService() *ArtifactService {
	return &ArtifactService{}
}

// Structure representing a row in the universe_artifacts table
type Artifact struct {
	ID           uuid.UUID
	SystemID     uuid.UUID
	ArtifactName string
	PosX         float64
	PosY         float64
	Texture      string
	Radius       float64
	Mass         float64
	Theta        float64
	Meta         Meta
}

// Retrieves all artifacts from the database
func (s ArtifactService) GetAllArtifacts() ([]Artifact, error) {
	artifacts := make([]Artifact, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load artifacts
	sql := `
				SELECT id, universe_systemid, artifactname, pos_x, pos_y, texture, radius, mass, theta, meta
				FROM public.universe_artifacts;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Artifact{}

		// scan into artifact structure
		rows.Scan(&r.ID, &r.SystemID, &r.ArtifactName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta, &r.Meta)

		// append to artifact slice
		artifacts = append(artifacts, r)
	}

	return artifacts, err
}

// Retrieves all artifacts in a given solar system from the database
func (s ArtifactService) GetArtifactsBySolarSystem(systemID uuid.UUID) ([]Artifact, error) {
	artifacts := make([]Artifact, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load artifacts
	sql := `
				SELECT id, universe_systemid, artifactname, pos_x, pos_y, texture, radius, mass, theta, meta
				FROM public.universe_artifacts
				WHERE universe_systemid = $1;
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Artifact{}

		// scan into artifact structure
		rows.Scan(&r.ID, &r.SystemID, &r.ArtifactName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta, &r.Meta)

		// append to artifact slice
		artifacts = append(artifacts, r)
	}

	return artifacts, err
}
