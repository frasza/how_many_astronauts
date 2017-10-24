/*
Astronauts. Some are in space. Program using API to check how many are there.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type persons struct {
	Number int `json:"number"`
}

func main() {

	url := "http://api.open-notify.org/astros" // API endpoint

	client := http.Client{ // HTTP client
		Timeout: time.Second * 2,
	}

	fetch(url, client)

}

func fetch(url string, client http.Client) {

	req, err := http.NewRequest("GET", url, nil) // Request
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "api-testing")

	res, err := client.Do(req) // Fire request
	if err != nil {
		log.Fatal(err)
	}

	parse(*res)

}

func parse(res http.Response) {

	body, err := ioutil.ReadAll(res.Body) // Get body out of resource
	if err != nil {
		log.Fatal(err)
	}

	persons1 := persons{}
	parsed := json.Unmarshal(body, &persons1) // JSON parse
	if parsed != nil {
		log.Fatal(parsed)
	}

	fmt.Println(persons1.Number)

}
