package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(c *config) error {
	defaultUrl := "https://pokeapi.co/api/v2/location-area"
	var res *http.Response
	var err error
	if c.Next == "" {
		res, err = http.Get(defaultUrl)
	} else {
		res, err = http.Get(c.Next)
	}
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	location := Location{}

	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(location.Results); i++ {
		fmt.Println(location.Results[i].Name)
	}
	c.Next = location.Next
	c.Previous = location.Previous

	return nil
}
