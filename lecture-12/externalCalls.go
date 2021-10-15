package main

import (
	"fmt"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	_http()
	_resty()
	_resty2()
}

func _http() {
	url := `http://localhost:8080/прикол/prikol`

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	//TODO:: что не хватает?

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("http.GET: ", string(body))
}

// https://github.com/go-resty/resty
func _resty() {
	client := resty.New()

	url := `http://localhost:8080/login`
	resp, err := client.R().
		SetFormData(map[string]string{
			"username": "jeeva",
			"password": "mypass",
		}).
		Post(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("resty.POST: ", resp)
}

func _resty2() {
	client := resty.New()

	url := `http://ip.jsontest.com/`
	resp, err := client.R().Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("resty.GET: ", resp)
}
