/*
* @Time    : 2020-08-20 16:00
* @Author  : CoderCharm
* @File    : ping.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package routes

import (
	"domain-check/config"
	"domain-check/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var currentCache saveDate

type saveDate struct {
	AccessToken string
	Expire      int
}

func (c *saveDate) SetExpire(expire int) {
	c.Expire = expire
}

func (c *saveDate) SetAccessToken(token string) {
	c.AccessToken = token
}

func addCheckRoutes(rg *gin.RouterGroup) {
	check := rg.Group("/domain")

	check.GET("/check", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "requests params fail",
				"data":    nil,
			})
			return
		}

		isRefresh := currentCache.Expire > int(time.Now().Unix())

		var AccessToken string
		if isRefresh {
			fmt.Println("不需要更新")
			fmt.Println(currentCache)
			// 不需要更新
			AccessToken = currentCache.AccessToken
		} else {
			// 获取accessToken
			AccessToken = utils.GetAccessToken(config.GlobConfig.AppId, config.GlobConfig.AppSecret)

			// 保存accessToken 给定过期两小时
			currentCache.SetExpire(int(time.Now().Unix() + 2*60*60))
			currentCache.SetAccessToken(AccessToken)
			fmt.Println("需要更新")
			fmt.Println(currentCache)
		}
		// 获取短网址
		shoutUrl := utils.GetShortUrl(AccessToken, url)
		// 检测是否冻结
		isFreeze := utils.CheckUrl(shoutUrl)

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"isFreeze": isFreeze,
				"url":      url,
			},
		})
	})
}
