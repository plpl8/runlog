package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
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

func (h *RunHandler) CreateRun(c *gin.Context) {
	var run Run

	if err := c.ShouldBindJSON(&run); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.repository.Create(
		c.Request.Context(),
		&run,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, run)
}

func (h *RunHandler) GetRun(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	run, err := h.repository.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "run not found",
		})
		return
	}

	c.JSON(200, run)
}

func (h *RunHandler) DeleteRun(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = h.repository.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}

func (h *RunHandler) UpdateRun(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	var run Run

	if err := c.ShouldBindJSON(&run); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.repository.Update(
		c.Request.Context(),
		id,
		&run,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	run.ID = id

	c.JSON(200, run)
}
