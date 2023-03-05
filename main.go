package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"

	webScrapper "web-backend/packages"
)

type Post struct {
	Title        string `json: "title"`
	Genre        string `json: "genre"`
	Director     string `json: "director"`
	Release_year int    `json: "release_year"`
}

var router *chi.Mux
var db *sql.DB

const (
	dbName = "movies"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "3306"
)

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbSource)
	catch(err)
}

func routers() *chi.Mux {
	router.Post("/movies/create", CreateMovie)
	return router
}

// CreatePost create a new post
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Insert movies SET title=?, genre=?, director=?, release_year=?")
	catch(err)

	_, er := query.Exec(post.Title, post.Genre, post.Director, post.Release_year)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func main() {
	go webScrapper.Crawl()
	go routers()
	http.ListenAndServe(":8005", Logger())
}
