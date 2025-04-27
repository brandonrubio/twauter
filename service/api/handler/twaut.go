package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brandonrubio/twauter/service/logger"
	"go.uber.org/dig"
)

type TwautResponseBody struct {
	Message string `json:"message"`
}

type ITwautHandlerService interface {
	HandleGetAll(w http.ResponseWriter, r *http.Request)
	HandleGetOne(w http.ResponseWriter, r *http.Request)
	HandlePost(w http.ResponseWriter, r *http.Request)
	HandlePut(w http.ResponseWriter, r *http.Request)
	HandleDelete(w http.ResponseWriter, r *http.Request)
}

type TwautHandlerService struct {
	loggerService logger.ILoggerService
}

func (ths *TwautHandlerService) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	ths.loggerService.Log("info", "sending response for twaut GET request", "TwautHandlerService.HandleGetAll")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(TwautResponseBody{Message: "GET twaut response"})
	if err != nil {
		ths.loggerService.Log("error", "failed to return response body", "TwautHandlerService.HandleGetAll")
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

func (ths *TwautHandlerService) HandleGetOne(w http.ResponseWriter, r *http.Request) {
	ths.loggerService.Log("info", "sending response for twaut GET request", "TwautHandlerService.HandleGetOne")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(TwautResponseBody{Message: "GET twaut response"})
	if err != nil {
		ths.loggerService.Log("error", "failed to return response body", "TwautHandlerService.HandleGetOne")
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

func (ths *TwautHandlerService) HandlePost(w http.ResponseWriter, r *http.Request) {
	ths.loggerService.Log("info", "sending response for twaut POST request", "TwautHandlerService.HandlePost")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(TwautResponseBody{Message: "POST twaut response"})
	if err != nil {
		ths.loggerService.Log("error", "failed to return response body", "TwautHandlerService.HandlePost")
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

func (ths *TwautHandlerService) HandlePut(w http.ResponseWriter, r *http.Request) {
	ths.loggerService.Log("info", "sending response for twaut PUT request", "TwautHandlerService.HandlePut")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(TwautResponseBody{Message: "PUT twaut response"})
	if err != nil {
		ths.loggerService.Log("error", "failed to return response body", "TwautHandlerService.HandlePut")
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

func (ths *TwautHandlerService) HandleDelete(w http.ResponseWriter, r *http.Request) {
	ths.loggerService.Log("info", "sending response for twaut DELETE request", "TwautHandlerService.HandleDelete")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(TwautResponseBody{Message: "GET twaut response"})
	if err != nil {
		ths.loggerService.Log("error", "failed to return response body", "TwautHandlerService.HandleDelete")
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

type TwautHandlerServiceDependencies struct {
	dig.In
	LoggerService logger.ILoggerService `name:"LoggerService"`
}

func CreateTwautHandlerService(dependencies TwautHandlerServiceDependencies) *TwautHandlerService {
	return &TwautHandlerService{
		loggerService: dependencies.LoggerService,
	}
}
