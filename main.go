package main

import (
	"context"
	"golang-chap47/infra"
	"golang-chap47/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	ctx.Cron.AddFunc("@every 1m", func() {
		log.Println("Running cron job to export order reports")
		err := ctx.Ctl.Order.ExportOrderReports("order_report.xlsx")
		if err != nil {
			ctx.Log.Error("Failed to export order reports", zap.Error(err))
		} else {
			log.Println("Order reports exported successfully")
		}
	})

	// Start the cron scheduler
	ctx.Cron.Start()
	defer ctx.Cron.Stop()

	r := routes.NewRoutes(*ctx)

	port := ctx.Cfg.Port
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		// Start the server
		log.Printf("Server running on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// Catching context timeout
	select {
	case <-shutdownCtx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
