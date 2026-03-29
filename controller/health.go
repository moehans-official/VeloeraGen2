package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ok",
		"data": gin.H{
			"version":    common.Version,
			"start_time": common.StartTime,
		},
	})
}

func Readyz(c *gin.Context) {
	type componentState struct {
		Status string `json:"status"`
		Error  string `json:"error,omitempty"`
	}

	checks := map[string]componentState{}
	ready := true

	if err := model.PingDB(); err != nil {
		ready = false
		checks["database"] = componentState{Status: "down", Error: err.Error()}
	} else {
		checks["database"] = componentState{Status: "up"}
	}

	if common.RedisEnabled && common.RDB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if _, err := common.RDB.Ping(ctx).Result(); err != nil {
			ready = false
			checks["redis"] = componentState{Status: "down", Error: err.Error()}
		} else {
			checks["redis"] = componentState{Status: "up"}
		}
	} else if common.RedisEnabled {
		ready = false
		checks["redis"] = componentState{Status: "down", Error: "redis enabled but client not initialized"}
	} else {
		checks["redis"] = componentState{Status: "disabled"}
	}

	statusCode := http.StatusOK
	message := "ready"
	if !ready {
		statusCode = http.StatusServiceUnavailable
		message = "not ready"
	}

	c.JSON(statusCode, gin.H{
		"success": ready,
		"message": message,
		"data": gin.H{
			"version": common.Version,
			"checks":  checks,
		},
	})
}

