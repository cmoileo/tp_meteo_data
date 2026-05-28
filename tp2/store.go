package main

import "tp_meteo_data/tp1"

type Store struct {
	stations map[string]tp1.Station
}

func NewStore() *Store {
	return &Store{stations: make(map[string]tp1.Station)}
}

func (s *Store) Put(st tp1.Station) {
	s.stations[st.ID] = st
}

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok
}

func (s *Store) Get(id string) (tp1.Station, bool) {
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

func (s *Store) All() []tp1.Station {
	all := make([]tp1.Station, 0, len(s.stations))
	for _, st := range s.stations {
		all = append(all, st)
	}
	return all
}
