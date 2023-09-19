package controllers

import (
	"context"
	"crud-go/internal/models"
	"crud-go/internal/repository"
	"crud-go/internal/responses"
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func GetStorage() []string {
	return []string{"Meat Lovers", "Bread Lovers", "PEPSI®", "DIET PEPSI®"}
}

func GetItems(ctx context.Context, repository repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var status int
		var message string

		res, e := repository.GetAllItems(ctx)
		if e != nil {
			status = http.StatusNotFound
			message = "error while fetching"
			log.Fatal("Error while fetching from db ", e)
		} else {
			status = http.StatusOK
			message = "success"
		}

		response := responses.ItemResponse{
			Status:  status,
			Message: message,
			Data:    map[string]interface{}{"data": res},
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
	}
}

func GetItemById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		itemId, _ := strconv.Atoi(params["id"])

		storage := GetStorage()

		var result models.Item
		status := http.StatusNotFound
		message := "not_found"

		for i, name := range storage {
			if i == itemId {
				result = models.Item{
					Id:   i,
					Name: name,
				}
				status = http.StatusOK
				message = "success"
			}
		}

		response := responses.ItemResponse{
			Status:  status,
			Message: message,
			Data:    map[string]interface{}{"data": result},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
