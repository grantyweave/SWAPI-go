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

	url := "https://swapi.dev/api/planets/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var planetResponse PlanetResponse
	err = json.Unmarshal(body, &planetResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("List of planets:")
	for _, planet := range planetResponse.Results {
		fmt.Println(planet.Name)
	}

	fmt.Println()

	sort.Slice(planetResponse.Results, func(i, j int) bool {
		return planetResponse.Results[i].Name < planetResponse.Results[j].Name
	})

	fmt.Println("List of planets (sorted alphabetically):")
	for _, planet := range planetResponse.Results {
		fmt.Println(planet.Name)
	}
}