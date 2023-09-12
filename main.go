package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type Planet struct {
	Name string `json:"name"`
}

type PlanetResponse struct {
	Results []Planet `json:"results"`
}

func main() {
	// URL for the SWAPI planets endpoint
	url := "https://swapi.dev/api/planets/"

	// Make an HTTP GET request to the SWAPI
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var planetResponse PlanetResponse
	err = json.Unmarshal(body, &planetResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the list of planets
	fmt.Println("List of planets:")
	for _, planet := range planetResponse.Results {
		fmt.Println(planet.Name)
	}

	fmt.Println()

	// Sort the list of planets alphabetically
	sort.Slice(planetResponse.Results, func(i, j int) bool {
		return planetResponse.Results[i].Name < planetResponse.Results[j].Name
	})

	// Print the sorted list of planets
	fmt.Println("List of planets (sorted alphabetically):")
	for _, planet := range planetResponse.Results {
		fmt.Println(planet.Name)
	}
}