package controller

import (
	"finance/internal/core/tinkoff"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TinkoffController struct {
	service *tinkoff.Service
}

func NewTinkoffController(service *tinkoff.Service) *TinkoffController {
	return &TinkoffController{service: service}
}

func (t *TinkoffController) AddRoutes(routerGroup *gin.RouterGroup) { //TODO think about changing value without ref; https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
	routerGroup = routerGroup.Group("/tinkoff")
	{
		routerGroup.GET("/get-all", t.getAll)
	}
}

func (t *TinkoffController) getAll(c *gin.Context) {
	getAllResponse, err := t.service.GetAll(c.Request.Context(), tinkoff.GetAllQuery{})
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, &getAllResponse)
}
