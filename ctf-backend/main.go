package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"ctf-backend/internal/db"
	"ctf-backend/internal/handlers"
	"ctf-backend/internal/middleware"

)

func main() {
	// ======================
	// LOAD ENV FILE
	// ======================
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// ======================
	// ENV CONFIG
	// ======================
	MONGO_URI := os.Getenv("MONGO_URI")
	DB_NAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("PORT")

	if MONGO_URI == "" || DB_NAME == "" || PORT == "" {
		log.Fatal("Missing required environment variables: MONGO_URI, DB_NAME, or PORT")
	}

	// ======================
	// DB CONNECT
	// ======================
	if err := db.ConnectMongo(MONGO_URI, DB_NAME); err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}

	// ======================
	// GIN SETUP
	// ======================
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ======================
	// ROUTES
	// ======================
	api := r.Group("/api")

	// AUTH
	api.POST("/team/login",middleware.RateLimiter(5.0/60.0, 5), handlers.LoginTeamHandler)
// signup
	api.POST("/auth/signup",middleware.RateLimiter(5.0/60.0, 5), handlers.Signup)
	// TEAM
	api.POST("/team/create",middleware.RateLimiter(5.0/60.0, 5), handlers.CreateTeam)
	api.POST("/team/join", middleware.RateLimiter(5.0/60.0, 5),handlers.JoinTeam)
	api.POST("/submit", middleware.RateLimiter(10.0/60.0, 3),middleware.TeamAuth(), handlers.SubmitChallengeHandler)
	api.GET("/challenges", middleware.RateLimiter(10.0/60.0, 3),handlers.GetChallenges)
	api.GET("/leaderboard",middleware.RateLimiter(10.0/60.0, 3), handlers.Leaderboard)

	api.Use(middleware.TeamAuth())
{
	api.GET("/team/me", handlers.GetMyTeam)
}

// routes.go

	// PROTECTED
	

	// ======================
	// SERVER
	// ======================
	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}

	go func() {
		log.Println("🚀 Server running on port", PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// ======================
	// GRACEFUL SHUTDOWN
	// ======================
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Forced shutdown:", err)
	}

	log.Println("✅ Server exited cleanly")
}