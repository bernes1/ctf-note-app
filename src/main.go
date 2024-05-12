package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"
)

type Platform struct {
	Id   int
	Name string
}

type Artist struct {
	Id   int
	Name string
}

type DJset struct {
	Id           int64
	Name         string
	Url          string
	PlatformId   int64
	PlatformName string
	ArtistId     int64
	ArtistName   string
}

func main() {
	ctx := context.Background()
	err := godotenv.Load("vars/.env")

	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	fmt.Println("Startup successful")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter 'add' to add a new DJ set, 'export' to export all DJ sets, or 'all' to list all DJ sets")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "add" {
		addNewDJSet(pool)
	} else if input == "export" {
		exportDjSets(pool)
	} else if input == "all" {
		getAllDJSets(pool)
	} else {
		fmt.Println("Invalid input")
	}
}

func getArtists(db *pgxpool.Pool) {
	ctx := context.Background()
	var artists []*Artist
	pgxscan.Select(ctx, db, &artists, "SELECT * FROM artist")
	fmt.Println("Artists:")
	for _, artist := range artists {
		fmt.Printf("id: %s , name: %s\n", strconv.Itoa(artist.Id), artist.Name)
	}
}

func addNewArtist(db *pgxpool.Pool, artistName string) (int, error) {
	// Insert artist and get ID
	var artistID int
	err := db.QueryRow(context.Background(), "INSERT INTO artist (name) VALUES ($1) RETURNING id", artistName).Scan(&artistID)
	if err != nil {
		return 0, err
	}
	return artistID, nil
}

func getPlatforms(db *pgxpool.Pool) {
	ctx := context.Background()
	var platforms []*Platform
	pgxscan.Select(ctx, db, &platforms, "SELECT * FROM platform")
	fmt.Println("Platforms:")
	for _, platform := range platforms {
		fmt.Printf("ID:%s, Name: %s\n", strconv.Itoa(platform.Id), platform.Name)
	}

}

func addNewPlatform(db *pgxpool.Pool, platformName string) (int, error) {
	var platformID int
	err := db.QueryRow(context.Background(), "INSERT INTO platform (name) VALUES ($1) RETURNING id", platformName).Scan(&platformID)
	if err != nil {
		return 0, err
	}
	return platformID, nil
}

func addDJSet(db *pgxpool.Pool, artistID int, platformID int, djsetName string, djsetUrl string) error {
	// Insert DJ set
	_, err := db.Exec(context.Background(), "INSERT INTO djset (name, url, platform_id, artist_id) VALUES ($1, $2, $3, $4)", djsetName, djsetUrl, platformID, artistID)
	if err != nil {
		fmt.Println("Error inserting into djsets:", err)
		return err
	}
	return nil
}
func getAllDJSets(db *pgxpool.Pool) {
	ctx := context.Background()
	var djsets []*DJset
	pgxscan.Select(ctx, db, &djsets, `SELECT djset.*, platform.name as platform_name, artist.name as artist_name FROM djset 
                                      JOIN platform ON djset.platform_id = platform.id
                                      JOIN artist ON djset.artist_id = artist.id`)

	for _, djset := range djsets {
		fmt.Printf("ID: %d, Name: %s, URL: %s, Platform: %s, Artist: %s\n", djset.Id, djset.Name, djset.Url, djset.PlatformName, djset.ArtistName)
	}
}

func addNewDJSet(db *pgxpool.Pool) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--Information about artist and platform--")
	getArtists(db)
	fmt.Println("")
	getPlatforms(db)
	fmt.Printf("-----------------------------------------\n")

	// Get the artist ID
	fmt.Print("Enter artist ID or 'new' to add a new artist: ")
	artistInput, _ := reader.ReadString('\n')
	artistInput = strings.TrimSpace(artistInput)

	var artistID int
	var err error
	if artistInput == "new" {
		fmt.Print("Enter new artist name: ")
		artistName, _ := reader.ReadString('\n')
		artistName = strings.TrimSpace(artistName)

		// Add the new artist to the database and get the ID
		artistID, err = addNewArtist(db, artistName)
		if err != nil {
			return err
		}
	} else {
		artistID, _ = strconv.Atoi(artistInput)
	}

	// Do the same for the platform
	fmt.Print("Enter platform ID or 'new' to add a new platform: ")
	platformInput, _ := reader.ReadString('\n')
	platformInput = strings.TrimSpace(platformInput)

	var platformID int
	if platformInput == "new" {

		fmt.Print("Enter new platform name: ")
		platformName, _ := reader.ReadString('\n')
		platformName = strings.TrimSpace(platformName)

		// Add the new platform to the database and get the ID
		platformID, err = addNewPlatform(db, platformName)
		if err != nil {
			return err
		}
	} else {
		platformID, _ = strconv.Atoi(platformInput)
	}

	// Get the DJ set details
	fmt.Print("Enter DJ set name: ")
	djsetName, _ := reader.ReadString('\n')
	djsetName = strings.TrimSpace(djsetName)

	fmt.Print("Enter DJ set URL: ")
	djsetUrl, _ := reader.ReadString('\n')
	djsetUrl = strings.TrimSpace(djsetUrl)

	// Add the new DJ set to the database
	err = addDJSet(db, artistID, platformID, djsetName, djsetUrl)
	if err != nil {
		return err
	}

	return nil

}

func exportDjSets(db *pgxpool.Pool) {
	//TODO export all djset in json
	ctx := context.Background()
	var djsets []*DJset
	pgxscan.Select(ctx, db, &djsets, `SELECT djset.*, platform.name as platform_name, artist.name as artist_name FROM djset 
                                      JOIN platform ON djset.platform_id = platform.id
                                      JOIN artist ON djset.artist_id = artist.id`)

	djsetList := make([]*DJset, 0)
	for _, djset := range djsets {
		djsetList = append(djsetList, djset)
		fmt.Println(djset)
	}
	jsonexport, err := json.Marshal(djsetList)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
	}
	err = os.WriteFile("djsetexport.json", jsonexport, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing json file:", err)
	}
}
