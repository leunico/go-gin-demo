package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"git.codepku.com/examinate/exam/pkg/app"
	"git.codepku.com/examinate/exam/pkg/e"
	"git.codepku.com/examinate/exam/pkg/util"
	"git.codepku.com/examinate/exam/models"
)

type auth struct {
	Certificates string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
	AdmissionTicket string `valid:"Required; MaxSize(50)`
	Category uint8 `valid:"Required`
}

// @Lzx Get Auth
// @Produce  json
// @Param certificates query string true "Certificates"
// @Param password query string true "Password"
// @Param admission_ticket query string true "AdmissionTicket"
// @Param category query string true "Category"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	r := app.Gin{C: c}
	valid := validation.Validation{}

	certificates := c.Query("certificates")
	password := c.Query("password")
	category := com.StrTo(c.Query("type")).MustUint8()
	admissionTicket := c.Query("admission_ticket")

	a := auth{Certificates: certificates, Password: password, AdmissionTicket: admissionTicket, Category: category}
	ok, _ := valid.Valid(&a)
	if !ok {
		app.MarkErrors(valid.Errors)
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	isExist, err := models.CheckAuth(certificates, admissionTicket, password, category)
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_EXAMTESTING_EXAMINEE, nil)
		return
	}

	if !isExist {
		r.Response(http.StatusUnauthorized, e.ERROR_PASSWORD_LOGIN, nil)
		return
	}

	token, err := util.GenerateToken(certificates, password)
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	r.ResponseE(http.StatusOK, e.SUCCESS, map[string]string{
		"access_token": token,
		"token_type": "bearer",
	})
}