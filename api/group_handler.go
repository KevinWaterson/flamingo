package main

import(
	"fmt"
	"log"
	"net/url"
	"net/http"
	"encoding/json"
)

type UserGroup struct {
        Id              string
        Name            string
        Description     string
}


func fetchAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")

	jsonMsg, err := getGroups(w, r)
	if err != nil {
		http.Error(w, "Oops", http.StatusInternalServerError)
	}
	fmt.Print("Json message is ")
	fmt.Fprintf(w, jsonMsg)
	fmt.Print("\n")
}


func getGroups(w http.ResponseWriter, r *http.Request) (string, error){

	query_params, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Print(query_params)
	fmt.Print("\n")

	fmt.Print("API: fetching all groups\n")

	// open admin db
	dba:=OpenDB()
	sql := "SELECT * FROM user_group;"
	fmt.Print(sql+"\n")

	rows, err := dba.Query(sql)
	if err != nil {
		fmt.Print("DB Error\n")
		fmt.Print(err)
		return "", err
		fmt.Print("\n")
	}

	// set a counter for the rows..
	rowCount := 0

	// a slice to for the result sets
	var rs []UserGroup

	for rows.Next() {
		// a container for
		ug := UserGroup{}

		if err := rows.Scan(&ug.Id, &ug.Name, &ug.Description); err != nil {
			fmt.Print(err)
		}
		rs = append(rs, ug)
		fmt.Print(ug.Name)
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


func groupAddHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")

	// lets check the api key is valid..
	queryParams, _ := url.ParseQuery(r.URL.RawQuery)
	apiKey, ok := queryParams["api_key"]
	log.Println(apiKey)
	if ok && apiKey[0] == API_KEY{
		// jsonMsg, err := userAdd(w, r)
		// fmt.Print(jsonMsg)
		// fmt.Fprintf(w, jsonMsg)
		ok, _ := groupAdd(w,r)
		if ok != true {
			http.Error(w, "Oops", http.StatusInternalServerError)
		}
	} else {
		log.Println("Warning: Invalid API KEY. Possible Break In Attempt")
		jsonMsg := ""
		// fmt.Print(jsonMsg)
		fmt.Fprintf(w, jsonMsg)
	}
}

func groupAdd(w http.ResponseWriter, r *http.Request) (bool, error){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")

	Name := r.FormValue("name")
	Description := r.FormValue("description")

	sql := "INSERT INTO user_group(name, description) VALUES (?,?)"
	dba := OpenDB()
	stmt, err := dba.Prepare(sql)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(501)
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec( Name, Description)
	// get affected rows
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(501)
		return false, err
	}
	affected, err := res.RowsAffected()
	fmt.Println("Affect Rows is: ")
	fmt.Println(affected)
	w.WriteHeader(201)
	return true, nil
}

