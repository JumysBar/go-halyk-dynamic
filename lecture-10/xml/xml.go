package main

import (
	"encoding/xml"
	"fmt"
)

// our struct which contains the complete
// array of all Users in the file
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
	Comment string   `xml:",comment"`
}

// a simple struct which contains all our
// social links
type Social struct {
	XMLName   xml.Name `xml:"social"`
	Facebook  string   `xml:"facebook"`
	Twitter   string   `xml:"twitter"`
	Youtube   string   `xml:"youtube"`
	Instagram string
}

func main() {

	byteValue := `
<?xml version="1.0" encoding="UTF-8"?>
<users>
  <user type="admin">
    <name>Elliot</name>
    <social>
      <facebook>https://facebook.com</facebook>
      <twitter>https://twitter.com</twitter>
      <youtube>https://youtube.com</youtube>
    </social>
  </user>
  <user type="reader">
    <name>Fraser</name>
    <social>
      <facebook>https://facebook.com</facebook>
      <twitter>https://twitter.com</twitter>
      <youtube>https://youtube.com</youtube>
	  <Instagram>https://instagram.com</Instagram>
    </social>
  </user>
</users>
	`

	// we initialize our Users array
	var users Users
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal([]byte(byteValue), &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Instagram Url: " + users.Users[i].Social.Instagram)
		users.Users[i].Social.Instagram = "@instagram"
		users.Users[i].Comment = "This is comment"
	}
	fmt.Println()

	data, err := xml.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Printf("XML marshal error: %v\n", err)
		return
	}
	fmt.Printf("XML: %s\n", data)
}
