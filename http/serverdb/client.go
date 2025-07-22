package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User model
type user struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func FindUserById(w http.ResponseWriter, r *http.Request, id int) {
	var u user
	db, err := sql.Open("mysql", "root:123456@/curso_go")
	if err != nil {
		panic(err)
	}
	db.QueryRow("SELECT * from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@/curso_go")
	if err != nil {
		panic(err)
	}

	rows, _ := db.Query("SELECT * from usuarios")

	var usuarios []user
	for rows.Next() {
		var users user
		rows.Scan(&users.ID, &users.Nome)
		usuarios = append(usuarios, users)
	}
	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

// User handler analisa o request e delega para função adequada
func UserHandler(w http.ResponseWriter, r *http.Request) {
	//O que está acontencendo aqui?
	//essa sid pega tudo depois do /usuarios/ e pega convertendo pra int
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		FindUserById(w, r, id)
	case r.Method == "GET":
		FindAllUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found... :(")
	}

}
