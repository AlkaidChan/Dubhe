package router

import (
	"github.com/alkaidchen/Dubhe/assets"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewRouter(logger *zap.Logger) (*gin.Engine, error) {
	g := gin.New()
	g.Use(cors.Default())
	g.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logger, true))

	//admin.RouteAdminAPI(g)
	// route for admin ui
	g.StaticFS("/admin", http.FS(assets.Static))
	return g, nil
}
