package handlers

import (
	"BIOTRACKERSERVICE/internal/auth"
	"BIOTRACKERSERVICE/internal/usecases"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"net/http"
)

// RegisterUserHandler expects user information in json format
func (c *Controller) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var userInfo usecases.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		ProcessError(w, ErrIncorrectUserRegistrationInfo, http.StatusBadRequest)
		return
	}

	err = c.Usecases.DbManager.RegisterUser(r.Context(), userInfo)
	if errors.Is(err, usecases.ErrUserAlreadyExists) {
		ProcessError(w, err, http.StatusBadRequest)
		return
	}
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
		return
	}
}

// LoginHandler expects username and password in json format
func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials usecases.CredentialsDTO
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		ProcessError(w, err, http.StatusBadRequest)
		return
	}
	token, err := c.Usecases.UserToken(r.Context(), credentials)
	if errors.Is(err, auth.ErrUnauthorized) {
		ProcessError(w, err, http.StatusUnauthorized)
		return
	}
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	body := make(map[string]string)
	body["token"] = token
	bodyJSON, _ := json.Marshal(body)
	_, err = w.Write(bodyJSON)
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
}

// UserPlotsInfoHandler expects user_id in json format
func (c *Controller) UserPlotsInfoHandler(w http.ResponseWriter, r *http.Request) {
	var UserId struct {
		Id int32 `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&UserId)
	if err != nil {
		ProcessError(w, err, http.StatusBadRequest)
	}
	plotsInfo, err := c.Usecases.DbManager.GetUserPlotsInfo(r.Context(), UserId.Id)
	if errors.Is(err, usecases.ErrUserNotFound) || errors.Is(err, pgx.ErrNoRows) {
		ProcessError(w, err, http.StatusNotFound)
	}
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
	body, err := json.Marshal(plotsInfo)
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
		return
	}
}

// UserPlotsByIdsHandler expects plot ids in json format
func (c *Controller) UserPlotsByIdsHandler(w http.ResponseWriter, r *http.Request) {
	var plotIds struct {
		Ids []int32 `json:"ids"`
	}
	err := json.NewDecoder(r.Body).Decode(&plotIds)
	if err != nil {
		ProcessError(w, err, http.StatusBadRequest)
	}
	plots, err := c.Usecases.DbManager.GetPlotsByIds(r.Context(), plotIds.Ids)
	if errors.Is(err, pgx.ErrNoRows) {
		ProcessError(w, err, http.StatusNotFound)
	}
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
	body, err := json.Marshal(plots)
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
}

func (c *Controller) AddPlotHandler(w http.ResponseWriter, r *http.Request) {
	var plotInfo usecases.AddPlotDTO
	err := json.NewDecoder(r.Body).Decode(&plotInfo)
	if err != nil {
		ProcessError(w, err, http.StatusBadRequest)
	}
	err = c.Usecases.DbManager.AddPlot(r.Context(), plotInfo)
	if errors.Is(err, usecases.ErrUserNotFound) {
		ProcessError(w, err, http.StatusNotFound)
	}
	if err != nil {
		ProcessError(w, err, http.StatusInternalServerError)
	}
}

func ProcessError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	body := make(map[string]string)
	body["error"] = err.Error()
	bodyJSON, _ := json.Marshal(body)
	_, _ = w.Write(bodyJSON)
}
