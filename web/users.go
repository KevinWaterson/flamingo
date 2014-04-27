package main

import(
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"crypto/tls"
)

type User struct {
	// Id      int    `json:"id"`
	Id      string    `json:"id"`
	Email string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) (string, error){

	log.Print("Fetching Users..\n")

	// set up a client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8080/api/user/fetchAll?api_key="+API_KEY)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print("Response is .. ")
	fmt.Print(resp)
	fmt.Print("\n")
	if err != nil{
		log.Print(err)
		return "", err
	}
	// convert response []byte array to string
	s := string(body)
	fmt.Print(s+"\n")
	return s, nil
}

