package routes

import (
	"github.com/gin-gonic/gin"
	"www.rsvpme.com/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/rsvp", controllers.CreateRSVP)

}

//GruhaPravesamRS
//RevanthSindhuHome@gmail.com
//pwd: Texas76227
//app password : wcab erwi rvyb rowl
