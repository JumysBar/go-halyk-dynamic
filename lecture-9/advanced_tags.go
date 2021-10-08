package main

import (
	"encoding/json"
	"fmt"
)

type FacebookCredentials struct {
	Login      string `json:"login"`
	Password   string `json:"-"`
	secretName string `json:"secret"`
}

type FacebookAccount struct {
	FirstName string `json:"first_name"`
	// LastName            string `json:"last_name,omitempty"`
	LastName string `json:"last_name"`
	FacebookCredentials
	// FacebookCredentials `json:"cred"`
}

func main() {
	account := &FacebookAccount{
		FirstName: "Vladimir",
		FacebookCredentials: FacebookCredentials{
			Login:      "admin",
			Password:   "admin",
			secretName: "secret",
		},
	}

	data, _ := json.MarshalIndent(account, "", "  ")
	fmt.Printf("JSON: %s\n", data)

	// JSON unmarshal
	x := `
	{
		"first_name": "Vladimir",
		"last_name": "Kim",
		"login": "admin",
		"Password": "admin",
		"secret": "my secret"
	}
	`

	unmarshalAccount := &FacebookAccount{}
	json.Unmarshal([]byte(x), &unmarshalAccount)

	fmt.Printf("first name: %s; last name: %s; login: %s; password: %s; secret: %s\n", unmarshalAccount.FirstName, unmarshalAccount.LastName, unmarshalAccount.Login, unmarshalAccount.Password, unmarshalAccount.secretName)
}
