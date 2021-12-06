package main

import (
	//"fmt"
	"net/http"
  "log"
)


func main() {
  const PORT string = ":9090"
  http.ListenAndServe(PORT, nil)
  log.Println("listen on", PORT)
  log.Fatal(http.ListenAndServe(PORT,nil))
  

}
                                                                                                             
