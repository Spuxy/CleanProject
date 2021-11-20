package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Spuxy/CleanArchitecture/pkg/database"
	"github.com/Spuxy/CleanArchitecture/pkg/logger"
)

const port int = 9999

func Serve() {
	logger := logger.NewLogger()

	addr := os.Getenv("ADDR_PORT")
	database, err := database.NewPostgres()

	if err != nil {
		logger.LogError(err.Error())
		fmt.Println(err)
	}

	handler := NewHandler(logger, database)

	http.HandleFunc("/", handler.Welcome)
	http.HandleFunc("/users", handler.Users)

	err = http.ListenAndServe(fmt.Sprintf(":%s", addr), nil)
	fmt.Println(err)

}
