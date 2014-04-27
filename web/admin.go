package main

import (
	"fmt"
	"net/http"
	// "io/ioutil"
	// "crypto/tls"
	"encoding/json"
)

type adminPage struct {
	Title	   string
	HostName	string
	MenuActive      string
	Username	string
	Email	   string
	Password	string
	Cpassword       string
	Protocol	string
	Groups	  []Group
	Users	   []User
	Errors map[string]string
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

	// get the user json content for the user tab
	u, err := getUsers(w, r)
	if err != nil{
		fmt.Print(err)
	}
	// a users slice
	users := []User{};
	errs := json.Unmarshal([]byte(u), &users)

	fmt.Print("user_admin.go: ")
	if errs == nil {
		fmt.Print("No errors: ")
		fmt.Printf("%+v\n", users)
	} else {
		fmt.Print("Errors here: ")
		fmt.Println(errs)
		fmt.Printf("%+v\n", users)
	}
	fmt.Print("\n")

	// get the groups json content for the group tab in the template (calls user_group.go)
	g, err := getUserGroups(w, r)
	if err != nil{
		fmt.Print(err)
	}
	// a groups slice
	groups := []Group{};
	gErrs := json.Unmarshal([]byte(g), &groups)

	if errs == nil {
		fmt.Printf("%+v\n", groups)
	} else {
		fmt.Println(gErrs)
		fmt.Printf("%+v\n", groups)
	}

	p := adminPage{Title:TITLE+" - Admin", HostName:r.Host, Protocol:SERVER_PROTOCOL, Groups:groups, Users:users}
	render(w, "admin", p)
}
