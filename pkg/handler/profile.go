package handler

import (
	"net/http"
	todo "to-do-list"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserProfile(c *gin.Context) {
	id, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	profile, err := h.services.GetUserProfile(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *Handler) updateUserProfile(c *gin.Context) {
	id, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var updatedUser todo.UserProfile

	if err := c.BindJSON(&updatedUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateUserProfile(updatedUser, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})
}
