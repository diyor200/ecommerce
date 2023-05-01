package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// @Summary DownloadExcel returns an excel file with all the orders data.
// @Description DownloadExcel creates an excel file with all the orders data and sends it as a response to the client.
// @Router /api/download [get]
func (h handler) DownloadExcel(ctx *gin.Context) {
	file, err := prepareExcelFile(h.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	// save to temp file
	tempFile, err := ioutil.TempFile("", "example.xlsx")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer os.Remove(tempFile.Name())

	err = file.Write(tempFile)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// set headers
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=orders.xlsx")
	ctx.Writer.Header().Set("Content-Type", ctx.Request.Header.Get("Content-Type"))
	ctx.File(tempFile.Name())
}

func prepareExcelFile(DB *sql.DB) (*excelize.File, error) {

	f := excelize.NewFile()
	const sheetname = "Sheet1"

	categories := map[string]string{"A1": "id", "B1": "firstname", "C1": "lastname", "D1": "product name", "E1": "Price", "F1": "time"}
	for k, v := range categories {
		f.SetCellValue(sheetname, k, v)
	}

	orders, err := db.GetOrders(DB)
	if err != nil {
		return nil, err
	}

	row := 2
	for _, value := range orders {
		f.SetCellValue(sheetname, fmt.Sprintf("A%d", row), value.ID)
		f.SetCellValue(sheetname, fmt.Sprintf("B%d", row), value.User.Firstname)
		f.SetCellValue(sheetname, fmt.Sprintf("C%d", row), value.User.Lastname)
		f.SetCellValue(sheetname, fmt.Sprintf("D%d", row), value.Productname)
		f.SetCellValue(sheetname, fmt.Sprintf("E%d", row), value.Payment)
		f.SetCellValue(sheetname, fmt.Sprintf("F%d", row), value.Time)
		row++
	}

	// err := f.SaveAs("excel1.xlsx")
	// if err != nil {
	// 	panic(err)
	// }
	return f, nil
}
