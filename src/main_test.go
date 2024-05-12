package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"
)

func TesaddDJSet(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("../vars/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	addNewArtist(db, "herman")
	addNewPlatform(db, "mixcloud")
	addDJSet(db, 1, 1, "https://www.mixcloud.com/artist/djset", "DJSet")
}

func TestExportDJSets(t *testing.T) {
	// Create a new test database
	ctx := context.Background()
	err := godotenv.Load("../vars/.env")

	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	exportDjSets(db)
}
func TestGetAllDJSets(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("../vars/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	getAllDJSets(db)
}

func TestGetArtists(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("../vars/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

}

func TestGetPlatforms(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("../vars/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	getPlatforms(db)
}
