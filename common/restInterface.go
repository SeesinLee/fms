package common

import (
	"github.com/gin-gonic/gin"
)

//restful公共接口

type RestInterface interface {
	Create()gin.HandlerFunc
	Update()gin.HandlerFunc
	Show()gin.HandlerFunc
	Delete()gin.HandlerFunc
}

type JSKeyValueCommon struct {
	Name string	`json:"name"`
	Value int	`json:"value"`
}

type AverageCompute struct {
	Value int
	Denominator int
}

type MonthSum struct {
	Value int
	Month int
}