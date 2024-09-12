package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	//set json type
	//access to params
	//loop over the movies
	//delete the movie of the id
	//add a new movie with the id

	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = param["id"]

			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)

			return

		}
	}

}

func handleHello(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, "Hello")

	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "32435", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "44324", Title: "Movie Two", Director: &Director{Firstname: "Razique", Lastname: "MK"}})

	movies = append(movies, Movie{ID: "3", Isbn: "43244", Title: "Movie Three", Director: &Director{Firstname: "Muhammed", Lastname: "Muthalib"}})
	r.HandleFunc("/", handleHello).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8000")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}

}
