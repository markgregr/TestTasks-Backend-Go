package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCoins(c *gin.Context) {
	coins, err := h.repo.GetAllCoins()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"coins": coins})
}

func (h *Handler) GetCoinByCurrency(c *gin.Context) {
	currency := c.Param("currency")

	coins, err := h.repo.GetCoinByCurrency(currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"coin": coins})
}