package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	mux.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		zipCode := r.URL.Query().Get("zipCode")

		usecase := InitializeGetDataWithViaCepApiUseCase(&http.Client{})
		service := InitializeGetCityAndWeatherByZipCode(usecase)

		cityName := service.Execute(r.Context(), zipCode)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cityName))
	})

	log.Fatal(srv.ListenAndServe())
}
