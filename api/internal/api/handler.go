package api

import (
	"net/http"

	"github.com/Spuxy/CleanArchitecture/api/pkg/database"
	"github.com/Spuxy/CleanArchitecture/api/pkg/logger"
)

type Handler struct {
	log logger.Loggerer
	db  *database.Postgres
}

func NewHandler(logger logger.Loggerer, database *database.Postgres) *Handler {
	return &Handler{logger, database}
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	h.log.LogInfo("ww")
}
