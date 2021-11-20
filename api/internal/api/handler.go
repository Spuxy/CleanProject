package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Spuxy/CleanArchitecture/pkg/database"
	"github.com/Spuxy/CleanArchitecture/pkg/logger"
)

type Handler struct {
	log logger.Loggerer
	db  *database.Postgres
}

type User struct {
	name string
	age  int
}

func NewHandler(logger logger.Loggerer, database *database.Postgres) *Handler {
	return &Handler{logger, database}
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	err := h.log.LogInfo("Yop")
	if err != nil {
		h.log.LogError(err.Error())
	}
}

func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	return
	// fmt.Fprint(w, data)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	err := h.log.LogInfo("Yop")
	if err != nil {
		h.log.LogError(err.Error())
	}
}
