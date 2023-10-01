package helpers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	fmt.Println("CheckUserType Called")
	userType := c.GetString("user_type")
	err = nil
	fmt.Println("User Type: ", userType, "User Role: ", role)
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	fmt.Println("MatchingUserTypeToUid Called")
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}
