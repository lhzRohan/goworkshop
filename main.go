package main

import (
	"encoding/json" //here i include this beacuse i get data in json formate by postman
	"fmt"
	"github.com/gorilla/mux" //external import by gitub
	"log"                    // this is for error
	"math/rand"              //this is for add new movie and ceated id
	"net/http"
	"strconv"
)

type Movie struct { //struct means object type movie and director
	ID       string    `json:"id"`
	Rohan    string    `json:"rohan"` //rohan for assign unique id
	Title    string    `json:"Title"`
	Director *Director `json:"director"` //dirotor is a type of director struct
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie //type slice its a variable

func getMovies(w http.ResponseWriter, r *http.Request) { //w reponsewriter and r response here send the request to the  postman
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("connect-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode((movies))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("cannot-type", "application/json")
	params := mux.Vars(r)         //vars use for access all the thing by params
	for _, item := range movies { // _ (blank identifier) without _ this we get an error
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("contant-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-tpe", " application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter() //declare and define the fuction methods

	movies = append(movies, Movie{ID: "1", Rohan: "24555", Title: "Dark one", Director: &Director{Firstname: "john", Lastname: "Depth"}})
	movies = append(movies, Movie{ID: "2", Rohan: "24975", Title: "Dark two", Director: &Director{Firstname: "steave", Lastname: "smith"}})
	r.HandleFunc("/movies", getMovies).Methods("GET") //routes
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("starting server at port 4000\n")
	log.Fatal(http.ListenAndServe(":4000", r))

}
