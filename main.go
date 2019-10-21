package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "strconv"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("views")))
  http.HandleFunc("/leap", CheckLeap)
  fmt.Println("Listening to port 9797...")
  http.ListenAndServe(":9797", nil)
}


// Handlers
/*#####################################################*/
func CheckLeap (res http.ResponseWriter, req *http.Request) {
  type reqStruct struct { Year string `json: year`}
  var data reqStruct
  GetBodyData(req, &data)

  year, convErr := strconv.Atoi(data.Year)
  CheckError(convErr)
  result := "Not a leap year"

  if year % 4 == 0 {
    if year % 100 != 0 {
      result = "Leap year"
    } else {
      if year % 400 == 0 {
        result = "Leap year"
      }
    }
  }

  fmt.Fprintf(res, result)
}


/*#####################################################*/


// Helper Functions
/*#####################################################*/
func CheckError(err error) {
  if err != nil {
		panic(err)
	}
}

func GetBodyData(req *http.Request, class interface{}) {
  err := json.NewDecoder(req.Body).Decode(&class)
	CheckError(err)
}
/*#####################################################*/
