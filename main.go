package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	operation := pathParts[1]
	aStr := pathParts[2]
	bStr := pathParts[3]
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	fmt.Fprintf(w, "Operation:"+operation+"\n")
	fmt.Fprintf(w, "a:"+aStr+"\n")
	fmt.Fprintf(w, "b:"+bStr+"\n")
	// fmt.Fprintf(w, r.URL.Path+"\n")
	result := 0
	if operation == "add" {
		result = Add(a, b)
		fmt.Fprintf(w, aStr+" + "+bStr)
	} else if operation == "sub" {
		result = Sub(a, b)
		fmt.Fprintf(w, aStr+" - "+bStr)
	} else if operation == "mul" {
		result = Mul(a, b)
		fmt.Fprintf(w, aStr+" * "+bStr)
	} else if operation == "div" {
		result, _ = Div(a, b)
		fmt.Fprintf(w, aStr+" / "+bStr)
	} else {
		fmt.Fprintf(w, "Invaild argument!")
	}

	fmt.Fprintf(w, " = "+strconv.FormatInt(int64(result), 10))
	if operation == "div" {
		_, result = Div(a, b)
		fmt.Fprintf(w, ", remainder = "+strconv.FormatInt(int64(result), 10))
	}
}

func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
func Mul(a int, b int) int {
	return a * b
}
func Div(a, b int) (int, int) {
	return a / b, a - b*(a/b)
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
