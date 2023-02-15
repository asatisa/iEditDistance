package main

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func test_GetExcel() string {
	initializeVariable()
	excel_filename := strings.Replace(original_excelCompareFileName, "{%TOPIC%}", "ชื่อบริษัท", -1)
	f, err := excelize.OpenFile(excel_filename)
	if err != nil {
		fmt.Println(err)
		return "N_A"
	}

	cellVal := f.GetCellValue("Sheet1", "A1")
	if err != nil {
		fmt.Println(err)
		return cellVal
	}
	fmt.Println("celvalue = " + cellVal)
	return cellVal
}

func jSonTestData() {
	CompanyName := "APPLE"
	Data := "MicroSoft"
	xjsonData := autocorrectAllData{Topic: "ชื่อบริษัท", Data: Data, MinPercentMatch: "80", ResultCode: 0, ResultMessage: "Init", ResultData: "_", ResultPercentMatch: "_"}
	fmt.Println("CompanyName", CompanyName)
	fmt.Println("xjsonData", xjsonData)

}
