package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/booomch/demo-golang/internal/app"
	"github.com/booomch/demo-golang/internal/appconfig"
	"github.com/booomch/demo-golang/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	contactrepo "github.com/booomch/demo-golang/internal/entities/contact/repositories"
	contactusecase "github.com/booomch/demo-golang/internal/entities/contact/usecase"

	userrepo "github.com/booomch/demo-golang/internal/entities/user/repositories"
	userusecase "github.com/booomch/demo-golang/internal/entities/user/usecase"
)

func main() {
	dispatcher()
	debug.SetPanicOnFault(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	// Initialize configuration service.
	appconfig.Init("configs/config.yaml")

	dsn := fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s sslmode=%s sslrootcert=%s",
		appconfig.Config.DB.Host,
		appconfig.Config.DB.Port,
		appconfig.Config.DB.Name,
		appconfig.Config.DB.User,
		appconfig.Config.DB.Password,
		appconfig.Config.DB.SslMode,
		appconfig.Config.DB.SslCert)
	master, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Fatal("master db connection", err)
	}
	defer master.Close()

	dbs := []*sqlx.DB{master}
	useruc := userusecase.New(userrepo.New(dbs))
	contactuc := contactusecase.New(contactrepo.New(dbs))

	// Setup application controller with its dependencies.
	ctr := controller.New(contactuc, useruc)

	// Configure and run fiber.
	a := fiber.New()
	app.Setup(a, ctr)
	addr := fmt.Sprintf(":%s", appconfig.Config.Port)
	log.Fatal(a.Listen(addr))
}

func dispatcher() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logrus.Info(fmt.Sprintf("Application gracefully shutdown at %s", os.Getenv("CURRENT_HOST_NAME")))
		select {
		case <-time.After(2 * time.Second):
			os.Exit(0)
			break
		}
	}()
}
