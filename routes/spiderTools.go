/*
* @Time    : 2020-09-16 17:25
* @Author  : CoderCharm
* @File    : ua.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package routes

import (
	"domain-check/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func addSpiderToolsRoutes(rg *gin.RouterGroup) {
	spiderTools := rg.Group("")

	spiderTools.GET("/random/ua", func(c *gin.Context) {
		// 微信UA
		uaType := c.DefaultQuery("type", "android")
		var currentUa string
		if uaType == "android" {
			// 随机选择一条 android 微信移动端UA
			currentUa = utils.MicroMessengerUaAndroidList[rand.Intn(len(utils.MicroMessengerUaAndroidList))]
		} else {
			// 随机获取一条 IOS 微信移动端UA
			currentUa = utils.MicroMessengerUaIosList[rand.Intn(len(utils.MicroMessengerUaIosList))]
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"uaType": uaType,
				"ua":     currentUa,
			},
		})
	})

	spiderTools.GET("/ip", func(c *gin.Context) {
		// 返回访问客户端信息
		ip := c.ClientIP()
		userUa := c.GetHeader("User-Agent")

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"ip": ip,
				"ua": userUa,
			},
		})
	})
}
