package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	//dependancy to load .env files
	"github.com/joho/godotenv"
)

type Data struct {
  ID   string `json:"id"`
  Name string `json:"name"`
  Username string `json:"username"`
}
type ResponseData struct {
  Data Data `json:"data"`
}


type DataTimeline struct {
	//ID   string `json:"id"`
	Text string `json:"text"`
}
type Meta struct {
	OldestID    string `json:"oldest_id"`
	NewestID    string `json:"newest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}
type TimeLineResponse struct {
	Data []DataTimeline `json:"data"`
	Meta Meta   `json:"meta"`
}

func main() {
  //Load env file
  err := godotenv.Load(".env")
  if(err != nil){
    log.Fatal("Error Loading .env File")
  }
  var user string
  fmt.Println("Enter Twitter username:")
  fmt.Scanf("%s",&user)

  var byte  = getApi(fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s",user),os.Getenv("TWITTER_BEARER_TOKEN"))
  var userData ResponseData
  errUnmarshall := json.Unmarshal(byte, &userData)
  if(errUnmarshall != nil){
    log.Fatal(errUnmarshall)
  }
  var tweetByte = getApi(fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets",string(userData.Data.ID)),os.Getenv("TWITTER_BEARER_TOKEN"))
  
  var timelineData TimeLineResponse
  errUnmarshall2 := json.Unmarshal(tweetByte,&timelineData)
  if(errUnmarshall2 != nil){
    log.Fatal(errUnmarshall2)
  }

  for _,v := range timelineData.Data{
    log.Printf("%s \n",string(v.Text))
  }

  getNextTweet(userData.Data.ID,timelineData.Meta.NextToken)

   






  //log.Println("server listening on port", os.Getenv("PORT"))
  //log.Fatal(http.ListenAndServe(os.Getenv("PORT"),nil))
  

}
func getApi( u string, token string)([]byte){
  client := &http.Client{}
  req,err := http.NewRequest("GET",u,nil)
  if(err != nil){
    log.Fatal(err)
  }
  req.Header = http.Header{
    "Content-Type": []string{"application/json"},
    "Authorization": []string{fmt.Sprintf("Bearer %s",token)},
  }

  res, error := client.Do(req)
  if(error != nil){
    log.Fatal(error)
  } 
  byte,err := io.ReadAll(res.Body)


  return byte


}
func getNextTweet(ID string,t string){
    var input string
    fmt.Println("Get next tweets? (next/exit):")
    fmt.Scanf("%s",&input)
    if(strings.ToLower(input) == "next" || strings.ToLower(input) == "exit"){
      if(strings.ToLower(input) == "next"){
        var nextPage TimeLineResponse
        var nextPageByte = getApi(fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets?pagination_token=%s",string(ID),string(t)),os.Getenv("TWITTER_BEARER_TOKEN"))
        errNextPage := json.Unmarshal(nextPageByte,&nextPage)
        if(errNextPage != nil){
          log.Fatal(errNextPage)
        }
        for _,v := range nextPage.Data{
          log.Printf("%s \n",string(v.Text))
        }
        getNextTweet(ID,nextPage.Meta.NextToken)

      } 
      if(strings.ToLower(input) == "exit"){
        return 
      }
    } 

}

