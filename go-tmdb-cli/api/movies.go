package api

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Movie struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Popularity float64 `json:"popularity"`
	ReleaseDate string `json:"release_date"`
}

func FetchMovies(movieType string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("TMDB_API_KEY")
	baseURL := "https://api.themoviedb.org/3/movie/"

	var url string
	switch movieType {
	case "playing":
		url = baseURL + "now_playing?page=1"
	case "popular":
		url = baseURL + "popular?page=1"
	case "top":
		url = baseURL + "top_rated?page=1"
	case "upcoming":
		url = baseURL + "upcoming?page=1"
	default:
		return fmt.Errorf("invalid movie type: %s. Use one of: playing, popular, top, upcoming", movieType)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	var data struct {
		Results []Movie `json:"results"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("error parsing JSON response: %v", err)
	}

	if len(data.Results) == 0 {
		return fmt.Errorf("no movies found for the type: %s", movieType)
	}

	fmt.Println("\033[1;36m=====================================")
	fmt.Printf("           \033[1;36m%s Movies\033[0m\n", strings.Title(movieType))
	fmt.Println("\033[1;36m=====================================\n")

	for _, movie := range data.Results {
		fmt.Printf("\033[1;34mTitle:\033[0m %s\n", movie.Title)
		fmt.Printf("\033[1;32mOverview:\033[0m %s\n", movie.Overview)
		fmt.Printf("\033[1;33mPopularity:\033[0m %f\n", movie.Popularity)
		fmt.Printf("\033[1;35mReleaseDate:\033[0m %s\n", movie.ReleaseDate)
		fmt.Println("\033[1;36m---------------------------------------\n")
	}

	return nil
}
