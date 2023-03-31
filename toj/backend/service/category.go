package service

import (
	"toj/define"
	"toj/helper"
	"toj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetCategoryList
// @Tags 公共方法
// @Summary 分类列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category-list [get]
func GetCategoryList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")

	categoryList := make([]*models.CategoryBasic, 0)
	err = models.DB.Model(new(models.CategoryBasic)).Where("name like ?", "%"+keyword+"%").
		Count(&count).Limit(size).Offset(page).Find(&categoryList).Error
	if err != nil {
		log.Println("GetCategoryList Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Failed to obtain the classification list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  categoryList,
			"count": count,
		},
	})
}

// CategoryCreate
// @Tags 管理员私有方法
// @Summary 分类创建
// @Param authorization header string true "authorization"
// @Param name formData string true "name"
// @Param parentId formData int false "parentId"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/category-create [post]
func CategoryCreate(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	category := &models.CategoryBasic{
		Identity:  helper.GetUUID(),
		Name:      name,
		ParentId:  parentId,
		CreatedAt: models.MyTime(time.Now()),
		UpdatedAt: models.MyTime(time.Now()),
	}
	err := models.DB.Create(category).Error
	if err != nil {
		log.Println("CategoryCreate Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Failed to create category",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Created successfully",
	})
}

// CategoryModify
// @Tags 管理员私有方法
// @Summary 分类修改
// @Param authorization header string true "authorization"
// @Param identity formData string true "identity"
// @Param name formData string true "name"
// @Param parentId formData int false "parentId"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/category-modify [put]
func CategoryModify(c *gin.Context) {
	identity := c.PostForm("identity")
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	if name == "" || identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "parameter is incorrect",
		})
		return
	}
	category := &models.CategoryBasic{
		Identity:  identity,
		Name:      name,
		ParentId:  parentId,
		UpdatedAt: models.MyTime(time.Now()),
	}
	err := models.DB.Model(new(models.CategoryBasic)).Where("identity = ?", identity).Updates(category).Error
	if err != nil {
		log.Println("CategoryModify Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Failed to modify category",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Modify Successfully",
	})
}

// CategoryDelete
// @Tags 管理员私有方法
// @Summary 分类删除
// @Param authorization header string true "authorization"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/category-delete [delete]
func CategoryDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "parameter is incorrect",
		})
		return
	}
	var cnt int64
	err := models.DB.Model(new(models.ProblemCategory)).Where("category_id = (SELECT id FROM category_basic WHERE identity = ? LIMIT 1)", identity).Count(&cnt).Error
	if err != nil {
		log.Println("Get ProblemCategory Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Failed to obtain the issue associated with the category",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "There is already a problem under this category and cannot be deleted",
		})
		return
	}
	err = models.DB.Where("identity = ?", identity).Delete(new(models.CategoryBasic)).Error
	if err != nil {
		log.Println("Delete CategoryBasic Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Delete failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Delete successfully",
	})
}
