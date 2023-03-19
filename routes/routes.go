package routes

import (
	cont "github.com/Modifa/finder_organization.git/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//RegisterOrganization
	V1 := r.Group("/api/devfinder/")
	{
		//controller.go
		V1.POST("GetDeveloperProfile", cont.RegisterOrganization)
		V1.POST("GetOrganizationProfile", cont.GetOrganizationProfile)

	}
}
