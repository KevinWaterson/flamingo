package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "crypto/tls"
        // "encoding/json"
)

type Group struct{
	Id      string
	Name    string
	Description string
}

// get a list of groups a user is a member of
func getUserGroups(w http.ResponseWriter, r *http.Request) (string, error){

	fmt.Print("Fetching User Groups..\n")

	// set up a client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	link := "https://localhost:8080/api/group/fetchAll?api_key="+API_KEY
	fmt.Println(link)
	resp, err := client.Get(link)
	// resp, err := client.Get("https://localhost:8080/api/group/fetchAll?api_key="+API_KEY)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print("Response is .. ")
	fmt.Print(resp)
	fmt.Print("\n")
	if err != nil{
		fmt.Print(err)
		return "", err
	}
	// convert response []byte array to string
	s := string(body)
	fmt.Print(s+"\n")
	return s, nil
}

