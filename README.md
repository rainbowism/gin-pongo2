# gin-pongo2
pongo2 middleware for Gin framework.

##Example:
```Go
package main

import (
    "os"

    "github.com/gin-gonic/gin"
    "github.com/rainbowism/gin-pongo2"
)

func main() {
    switch os.Getenv("MODE") {
    case "RELEASE":
        gin.SetMode(gin.ReleaseMode)

    case "DEBUG":
        gin.SetMode(gin.DebugMode)

    case "TEST":
        gin.SetMode(gin.TestMode)

    default:
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.New()
    router.Use(gin.Recovery())

    if gin.IsDebugging() {
        router.HTMLRender = render.NewDebug("resources")
    } else {
        router.HTMLRender = render.NewProduction("resources")
    }

    router.Static("/static", "resources/static")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tpl", render.Context{"title": "Gin-pongo2!"})
    })

    router.Run(":3000")
}
```
