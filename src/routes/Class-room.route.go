package routes

import "github.com/gin-gonic/gin"

func ClassRoomRoute(g *gin.RouterGroup) {
	g.GET("/recommend-course", func(r *gin.Context) {
		r.AbortWithStatusJSON(200, gin.H{})
	})
	g.GET("/for-you", func(r *gin.Context) {
		r.AbortWithStatusJSON(200, gin.H{})
	})
	g.GET("/categories", func(r *gin.Context) {
		r.AbortWithStatusJSON(200, gin.H{})
	})
	g.GET("/purchase/:slug", func(r *gin.Context) {
		r.AbortWithStatusJSON(200, gin.H{})
	})
	g.GET("/purchased/slug/:slug/courses/:courseId/:videoId", func(r *gin.Context) {
		r.AbortWithStatusJSON(200, gin.H{})
	})
}
