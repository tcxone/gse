package gse

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	router     *http.ServeMux
	middleware []http.HandlerFunc
}

func new() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

func (s *Server) use(handler http.HandlerFunc) *Server {
	s.middleware = append(s.middleware, handler)
	return s
}

func (s *Server) get(pattern string, handler http.HandlerFunc) *Server {
	s.router.HandleFunc(pattern, s.wrapMiddleware(handler))
	return s
}

func (s *Server) post(pattern string, handler http.HandlerFunc) *Server {
	s.router.HandleFunc(pattern, s.wrapMiddleware(handler))
	return s
}

func (s *Server) listen(port string) {
	log.Fatal(http.ListenAndServe(":"+port, s.router))
}

func (s *Server) sendjson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) sendhtml(w http.ResponseWriter, statusCode int, templateFile string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) wrapMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, mw := range s.middleware {
			mw(w, r)
		}
		handler(w, r)
	}
}