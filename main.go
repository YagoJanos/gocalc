package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/sum/{number1}/{number2}", sum).Methods("GET")
	router.HandleFunc("/sub/{number1}/{number2}", sub).Methods("GET")
	router.HandleFunc("/mult/{number1}/{number2}", mult).Methods("GET")
	router.HandleFunc("/div/{number1}/{number2}", div).Methods("GET")
	router.HandleFunc("/history", getHistory).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func sum(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	firstParam := vars["number1"]

	secondParam := vars["number2"]

	firstNumber, err1 := strconv.ParseFloat(firstParam, 64)
	secondNumber, err2 := strconv.ParseFloat(secondParam, 64)

	if err1 != nil || err2 != nil {
		fmt.Fprint(w, "Invalid argument")

	} else {

		var result float64 = firstNumber + secondNumber
		resultStr := fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resultStr)
	}
	saveLog(firstParam, "+", secondParam)
}

func sub(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	firstParam := vars["number1"]

	secondParam := vars["number2"]

	firstNumber, err1 := strconv.ParseFloat(firstParam, 64)
	secondNumber, err2 := strconv.ParseFloat(secondParam, 64)

	if err1 != nil || err2 != nil {
		fmt.Fprint(w, "Invalid argument")

	} else {

		var result float64 = firstNumber - secondNumber
		resultStr := fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resultStr)
	}
	saveLog(firstParam, "-", secondParam)
}

func mult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	firstParam := vars["number1"]

	secondParam := vars["number2"]

	firstNumber, err1 := strconv.ParseFloat(firstParam, 64)
	secondNumber, err2 := strconv.ParseFloat(secondParam, 64)

	if err1 != nil || err2 != nil {
		fmt.Fprint(w, "Invalid argument")
	} else {

		var result float64 = firstNumber * secondNumber
		resultStr := fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resultStr)
	}
	saveLog(firstParam, "*", secondParam)
}

func div(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	firstParam := vars["number1"]

	secondParam := vars["number2"]

	firstNumber, err1 := strconv.ParseFloat(firstParam, 64)
	secondNumber, err2 := strconv.ParseFloat(secondParam, 64)

	if err1 != nil || err2 != nil {
		fmt.Fprint(w, "Invalid argument")

	} else if secondNumber == 0 {
		fmt.Fprint(w, "Division by zero")
	} else {

		var result float64 = firstNumber / secondNumber
		resultStr := fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resultStr)
	}
	saveLog(firstParam, "/", secondParam)
}

type node struct {
	data string
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

var operationsList linkedList = linkedList{}

func saveLog(firstNumber string, operation string, secondNumber string) {

	fullOperation := firstNumber + " " + operation + " " + secondNumber
	node1 := &node{data: fullOperation}
	operationsList.prepend(node1)
}

func getHistory(w http.ResponseWriter, r *http.Request) {
	operationsList.printListData(w)
}

func (l linkedList) printListData(w http.ResponseWriter) {
	toPrint := l.head
	for l.length != 0 {
		fmt.Fprintln(w, toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("\n")
}
