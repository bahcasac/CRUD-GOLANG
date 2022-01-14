package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-banco-de-dados/database"
	"io/ioutil"
	"net/http"
	"strconv"
)

type user struct {
	ID    uint32 `json:id`
	Name  string `json:name`
	Email string `json:email`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read body!"))
		return
	}
	var user user
	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error to convert user to struct"))
		return
	}
	db, err := database.Connection()
	if err != nil {
		w.Write([]byte("Error to connect to database"))

	}
	defer db.Close()
	//PREPARE STATEMENT - to avoid sql injection
	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Error to create the statement"))
		return
	}
	defer statement.Close()
	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Error to insert user"))
		return
	}
	userId, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Error to get user id"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Success to insert user! %d", userId)))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connection()
	if err != nil {
		w.Write([]byte("Erro to connect to database"))
		return
	}
	defer db.Close()
	lines, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Erro to search users"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user
		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error to scan the user"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error to encode users"))
		return
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	ID, err := getUserId(w, r)
	db, err := database.Connection()
	if err != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}
	defer db.Close()
	line, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Error to search user"))
		return
	}
	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error to scan user"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error to encode user"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ID, err := getUserId(w, r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read body"))
		return
	}
	var user user
	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error to convert user to struct"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("update users set name = ?, email = ?  where id = ?")
	if err != nil {
		w.Write([]byte("Error to create the statement"))
		return
	}
	defer statement.Close()
	updateUser, err := statement.Exec(user.Name, user.Email, ID)
	if err != nil {
		w.Write([]byte("Error to update user"))
		return
	}
	if _, err := updateUser.LastInsertId(); err != nil {
		w.Write([]byte("Error to get user id"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Success to update user! %d", ID)))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ID, err := getUserId(w, r)
	db, err := database.Connection()
	if err != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Error to create the statement"))
		return
	}
	defer statement.Close()
	deleteUser, err := statement.Exec(ID)
	if err != nil {
		w.Write([]byte("Error to delete user"))
		return
	}
	if _, err := deleteUser.LastInsertId(); err != nil {
		w.Write([]byte("Error to get user id"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Success to delete user! %d", ID)))
}

func getUserId(w http.ResponseWriter, r *http.Request) (uint64, error) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to get user id param"))
	}
	return ID, err
}
