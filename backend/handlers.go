package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type RunHandler struct {
	repository *RunRepository
}

func NewRunHandler(repository *RunRepository) *RunHandler {
	return &RunHandler{
		repository: repository,
	}
}

func (h *RunHandler) GetRuns(c *gin.Context) {
	runs, err := h.repository.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, runs)
}
