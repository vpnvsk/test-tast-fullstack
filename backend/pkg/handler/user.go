package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/internal/models"
	"net/http"
	"strconv"
)

// @Summary Add new user
// @Description Add new user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user info"
// @Success 200 {integer} integer 1
// @Failure 400,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user [post]
func (h *Handler) createUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.repository.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all users
// @Description Get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.User
// @Failure 400,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	list, err := h.repository.GetAllUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Get User By Id
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/:id [get]
func (h *Handler) getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	item, err := h.repository.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary Delete user
// @Description Delete user
// @ID delete-user
// @Accept  json
// @Produce  json
// @Success 200 {string} string ok
// @Failure 400,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/:id [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.repository.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// @Summary Update user
// @Description Update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body models.UserUpdate true "user info"
// @Success 200 {string} string ok
// @Failure 400,403,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/:id [put]
func (h *Handler) updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.UserUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.repository.UpdateUser(id, input); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}
