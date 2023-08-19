package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var movies []Movie

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if (params["id"] == item.ID) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	fmt.Print(movie.ID)
	fmt.Print(movie.Title)
	fmt.Print(&movie.Director)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// params := mux.Vars(r)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if (params["id"] == item.ID) {
			movies = append(movies[:index], movies[index + 1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	// r.HandleFunc("/", HomeHandler)
	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Walking dead", Director: &Director{FirstName: "Quin", LastName: "Scorces"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4654", Title: "Breaking bad", Director: &Director{FirstName: "John", LastName: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "7654", Title: "Dexter", Director: &Director{FirstName: "Dexter", LastName: "Morgan"}})
	movies = append(movies, Movie{ID: "4", Isbn: "1287", Title: "Miami police", Director: &Director{FirstName: "Maria", LastName: "Laguerta"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Listening Server on port:8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}