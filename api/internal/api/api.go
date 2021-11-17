package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Spuxy/CleanArchitecture/api/pkg/database"
	"github.com/Spuxy/CleanArchitecture/api/pkg/logger"
)

const port int = 9999

func Serve() {
	logger := logger.NewLogger()

	addr := os.Getenv("ADDR_PORT")
	database, err := database.NewPostgres()
	if err != nil {
		logger.LogError(err.Error())
	}
	defer database.Db.Close()
	handler := NewHandler(logger, database)
	router := NewRouter(handler)
	router.Routes()
	http.ListenAndServe(fmt.Sprintf(":%s", addr), nil)

}
