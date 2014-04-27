package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {

//	http.HandleFunc("/api/user/add", userAddHandler)
//	http.HandleFunc("/api/user/del", userDelHandler)
	http.HandleFunc("/api/user/fetchAll", fetchAllUsersHandler)

	http.HandleFunc("/api/group/fetchAll", fetchAllGroupsHandler)
	http.HandleFunc("/api/group/add", groupAddHandler)
//	http.HandleFunc("/api/user/group/add", userGroupAddHandler)	     // Adds a user to a group
//	http.HandleFunc("/api/user/group/del", userGroupDelHandler)	     // Deletes a user from a group
//	http.HandleFunc("/api/user/group/fetchAll", fetchAllUserGroupHandler)   // fetches all groups a user is a member of

	fmt.Println("Listening on 8080...")
	uErr := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if uErr != nil {
		log.Fatal(uErr)
	}
}

