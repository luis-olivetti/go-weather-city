package main

import (
	"encoding/json"
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

		viaCepUsecase := InitializeGetDataWithViaCepApiUseCase(&http.Client{})
		weatherUsecase := InitializeGetTemperatureWithWeatherApiUseCase(&http.Client{})
		service := InitializeGetCityAndWeatherByZipCode(viaCepUsecase, weatherUsecase)

		response, err, statusCode := service.Execute(r.Context(), zipCode)
		if err != nil {
			fmt.Println(err)

			if statusCode == 422 {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("invalid zipcode"))
				return
			}

			if statusCode == 404 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("cannot found zipcode"))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		jsonResponse, err := json.Marshal(response)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error - weather"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	log.Fatal(srv.ListenAndServe())
}
