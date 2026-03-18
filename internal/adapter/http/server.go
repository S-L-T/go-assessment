package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/S-L-T/go-assessment/helper"
	"github.com/S-L-T/go-assessment/internal/core/domain"
	"github.com/S-L-T/go-assessment/internal/core/port"
	netHTTP "net/http"
	"os"
	"strings"
	"time"
)

type Server struct {
	Router         *mux.Router
	CompanyUseCase port.CompanyUseCase
}

func NewServer(c port.CompanyUseCase) *Server {
	s := Server{
		Router:         mux.NewRouter(),
		CompanyUseCase: c,
	}

	s.initializeRoutes()
	return &s
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/company", s.company).Methods(
		netHTTP.MethodGet,
		netHTTP.MethodPut,
		netHTTP.MethodPatch,
		netHTTP.MethodDelete,
		netHTTP.MethodOptions,
	)
	s.Router.HandleFunc("/healthcheck", s.healthcheck).Methods(
		netHTTP.MethodGet,
		netHTTP.MethodOptions,
	)
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
}

func (s *Server) writeResponse(w netHTTP.ResponseWriter, resData interface{}) {
	res, err := json.Marshal(resData)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(netHTTP.StatusInternalServerError)
	}

	_, err = w.Write(res)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(netHTTP.StatusInternalServerError)
	}
}

func (s *Server) company(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case netHTTP.MethodGet:
		s.getHandler(w, r)
		break
	case netHTTP.MethodPut:
		s.putHandler(w, r)
		break
	case netHTTP.MethodPatch:
		s.patchHandler(w, r)
		break
	case netHTTP.MethodDelete:
		s.deleteHandler(w, r)
		break
	case netHTTP.MethodOptions:
		s.optionsHandler(w)
		break
	default:
		w.WriteHeader(netHTTP.StatusMethodNotAllowed)
	}
}

func getBearerToken(r *netHTTP.Request) (string, error) {
	h := r.Header.Get("Authorization")
	if h == "" {
		return "", errors.New("no Authorization header")
	}

	split := strings.Split(h, "Bearer ")
	if len(split) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	return split[1], nil
}

func (s *Server) isAuthenticated(w netHTTP.ResponseWriter, r *netHTTP.Request) bool {
	t, err := getBearerToken(r)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return false
	}

	err = helper.IsAuthorized(t, os.Getenv("JWT_KEY"))
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return false
	}

	return true
}

func (s *Server) getHandler(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	req := GetReqAdapter{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	company, err := s.CompanyUseCase.Get(id)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(netHTTP.StatusInternalServerError)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	res := NewGetResAdapter(company)
	w.WriteHeader(netHTTP.StatusOK)
	s.writeResponse(w, res)
}

func (s *Server) putHandler(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	if !s.isAuthenticated(w, r) {
		return
	}

	req := PutReqAdapter{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	id, err := s.CompanyUseCase.Create(
		req.Name,
		req.Description,
		req.TotalEmployees,
		req.IsRegistered,
		domain.CompanyType(req.Type),
	)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(netHTTP.StatusInternalServerError)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}
	resData := PutResAdapter{
		ID: id.String(),
	}
	w.WriteHeader(netHTTP.StatusCreated)
	s.writeResponse(w, resData)

}

func (s *Server) patchHandler(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	if !s.isAuthenticated(w, r) {
		return
	}

	req := PatchReqAdapter{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	err = s.CompanyUseCase.Update(
		id,
		req.Name,
		req.Description,
		req.TotalEmployees,
		req.IsRegistered,
		domain.CompanyType(req.Type),
	)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(netHTTP.StatusInternalServerError)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	w.WriteHeader(netHTTP.StatusOK)
}

func (s *Server) deleteHandler(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	if !s.isAuthenticated(w, r) {
		return
	}
	
	req := DeleteReqAdapter{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		w.WriteHeader(netHTTP.StatusBadRequest)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}
	err = s.CompanyUseCase.Delete(id)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(netHTTP.StatusInternalServerError)
		resData := ErrorRes{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	w.WriteHeader(netHTTP.StatusOK)
}

func (s *Server) optionsHandler(w netHTTP.ResponseWriter) {
	w.Header().Add("Allow", "GET,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,PATCH,DELETE,OPTIONS")
	w.WriteHeader(netHTTP.StatusOK)
}

func (s *Server) healthcheck(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !s.CompanyUseCase.IsAlive(ctx) {
		w.WriteHeader(netHTTP.StatusFailedDependency)
		return
	}

	w.WriteHeader(netHTTP.StatusOK)
}
