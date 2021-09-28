package utils

import (
	"github.com/boshangad/go-api/ent"
	"github.com/gin-gonic/gin"
)

// GetGinApp 获取应用
func GetGinApp(c *gin.Context) *ent.App {
	if value, ok := c.Get("$.app"); ok {
		if data, ok := value.(*ent.App); ok {
			return data
		}
	}
	return nil
}

// GetGinAppUserToken 获取应用用户token
func GetGinAppUserToken(c *gin.Context) *ent.AppUserToken {
	if value, ok := c.Get("$.appUserToken"); ok {
		if data, ok := value.(*ent.AppUserToken); ok {
			return data
		}
	}
	return nil
}

// GetGinAppUser 获取应用用户
func GetGinAppUser(c *gin.Context) *ent.AppUser {
	if value, ok := c.Get("$.appUser"); ok {
		if data, ok := value.(*ent.AppUser); ok {
			return data
		}
	}
	return nil
}

// GetGinUser 获取用户
func GetGinUser(c *gin.Context) *ent.User {
	if value, ok := c.Get("$.user"); ok {
		if data, ok := value.(*ent.User); ok {
			return data
		}
	}
	return nil
}

// SetGinApp 设置应用
func SetGinApp(c *gin.Context, data *ent.App) {
	c.Set("$.app", data)
	return
}

// SetGinAppUserToken 设置应用用户token
func SetGinAppUserToken(c *gin.Context, data *ent.AppUserToken) {
	c.Set("$.appUserToken", data)
	return
}

// SetGinAppUser 设置应用用户
func SetGinAppUser(c *gin.Context, data *ent.AppUser) {
	c.Set("$.appUser", data)
	return
}

// SetGinUser 设置用户
func SetGinUser(c *gin.Context, data *ent.User) {
	c.Set("$.user", data)
	return
}