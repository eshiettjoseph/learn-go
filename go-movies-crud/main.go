package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Structs are like templates for building. Sample struct below is Json format structure.
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
		movies = append(movies[:index], movies[index+1:]...)
		break
		}
	json.NewEncoder(w).Encode(movies)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
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
	
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "438227", Title: "Schindlers List", Director: &Director{
			Firstname: "Joseph", Lastname: "Eshiett"}})
	movies = append(movies, Movie{
		ID: "2", Isbn: "438228", Title: "Jesus of Nazareth", Director: &Director{
			Firstname: "Eshiett", Lastname: "Joseph"}})

	// Create movies route. Only Get requests allowed
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// Get movie via unique ID
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// Create movie
	r.HandleFunc("/movies/movies", createMovie).Methods("POST")
	// Update movie via ID
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// Delete movie
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	// Catch any errors that may occur when starting server
	log.Fatal(http.ListenAndServe(":8000", r))

}