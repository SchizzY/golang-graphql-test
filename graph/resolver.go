package graph

import (
	"database/sql"
	"example/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore map[string]model.Character
	//add to my db
	DB *sql.DB
 }
