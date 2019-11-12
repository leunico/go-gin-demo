package routers

import (
    "github.com/gin-gonic/gin"

    "git.codepku.com/examinate/exam/routers/api"
    "git.codepku.com/examinate/exam/routers/api/v1"
    "git.codepku.com/examinate/exam/middleware/jwt"
)

func InitRouter() *gin.Engine {
    r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.GET("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test go!!!!",
        })
    })
    
    // login
    r.POST("/auth", api.GetAuth)

    apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
    {
        // 考试详情
        apiv1.GET("/:admissionTicket/testing", v1.GetExaminationExamineeDetail)
    }
    
    // apiv1 := r.Group("/api/v1")
    // {
    //     // 考试详情
    //     apiv1.GET("/:admissionTicket/testing", v1.GetExaminationExamineeDetail)
    // }

    return r
}