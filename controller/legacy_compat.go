package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LegacyPlanPurchaseNotSupported(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "Legacy endpoint deprecated: please use /api/subscription/* (旧接口已废弃，请改用 /api/subscription/*)",
	})
}

func LegacyPlanAdminNotSupported(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "Legacy endpoint deprecated: please use /api/subscription/admin/* (旧接口已废弃，请改用 /api/subscription/admin/*)",
	})
}

