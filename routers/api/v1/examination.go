package v1

import (
    // "fmt"
    // "log"
    "time"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    //"github.com/astaxie/beego/validation"
    "github.com/unknwon/com"

    "git.codepku.com/examinate/exam/models"
    "git.codepku.com/examinate/exam/pkg/e"
    "git.codepku.com/examinate/exam/pkg/app"
    "git.codepku.com/examinate/exam/pkg/util"
    // "git.codepku.com/examinate/exam/pkg/setting"
)

// 获取考试详情页面
func GetExaminationExamineeDetail(c *gin.Context) {
    r := app.Gin{C: c}
    maps := make(map[string]interface{})
    maps["admission_ticket"] = com.StrTo(c.Param("admissionTicket")).MustInt()
    maps["examinee_id"] = 891

    data, err := models.GetExaminationExamineeDetail(maps)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            r.ResponseE(http.StatusForbidden, e.ERROR_NOT_EXIST_EXAMTESTING, nil)
            return
        }
        r.ResponseE(http.StatusInternalServerError, e.ERROR, nil)
        return
    }

    val, err := util.StructJsonToMap(data)
    if err != nil {
        r.ResponseE(http.StatusInternalServerError, e.ERROR_JSON_HANDLE, nil)
        return
    }

    // 以下为了适配Laravel之前写的
    if val["status"].(float64) == models.STATUS_EXAMINATION {
        val["achievements"] = map[string]float64{
            "objective_score": val["objective_score"].(float64),
            "subjective_score": val["subjective_score"].(float64),
            "rank": val["objective_score"].(float64),
        }
    }

    // 适配状态
    val["examinee_tencent_face_testing"] = false
    for _, v := range val["examinee_tencent_faces"].([]interface{}) {
        tmps := v.(map[string]interface{})
        if tmps["type"].(int) == models.TENCENT_FACE_TYPE_BEFORE && tmps["result"].(string) == "Success" {
            val["examinee_tencent_face_testing"] = true
            break
        }
    }

    // 适配时间
    now := util.JSONTime{time.Now()}.String()
    if now < val["start_at"].(string) {
        val["examination_status"] = 0
    } else if now >= val["start_at"].(string) && now <= val["end_at"].(string) {
        val["examination_status"] = 1
    } else if now > val["end_at"].(string) {
        val["examination_status"] = 2
    } else {
        val["examination_status"] = -1
    }

    r.ResponseE(http.StatusOK, e.SUCCESS, val)
}