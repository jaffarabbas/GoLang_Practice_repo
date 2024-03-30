package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized")
		return err
	}
	return err
}

func MatachUserTypeToUid(c *gin.Context, userID string) (err error) {
	userType := c.GetString("userType")
	uid := c.GetString("uid")
	err = nil
	if userType == "USER" && uid != userID {
		err = errors.New("Unauthorized")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
