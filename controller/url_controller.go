package controller

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dora-exku/shorturl/model"
	"github.com/dora-exku/shorturl/pkg/config"
	"github.com/dora-exku/shorturl/pkg/database"
	"github.com/dora-exku/shorturl/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlController struct {
}
type urlParams struct {
	Url string `form:"url" json:"url" valid:"required~参数错误,url~链接错误"`
}

func Jump(c *gin.Context) {

	var urlModel model.Url
	key := c.Param("key")

	result := database.DB.Table("urls").Where("path = ?", key).First(&urlModel)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求错误",
		})
		return
	}

	c.Redirect(302, urlModel.OriginalUrl)
}

func Store(c *gin.Context) {

	var params urlParams
	c.Bind(&params)

	result, err := govalidator.ValidateStruct(params)

	if false == result {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": err.Error(),
		})
		return
	}

	var url model.Url
	url.Path = util.RandChar(5)
	url.OriginalUrl = params.Url

	res := database.DB.Table("urls").Create(&url)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求频繁",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"url": fmt.Sprintf("%s/%s", config.GetString("app.url"), url.Path),
		},
	})
}
