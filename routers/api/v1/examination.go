package v1

import (
    "fmt"
    // "log"
    "net/http"
    "encoding/json"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    //"github.com/astaxie/beego/validation"
    "github.com/unknwon/com"

    "git.codepku.com/examinate/exam/models"
    "git.codepku.com/examinate/exam/pkg/e"
    "git.codepku.com/examinate/exam/pkg/app"
    // "git.codepku.com/examinate/exam/pkg/util"
    // "git.codepku.com/examinate/exam/pkg/setting"
)

// 获取考试详情页面
func GetExaminationExamineeDetail(c *gin.Context) {
    r := app.Gin{C: c}
    admissionTicket := com.StrTo(c.Param("admissionTicket")).MustInt()
    maps := make(map[string]interface{})
    maps["admission_ticket"] = admissionTicket
    maps["examinee_id"] = 1277

    data, err := models.GetExaminationExamineeDetail(maps)
    // log.Printf("%+v", err)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            r.ResponseE(http.StatusForbidden, e.ERROR_NOT_EXIST_EXAMTESTING, nil)
            return
        }
        r.ResponseE(http.StatusInternalServerError, e.ERROR, nil)
        return
    }

    val, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Umarshal failed:", err)
        return
    }

    for _, v := range val {
        fmt.Printf(v)
    }

    // v := reflect.ValueOf(data)
    // for i := 0; i < v.NumField(); i++ {
    //     fmt.Printf("Field %d: %v, %s; \n", i, v.Field(i), v.FieldByName("s1"))
    // }

    r.ResponseE(http.StatusOK, e.SUCCESS, data)
}