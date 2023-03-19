package controller

import (
	"fmt"
	"net/http"

	models "github.com/Modifa/finder_organization.git/models"
	services "github.com/Modifa/finder_organization.git/services"

	"github.com/Modifa/finder_organization.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterOrganization(c *gin.Context) {

	db := services.DB{}

	var rb models.Returnblock
	var org models.OrganizationRequest
	if err := c.ShouldBindBodyWith(&org, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	number, _ := utils.FormatMobileNumber(org.Mobile_Number)
	org.Mobile_Number = number
	// _, err := db.RegisterDeveloper("developer.fn_Register_system_user_developer", u)
	errDel := db.SAVEONDBNPRETURN("organization.fn_Register_Organization", org)

	if errDel != nil {
		fmt.Println("QueryRow failed: ", errDel.Error())
		errormessage := fmt.Sprintf("%v\n", errDel)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))

		return
	}
	var t models.GetOrganizationProfile
	t.EmailAddress = org.Email_address

	response, err := db.GetOrganizationProfile("organization.fn_get_organiation_profile_email", t)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}

	go func() {
		err1 := services.SaveOrganizationprofile(response[0])
		if err != nil {
			fmt.Println(err1)
		}
	}()
}

//
func GetOrganizationProfile(c *gin.Context) {
	var rb models.Returnblock
	var u models.OrganizationRequest1
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	d, profile := services.GetOrganizationProfileRedis(u.Email)
	active := utils.StringToBool(profile.Status)
	// developerID := profile.Id
	if !d {
		c.JSON(http.StatusOK, rb.New(false, "Acoount Does Not Exist", profile))
		return
	} else if !active {
		c.JSON(http.StatusOK, rb.New(false, "Account Not Activated", profile))
		return
	} else if profile.Password != u.Password {
		c.JSON(http.StatusOK, rb.New(false, "Password Incorrect", profile))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", profile))
		return
	}
}

//Activate Account
func ActivateOrganization(c *gin.Context) {

	db := services.DB{}

	var rb models.Returnblock
	var org models.GetOrganizationProfileID
	if err := c.ShouldBindBodyWith(&org, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	errDel := db.SAVEONDBNPRETURN("organization.fn_activate_organization", org)

	if errDel != nil {
		fmt.Println("QueryRow failed: ", errDel.Error())
		errormessage := fmt.Sprintf("%v\n", errDel)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))

		return
	}

	response, err := db.GetOrganizationProfile("organization.fn_get_organiation_profile_id", org)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}

	go func() {
		err1 := services.SaveOrganizationprofile(response[0])
		if err != nil {
			fmt.Println(err1)
		}
	}()
}
