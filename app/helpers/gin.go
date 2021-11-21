package helpers

import (
	"github.com/boshangad/v1/ent"
	"github.com/gin-gonic/gin"
)

const (
	ginAppName          string = "$.app"
	ginAppUserName      string = "$.appUser"
	ginAppUserTokenName string = "$.appUserToken"
	ginUserName         string = "$.user"
)

// GetGinApp 获取应用
func GetGinApp(c *gin.Context) *ent.App {
	if value, ok := c.Get(ginAppName); ok {
		if data, ok := value.(*ent.App); ok {
			return data
		}
	}
	return nil
}

// GetGinAppUserToken 获取应用用户token
func GetGinAppUserToken(c *gin.Context) *ent.AppUserToken {
	if value, ok := c.Get(ginAppUserTokenName); ok {
		if data, ok := value.(*ent.AppUserToken); ok {
			return data
		}
	}
	return nil
}

// GetGinAppUser 获取应用用户
func GetGinAppUser(c *gin.Context) *ent.AppUser {
	if value, ok := c.Get(ginAppUserName); ok {
		if data, ok := value.(*ent.AppUser); ok {
			return data
		}
	}
	return nil
}

// GetGinUser 获取用户
func GetGinUser(c *gin.Context) *ent.User {
	if value, ok := c.Get(ginUserName); ok {
		if data, ok := value.(*ent.User); ok {
			return data
		}
	}
	return nil
}

// SetGinApp 设置应用
func SetGinApp(c *gin.Context, data *ent.App) {
	c.Set(ginAppName, data)
}

// SetGinAppUserToken 设置应用用户token
func SetGinAppUserToken(c *gin.Context, data *ent.AppUserToken) {
	c.Set(ginAppUserTokenName, data)
}

// SetGinAppUser 设置应用用户
func SetGinAppUser(c *gin.Context, data *ent.AppUser) {
	c.Set(ginAppUserName, data)
}

// SetGinUser 设置用户
func SetGinUser(c *gin.Context, data *ent.User) {
	c.Set(ginUserName, data)
}
