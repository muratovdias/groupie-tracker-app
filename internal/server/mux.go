package server

import (
	"net/http"
	"tracker/internal/handlers"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Route() *http.ServeMux {
	s.mux.HandleFunc("/", handlers.HomePage)
	s.mux.HandleFunc("/artists/", handlers.ArtistPage)
	s.mux.HandleFunc("/searched/", handlers.SearchHandler)
	s.mux.HandleFunc("/filter/", handlers.Filter)
	s.mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))

	return s.mux
}
