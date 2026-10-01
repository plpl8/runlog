package main

import "github.com/gin-gonic/gin"

func main() {

	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := NewRunRepository(db)
	handler := NewRunHandler(repository)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/api/runs", handler.GetRuns)
	router.POST("/api/runs", handler.CreateRun)
	router.GET("/api/runs/:id", handler.GetRun)
	router.DELETE("/api/runs/:id", handler.DeleteRun)
	router.PUT("/api/runs/:id", handler.UpdateRun)

	router.Run(":8080")
}
