package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.randomnumberapi.com/api/v1.0/randomnumber")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var number []int
	json.NewDecoder(resp.Body).Decode(&number)

	fmt.Println("Random number:", number[0])
}

