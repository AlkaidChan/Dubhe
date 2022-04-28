package router

import (
	"github.com/alkaidchen/Dubhe/assets"
	"github.com/alkaidchen/Dubhe/internal/api"
	"github.com/alkaidchen/Dubhe/internal/middleware/cors"
	"github.com/alkaidchen/Dubhe/internal/middleware/jwt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"strings"
	"time"
)

func NewRouter(logger *zap.Logger) (*gin.Engine, error) {
	g := gin.New()
	g.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logger, true))
	g.Use(cors.Cors())
	g.StaticFS("/admin", http.FS(assets.Static))
	a := g.Group("/api")
	a.Use(jwt.Jwt())
	a.POST("/login", api.Login)
	a.GET("/check", api.Check)
	//admin.RouteAdminAPI(g)
	// route for admin ui
	g.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/admin") {
			//设置响应状态
			c.Writer.WriteHeader(http.StatusOK)
			indexHTML, err := fs.ReadFile(assets.Static, "index.html")
			if err != nil {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
			_, err = c.Writer.Write(indexHTML)
			if err != nil {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Flush()
			return
		}
	})
	return g, nil
}
