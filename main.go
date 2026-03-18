package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/S-L-T/go-assessment/helper"
	"github.com/S-L-T/go-assessment/internal/adapter/http"
	"github.com/S-L-T/go-assessment/internal/core/use_case"
	"github.com/S-L-T/go-assessment/internal/repository/company"
	"log"
	netHTTP "net/http"
	"time"
)

func main() {
	err := helper.InitializeLogger(helper.TraceLevel)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		helper.Log(err, helper.FatalLevel)
		return
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	r, err := company.NewMySQL()
	if err != nil {
		helper.Log(err, helper.FatalLevel)
		return
	}
	startHTTPServer(http.NewServer(use_case.NewCompanyUseCase(ctx, r)))
}

func startHTTPServer(s *http.Server) {
	err := netHTTP.ListenAndServe(":8080", s.Router)

	if err != nil {
		helper.Log(err, helper.FatalLevel)
	}
}
