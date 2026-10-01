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

	router.Run(":8080")
}
