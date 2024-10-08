package main

import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"

)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"secondname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main(){
	// fmt.Println("first commit")

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "John",Lastname: "William"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie two", Director: &Director{Firstname: "Steve", Lastname:"Silby"}})
	r := mux.NewRouter()
    r.HandleFunc("/movies", getMovies).Methods("GET")
    r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	
	fmt.Printf("Starting the server at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))



}