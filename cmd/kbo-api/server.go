package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"net/http"

	"encoding/json"
	"time"

	"os/signal"

	"github.com/gorilla/mux"
	"github.com/seeeturtle/kbo"
)

var (
	logger *log.Logger
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	var (
		logName string
		logFile *os.File
		port    string
		err     error
	)

	flag.StringVar(&logName, "log", "kbo-api.log", "file name to log")
	flag.StringVar(&port, "port", "8080", "port to listen")
	flag.Parse()

	logFile, err = os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	logger = log.New(logFile, "kbo-api: ", log.LstdFlags|log.Lshortfile)

	r := mux.NewRouter()
	r.HandleFunc("/game", GameHandler).
		Methods("GET").
		Queries("year", "{year}").
		Queries("month", "{month}").
		Queries("day", "{day}")

	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,
	}

	go func() {
		logger.Printf("listening on http://0.0.0.0:%s\n", port)

		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-stop

	logger.Println("shutting down the server...")

	server.Shutdown(nil)

	logger.Println("server gracefully stopped")
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		http.Error(w, "Please provide proper query parameters", http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(vars["month"])
	if err != nil {
		http.Error(w, "Please provide proper query parameters", http.StatusBadRequest)
		return
	}

	day, err := strconv.Atoi(vars["day"])
	if err != nil {
		http.Error(w, "Please provide proper query parameters", http.StatusBadRequest)
		return
	}

	result, err := kbo.NewParser(kbo.URL, &http.Client{}).Parse(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))

	if err != nil {
		http.Error(w, "Error occured while parsing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
