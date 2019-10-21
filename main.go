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
  type reqStruct struct { Num string `json: num`}
  var data reqStruct
  GetBodyData(req, &data)

  N, convErr := strconv.Atoi(data.Num)
  CheckError(convErr)

  sum := 0
  squares := 0

  for i := 1; i<=N; i++ {
    sum += i
    squares += (i * i)
  }

  diff := Abs(squares - (sum * sum))
  fmt.Fprintf(res, strconv.Itoa(diff))
}

func SumOfMultiples(res http.ResponseWriter, req *http.Request){
  type reqStruct struct {
    Num string `json: num`
    N string `json N`
  }
  var data reqStruct
  GetBodyData(req, &data)

  num, convErr1 := strconv.Atoi(data.Num)
  CheckError(convErr1)
  N, convErr2 := strconv.Atoi(data.N)
  CheckError(convErr2)

  sum := 0
  for i := 1; i<N; i++ {
    sum += i * num
  }

  fmt.Fprintf(res, strconv.Itoa(sum))
}

func ConvertToBinary(res http.ResponseWriter, req *http.Request){
  type reqStruct struct { Num string `json: num`}
  var data reqStruct
  GetBodyData(req, &data)

  binaryStr := data.Num
  sum := 0.0
  length := len(binaryStr)

  for i := length -1; i>=0; i-- {
    if binaryStr[i] == 49 {
      sum += math.Pow(2, float64(length -i -1))
    }
  }

  fmt.Fprintf(res, strconv.Itoa(int(sum)))
}

func CheckPalindrom(res http.ResponseWriter, req *http.Request){
  type reqStruct struct { Str string `json: str`}
  var data reqStruct
  GetBodyData(req, &data)

  str := data.Str
  length := len(str)

  for i := length -1; i>=length/2; i-- {
    if str[i] != str[length -i -1] {
      fmt.Fprintf(res, "Not a palindrom")
      return
    }
  }
  fmt.Fprintf(res, "Palindrom")
}

func PrimeFactors(res http.ResponseWriter, req *http.Request){
  type reqStruct struct { Num string `json: num`}
  var data reqStruct
  GetBodyData(req, &data)

  num, convErr1 := strconv.Atoi(data.Num)
  CheckError(convErr1)
  result := make([]int, 0)

  for num%2 == 0 {
		result = append(result, 2)
		num = num / 2
	}

	for i := 3; i*i <= num; i = i + 2 {
		for num%i == 0 {
			result = append(result, i)
			num = num / i
		}
	}

	if num > 2 {
		result = append(result, num)
	}

  jsonArr, err := json.Marshal(result)
  CheckError(err)
  fmt.Fprintf(res, string(jsonArr))
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

func Abs(num int) int{
  if num > 0 {
    return num
  } else {
    return -num
  }
}
/*#####################################################*/
