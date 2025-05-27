package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.randomnumberapi.com/api/v1.0/randomnumber"

	// Realiza la solicitud GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}


	var numbers []int
	err = json.Unmarshal(body, &numbers)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return
	}
	fmt.Prin
