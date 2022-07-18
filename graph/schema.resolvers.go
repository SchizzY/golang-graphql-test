package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"fmt"
	"strconv"
)

// UpsertCharacter is the resolver for the upsertCharacter field.
func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	id := input.ID
	var character model.Character
	character.Name = input.Name
	character.CliqueType = input.CliqueType
	

	//get a count from the database and scan it to present the data in the graphql response
	res, err := r.Resolver.DB.Query("SELECT COUNT(*) FROM characters")
	if err != nil {
		r.Resolver.DB.Exec("CREATE TABLE characters (id int NOT NULL AUTO_INCREMENT, name varchar(255), is_hero TINYINT,clique_type varchar(255), PRIMARY KEY (id))")

		r.Resolver.DB.Exec("INSERT INTO characters (name, is_hero, clique_type) VALUES (?, ?, ?)", character.Name, character.IsHero, character.CliqueType)
		return &character, nil
	}
	var count int
	res.Next()
	res.Scan(&count)
	character.ID = strconv.Itoa(count + 1)

	
	// n := len(r.Resolver.CharacterStore)
	if count == 0 {
		fmt.Println(r.Resolver.DB.Exec("INSERT INTO characters (name, is_hero, clique_type) VALUES (?, ?, ?)", character.Name, character.IsHero, character.CliqueType))
	}

	if id != nil {
		cs, ok := r.Resolver.CharacterStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		} else {
			character.IsHero = cs.IsHero
		}
		r.Resolver.CharacterStore[*id] = character
	} else {
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		}

	}

	r.Resolver.DB.Query("INSERT INTO characters (id, name, is_hero, clique_type) VALUES (?, ?, ?, ?)", character.ID, character.Name, character.IsHero, character.CliqueType)

	return &character, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	res, err := r.Resolver.DB.Query("SELECT * FROM characters WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	var character model.Character
	res.Next()
	res.Scan(&character.ID, &character.Name, &character.IsHero, &character.CliqueType)
	return &character, nil
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context, cliqueType model.CliqueType) ([]*model.Character, error) {
	res, err := r.Resolver.DB.Query("SELECT * FROM characters WHERE clique_type = ?", cliqueType)
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	var characters []*model.Character
	for res.Next() {
		var character model.Character
		res.Scan(&character.ID, &character.Name, &character.IsHero, &character.CliqueType)
		characters = append(characters, &character)
	}

	return characters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}
