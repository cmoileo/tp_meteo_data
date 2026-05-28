package main

import (
	"encoding/json"
	"net/http"
	"tp_meteo_data/weather"
)

type createStationRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country_code"`
	Altitude int16  `json:"altitude"`
}

type App struct {
	store *Store
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	stations := a.store.All()
	writeJSON(w, http.StatusOK, stations)
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "station not found")
		return
	}
	writeJSON(w, http.StatusOK, st)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	var req createStationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if a.store.Has(req.ID) {
		writeError(w, http.StatusConflict, "station already exists")
		return
	}

	st := weather.Station{
		ID:       req.ID,
		Name:     req.Name,
		Country:  req.Country,
		Altitude: req.Altitude,
	}
	a.store.Put(st)
	writeJSON(w, http.StatusCreated, st)
}
