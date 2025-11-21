package handler

import (
	"avito/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) AddTeam(w http.ResponseWriter, r *http.Request) {
	team := models.Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		jsonError(w, NotValid, "Not valid data", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	err = h.db.AddTeam(&team)
	if err != nil {
		jsonError(w, TeamExists, "team_name already exists", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)

}

func (h *Handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("team_name")
	if teamName == "" {
		jsonError(w, NotFound, "team_name not found", http.StatusNotFound)
		return
	}
	team, err := h.db.GetTeamByName(teamName)
	if err != nil {
		jsonError(w, NotFound, "team_name not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}
