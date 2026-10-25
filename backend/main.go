package main

import (
	"encoding/json"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := NewRunRepository(db)
	handler := NewRunHandler(repository)

	router := gin.New()

	// Collect HTTP request metrics.
	router.Use(PrometheusMiddleware())

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		level := "info"
		if param.StatusCode >= 500 {
			level = "error"
		} else if param.StatusCode >= 400 {
			level = "warn"
		}

		entry := map[string]interface{}{
			"timestamp": param.TimeStamp.Format(time.RFC3339Nano),
			"level":     level,
			"message":   "HTTP request",
			"service": map[string]string{
				"name": "runlog-backend",
			},
			"method":     param.Method,
			"path":       param.Path,
			"status":     param.StatusCode,
			"latency_ms": float64(param.Latency.Microseconds()) / 1000,
			"client_ip":  param.ClientIP,
		}

		if param.ErrorMessage != "" {
			entry["error"] = param.ErrorMessage
		}

		data, err := json.Marshal(entry)
		if err != nil {
			return "{}\n"
		}
		return string(data) + "\n"
	}))

	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://192.168.49.2:30080",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false,
	}))

	router.GET("/metrics", MetricsHandler())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/api/runs", handler.GetRuns)
	router.POST("/api/runs", handler.CreateRun)
	router.GET("/api/runs/:id", handler.GetRun)
	router.DELETE("/api/runs/:id", handler.DeleteRun)
	router.PUT("/api/runs/:id", handler.UpdateRun)

	router.Run(":8080")
}
