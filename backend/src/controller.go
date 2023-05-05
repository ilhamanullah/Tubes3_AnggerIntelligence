package src

import (
	"fmt"
	"net/http"

	"backend/src/database"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []database.Product

	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {
	// var product database.Product
	db, _ := sql.Open("mysql", "root:angger123@tcp(localhost:3306)/tubesStima")
	var id string
	id = c.Param("pertanyaan")
	// fmt.Println("================")
	// fmt.Println(id)
	var isUpdate bool
	var cek string
	cek, isUpdate = AddPertanyaan(id, db)
	fmt.Println(cek)
	if cek != "" {
		if isUpdate {
			c.JSON(http.StatusOK, gin.H{"product": cek})
		} else {
			c.JSON(http.StatusOK, gin.H{"product": cek})
		}

	}
	// answer := getDay(id)
	// if answer != "String input tidak valid" {
	// 	c.JSON(http.StatusOK, gin.H{"product": answer})
	// }
	// else{
	// 	// answer = calculator
	// 	// if(answer calculator){

	// 	// }
	// 	else {

	// 	}
	// }

	// if err := database.DB.First(&product, id).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
	// 		return
	// 	default:
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// }

}

// func Create(c *gin.Context) {

// 	var product database.Product

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	database.DB.Create(&product)
// 	c.JSON(http.StatusOK, gin.H{"product": product})
// }

// func Update(c *gin.Context) {
// 	var product database.Product
// 	id := c.Param("id")

// 	// if err := c.ShouldBindJSON(&product); err != nil {
// 	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 	// 	return
// 	// }

// 	// if database.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
// 	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

// }

// func Delete(c *gin.Context) {

// 	var product database.Product

// 	var input struct {
// 		Id json.Number
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	id, _ := input.Id.Int64()
// 	if database.DB.Delete(&product, id).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
// }
