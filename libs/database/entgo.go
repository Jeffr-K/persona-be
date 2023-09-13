package database

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	"persona/libs/database/ent"
)

var Client *ent.Client

func InitializeDatabase() {
	client, err := ent.Open("postgres", "host=localhost user=postgres dbname=persona_db sslmode=disable password=root")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	Client = client

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
