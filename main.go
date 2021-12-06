package main

import (
	//"fmt"
	"log"
	"net/http"
  "os"

	//dependancy to load .env files
	"github.com/joho/godotenv"
)


func main() {
  err := godotenv.Load(".env")
  if(err != nil){
    log.Fatal("Error Loading .env File")
  }
  log.Println("listen on", os.Getenv("PORT"))
  log.Fatal(http.ListenAndServe(os.Getenv("PORT"),nil))
  

}
                                                                                                             
