package cli

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/rysmaadit/go-template/model"
	"github.com/rysmaadit/go-template/repository"

	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/router"
	"github.com/rysmaadit/go-template/service"
	log "github.com/sirupsen/logrus"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (c *Cli) Run(application *app.Application) {
	log.SetLevel(log.InfoLevel)
	log.StandardLogger()
	log.SetOutput(os.Stdout)

	if strings.ToLower(application.Config.LogLevel) == log.DebugLevel.String() {
		log.SetLevel(log.DebugLevel)
	}

	log.SetReportCaller(true)

	// butuh ini untuk migrasi
	db, err := repository.Connect(application.Config.DBDSN)
	if err != nil {
		panic(err)
	}

	//migration
	if len(c.Args) >= 2 && c.Args[1] == "migrate" {
		db.Migrator().AutoMigrate(&model.Movie{})
		fmt.Println("Migration run successfully")
		os.Exit(0)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", application.Config.AppPort),
		Handler: router.NewRouter(service.InstantiateDependencies(application), db),
	}

	log.Println(fmt.Sprintf("starting application { %v } on port :%v", application.Config.AppName, application.Config.AppPort))

	go listenAndServe(srv)
	waitForShutdown(srv)
}

func listenAndServe(apiServer *http.Server) {
	err := apiServer.ListenAndServe()

	if err != nil {
		log.WithField("error", err.Error()).Fatal("unable to serve")
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)

	_ = <-sig

	log.Warn("shutting down")

	if err := apiServer.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}

	log.Warn("shutdown complete")
}
