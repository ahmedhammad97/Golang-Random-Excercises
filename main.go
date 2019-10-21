package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "strconv"
  "time"
  "math"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("views")))

  // Routes
  http.HandleFunc("/leap", CheckLeap)
  http.HandleFunc("/lived", CalculateLivedMoment)
  http.HandleFunc("/squares", SumOfSquares)
  http.HandleFunc("/multiples", SumOfMultiples)
  http.HandleFunc("/binary", ConvertToBinary)
  http.HandleFunc("/palindrom", CheckPalindrom)
  http.HandleFunc("/prime", PrimeFactors)
  http.HandleFunc("/search", BinarySearch)

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

func CalculateLivedMoment(res http.ResponseWriter, req *http.Request){
  type reqStruct struct { Time string `json: time`}
  var data reqStruct
  GetBodyData(req, &data)

  t, err := time.Parse(time.RFC3339, data.Time)
  CheckError(err)

	tEpoch := float64(t.Unix())
	gigaSecond := math.Pow(10.0, 9)
	tPlusGiga := tEpoch + gigaSecond
	t = time.Unix(int64(tPlusGiga), 0)

  fmt.Fprintf(res, t.String())
}

func SumOfSquares(res http.ResponseWriter, req *http.Request){

}

func SumOfMultiples(res http.ResponseWriter, req *http.Request){

}

func ConvertToBinary(res http.ResponseWriter, req *http.Request){

}

func CheckPalindrom(res http.ResponseWriter, req *http.Request){

}

func PrimeFactors(res http.ResponseWriter, req *http.Request){

}

func BinarySearch(res http.ResponseWriter, req *http.Request){

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
