module git.codepku.com/examinate/exam

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.48.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20191105034135-c7e5f84aec59
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/ini.v1 v1.49.0 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace (
	git.codepku.com/examinate/exam/conf => /Users/lizhixin/go/src/projects/pkg/conf
	git.codepku.com/examinate/exam/middleware => /Users/lizhixin/go/src/projects/middleware
	git.codepku.com/examinate/exam/models => /Users/lizhixin/go/src/projects/exam/models
	git.codepku.com/examinate/exam/pkg/app => /Users/lizhixin/go/src/projects/exam/pkg/app
	git.codepku.com/examinate/exam/pkg/e => /Users/lizhixin/go/src/projects/exam/pkg/e
	git.codepku.com/examinate/exam/pkg/file => /Users/lizhixin/go/src/projects/exam/pkg/file
	git.codepku.com/examinate/exam/pkg/logging => /Users/lizhixin/go/src/projects/exam/pkg/logging
	git.codepku.com/examinate/exam/pkg/setting => /Users/lizhixin/go/src/projects/exam/pkg/setting
	git.codepku.com/examinate/exam/routers => /Users/lizhixin/go/src/projects/exam/routers
)
