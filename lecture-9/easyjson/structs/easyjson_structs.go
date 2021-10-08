package easyjson_struct

import (
	_ "github.com/mailru/easyjson/gen"
)

type Credentials struct {
	Login    string
	Password string
}

type Language struct {
	Title string
	Words []string
}

type Profile struct {
	Name      string
	Age       int64
	Languages []Language
	Credentials
}
