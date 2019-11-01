package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.codepku.com/examinate/exam/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseExam struct {
	Message string `json:"message"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// ResponseExam setting gin.JSON
func (g *Gin) ResponseE(httpCode, errCode int, data interface{}) {
	if (httpCode != http.StatusOK) {
		g.C.JSON(httpCode, ResponseExam{
			Message:  e.GetMsg(errCode),
		})
	} else {
		g.C.JSON(httpCode, data)
	}
	return
}