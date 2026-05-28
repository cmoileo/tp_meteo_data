package main

import "tp_meteo_data/weather"

type Store struct {
	stations map[string]weather.Station
}

func NewStore() *Store {
	return &Store{stations: make(map[string]weather.Station)}
}

func (s *Store) Put(st weather.Station) {
	s.stations[st.ID] = st
}

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok
}

func (s *Store) Get(id string) (weather.Station, bool) {
	st, ok := s.stations[id]
	return st, ok
}

func (s *Store) Delete(id string) bool {
	if _, ok := s.stations[id]; !ok {
		return false
	}
	delete(s.stations, id)
	return true
}

func (s *Store) All() []weather.Station {
	all := make([]weather.Station, 0, len(s.stations))
	for _, st := range s.stations {
		all = append(all, st)
	}
	return all
}
