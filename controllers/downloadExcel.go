package controllers

import (
	"github.com/diyor200/ecommerce/model"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func (h handler) DownloadExcel(file *excelize.File, ctx gin.Context) {
	ctx.JSON(200, gin.H{"message": "excel file"})
}

func prepareExcelFile(order []model.Orders) *excelize.File {

	file := excelize.NewFile()
	file.SetCellValue("Sheet1", "A1", "Username")
	file.SetCellValue("Sheet1", "A2", order[0].ID)
	file.SetCellValue("Sheet1", "B1", "Location")
	file.SetCellValue("Sheet1", "B2", order[0].Productname)
	file.SetCellValue("Sheet1", "C1", "Occupation")
	file.SetCellValue("Sheet1", "C2", order[0].Payment)
	return file
}
