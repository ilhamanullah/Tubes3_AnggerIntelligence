package src

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
)

// func Index(c *gin.Context) {

// 	var products []database.Product

// 	database.DB.Find(&products)
// 	c.JSON(http.StatusOK, gin.H{"products": products})

// }
var isKmp bool

func AlgoOption(c *gin.Context) {
	// fmt.Println("halo")
	type Algo struct {
		Value string `json:"value"`
	}
	var algo Algo
	if err := c.ShouldBindJSON(&algo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(algo.Value)
	if algo.Value == "Option 1" {
		isKmp = true
	} else if algo.Value == "Option 2" {
		isKmp = false
	}
	// fmt.Println("hal00o")
}

func Show(c *gin.Context) {
	// var product database.Product
	db, _ := sql.Open("mysql", "root:angger123@tcp(localhost:3306)/tubesStima")
	var id string
	id = c.Param("pertanyaan")
	// var isKmp bool
	isKmp = true
	var cek string

	cek = AddPertanyaan(id, db) // add/update pertanyaan
	// fmt.Println("1")
	if cek != "" {
		c.JSON(http.StatusOK, gin.H{"product": cek})
	} else {
		cek = hapusPertanyaan(id, db) // del pertanyaan
		// fmt.Println("2")
		if cek != "" {
			c.JSON(http.StatusOK, gin.H{"product": cek})
		} else {
			cek = getDay(id) // cek hari
			// fmt.Println("3")
			if cek != "String input tidak valid" {
				c.JSON(http.StatusOK, gin.H{"product": cek})
			} else {
				cek = getVal(id)
				println(id)
				println(cek)
				// println("qwer")
				if cek != "" {
					c.JSON(http.StatusOK, gin.H{"product": cek})
				} else {
					cek = findJawaban(id, db, isKmp)
					if cek != "" {
						c.JSON(http.StatusOK, gin.H{"product": cek})
					}
				}
			}
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
