package main

import(
	"fmt"
	"log"
	"net/url"
	"net/http"
	"encoding/json"
	_ "github.com/lib/pq"
)

type User struct {
        Id              string
        Email           string
}

func fetchAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")

	jsonMsg, err := getUsers(w, r)
	if err != nil {
		http.Error(w, "Oops", http.StatusInternalServerError)
	}
	fmt.Print("Json message is ")
	fmt.Fprintf(w, jsonMsg)
	fmt.Print("\n")
}

func getUsers(w http.ResponseWriter, r *http.Request) (string, error){

	query_params, _ := url.ParseQuery(r.URL.RawQuery)
	log.Println(query_params)

	log.Println("API: fetching all groups")

	dba:=OpenDB()
	sql := "SELECT id, email FROM users;"
	log.Println(sql)

	rows, err := dba.Query(sql)
	if err != nil {
		log.Println("DB Error")
		log.Println(err)
		return "", err
	}

	// set a counter for the rows..
	rowCount := 0

	// a slice to for the result sets
	var rs []User
	for rows.Next() {
		// a container for
		u := User{}

		if err := rows.Scan(&u.Id, &u.Email); err != nil {
			fmt.Print(err)
		}
		rs = append(rs, u)
		fmt.Print(u.Email)
		fmt.Print("\n")
		rowCount++
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Print("Prepare JSON\n")

	jbMsg, err := json.Marshal(rs)
	jsonMsg := string(jbMsg[:]) // converting byte array to string
	return jsonMsg, nil
}

