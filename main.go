package main 

import (
    "fmt"
	 "log"
	 "net/http"
	 "encoding/json"
	 "github.com/gorilla/mux"
)



type Article  struct{
  Title string `json:"Title"`
  Dec string  `json: "Descrption"`
  Content string `jsn:"Content"`
}

type Articles [] Article

func allArticles(w http.ResponseWriter , r *http.Request){
	articles := Articles{Article{Title: "My name is Namam", Dec : "Description" , Content: "Content is very good"}}
	fmt.Println("Hits the end point for arcticles" )
	json.NewEncoder(w).Encode(articles)

}

func testPostArticles(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "testPostArticles:-----")	

}

func homePage(w http.ResponseWriter , r *http.Request ){
fmt.Fprintf(w, "HomePage Hit:-----")
}


func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/" , homePage )
	myRouter.HandleFunc("/articles" , allArticles).Methods("GET")
	myRouter.HandleFunc("/articles" , testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081",myRouter))

}

func main(){
     handleRequests() 
}