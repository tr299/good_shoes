package router

import (
    "time"

    "good_shoes/logger"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
)

type login struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
    UserName  string
    FirstName string
    LastName  string
}

func initJwt() (*jwt.GinJWTMiddleware, error) {
    var identityKey = "good-shoes-id-key"

    return jwt.New(&jwt.GinJWTMiddleware{
        Realm:       "good shoes",
        Key:         []byte("good-shoes-6789"),
        Timeout:     24 * time.Hour,
        MaxRefresh:  3 * 24 * time.Hour,
        IdentityKey: identityKey,
        PayloadFunc: func(data interface{}) jwt.MapClaims {
            if v, ok := data.(*User); ok {
                return jwt.MapClaims{
                    identityKey: v.UserName,
                }
            }
            return jwt.MapClaims{}
        },
        IdentityHandler: func(c *gin.Context) interface{} {
            claims := jwt.ExtractClaims(c)
            return &User{
                UserName: claims[identityKey].(string),
            }
        },
        Authenticator: func(c *gin.Context) (interface{}, error) {
            var loginVals login
            if err := c.ShouldBind(&loginVals); err != nil {
                return "", jwt.ErrMissingLoginValues
            }
            userID := loginVals.Username
            password := loginVals.Password

            logger.Info(userID)
            logger.Info(password)

            if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
                return &User{
                    UserName:  userID,
                    LastName:  "last-name",
                    FirstName: "first-name",
                }, nil
            }

            return nil, jwt.ErrFailedAuthentication
        },
        Authorizator: func(data interface{}, c *gin.Context) bool {
            if v, ok := data.(*User); ok && v.UserName == "admin" {
                return true
            }

            return false
        },
        Unauthorized: func(c *gin.Context, code int, message string) {
            c.JSON(code, gin.H{
                "code":    code,
                "message": message,
            })
        },
        // TokenLookup is a string in the form of "<source>:<name>" that is used
        // to extract token from the request.
        // Optional. Default value "header:Authorization".
        // Possible values:
        // - "header:<name>"
        // - "query:<name>"
        // - "cookie:<name>"
        // - "param:<name>"
        TokenLookup: "header: Authorization, query: token, cookie: jwt",
        // TokenLookup: "query:token",
        // TokenLookup: "cookie:token",

        // TokenHeadName is a string in the header. Default value is "Bearer"
        TokenHeadName: "Bearer",

        // TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
        TimeFunc: time.Now,
    })
}
