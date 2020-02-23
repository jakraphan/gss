package controllers

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type Articles struct {
	DB *gorm.DB
}

type CreateArticlePayload struct {
	Title   string `binding:"required"`
	Excerpt string `binding:"required"`
	Body    string `binding:"required"`
}

type UpdateArticlePayload struct {
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Body    string `json:"body"`
}

func (a *Articles) FindAll(c *gin.Context) {
	var articles []models.Article

	if result := a.DB.Find(&articles); result.RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

func (a *Articles) FindOne(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if result := a.DB.First(&article, id); result.RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"articles": article})
}

func (a *Articles) Create(c *gin.Context) {
	var form CreateArticlePayload

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var article models.Article
	copier.Copy(&article, &form)
	a.DB.Create(&article)

	c.JSON(http.StatusCreated, gin.H{"article": article})
}

func (a *Articles) Update(c *gin.Context) {
	var form UpdateArticlePayload

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var article models.Article

	if result := a.DB.First(&article, id); result.RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	a.DB.First(&article, id)
	copier.Copy(&article, &form)
	a.DB.Save(&article)

	c.JSON(http.StatusOK, gin.H{"article": article})
}

func (a *Articles) Delete(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if result := a.DB.First(&article, id); result.RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	if err := a.DB.Unscoped().Delete(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
