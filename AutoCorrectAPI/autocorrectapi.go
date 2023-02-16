// Title: Auto Correct API
// Filename: AutoCorrectAPI.go
// Editor: Atthapol.w
// Update: 2023-02-14 Valentine's day.
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"

	//"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/360EntSecGroup-Skylar/excelize"
	//"github.com/xuri/excelize"
)

// ************ Start: Constant of Application ************//
const app_version string = "1.0.230216.02"
const app_comment string = "Beta version: remove debugmode"
const max_excel_rows int = 500

type debug_mode int

const (
	mode_none = iota //do nothing
	mode_default
	mode_deep_dive
)

const iDEBUG_MODE debug_mode = mode_none //iDEBUG_MODE_DEFAULT
const isINFO_MODE = true

//const prefixExcelCompareFileName = "./Data_Compare_" //.xlsx"
//************ End: Constant of Application ************//

// ************ Start: Structure of Application ************//
type mappingValue struct {
	KEY   string `json:"KEY"`
	Value string `json:"Value"`
}
type autocorrectInput struct {
	Topic           string `json:"Topic"`
	Data            string `json:"Data"`
	MinPercentMatch string `json:"MinPercentMatch"`
}
type autocorrectOutput struct {
	ResultCode         int    `json:"ResultCode"`
	ResultMessage      string `json:"ResultMessage"`
	ResultData         string `json:"ResultData"`
	ResultPercentMatch string `json:"ResultPercentMatch"`
}
type autocorrectAllData struct {
	Topic              string `json:"Topic"`
	Data               string `json:"Data"`
	MinPercentMatch    string `json:"MinPercentMatch"`
	ResultCode         int    `json:"ResultCode"`
	ResultMessage      string `json:"ResultMessage"`
	ResultData         string `json:"ResultData"`
	ResultPercentMatch string `json:"ResultPercentMatch"`
}

type returnCompare struct {
	ResultData         string
	ResultPercentMatch float32
}

//************ End: Structure of Application ************//

// ************ Start: Variable and declaration of Application ************//
var i_count int = 0
var resultPercent float32 = 0.0                 //Result oof search in percent
const original_excelCompareFileName string = "" //read from INI
var excelCompareFileName string = ""
var isInitialized = false // Is initailized
var executionDir string   // Execution directory

var version = []mappingValue{
	{KEY: "Version", Value: app_version},
}

var helloworld = []mappingValue{
	{KEY: "Greeting", Value: "Hello! World."},
}

// TEST DATA
var autoCorrects = []autocorrectAllData{
	{Topic: "ชื่อบริษัท", Data: "APPLA", MinPercentMatch: "80"},
	{Topic: "TAXID-สัญญารายย่อย", Data: "1999892001002", MinPercentMatch: "90"},
	{Topic: "เลขที่บัตรประชาชน", Data: "1999892001002", MinPercentMatch: "80"},
}

var inputDistanceCorrects = []autocorrectInput{
	{Topic: "-", Data: "-", MinPercentMatch: "-"},
}

var outputDistanceCorrects = []autocorrectAllData{
	{Topic: "-", Data: "-", MinPercentMatch: "-", ResultCode: 0, ResultMessage: "-", ResultData: "-", ResultPercentMatch: "-"},
}

var outputDistanceCorrectsLatest = []autocorrectAllData{
	{Topic: "-", Data: "-", MinPercentMatch: "-", ResultCode: 0, ResultMessage: "-", ResultData: "-", ResultPercentMatch: "-"},
}

//************ End: Variable and declaration of Application ************//

// ************ Start: Function declaration and initialization ************//
func getVersion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, version)
}

func initializeVariable() (isInit bool) {
	if isInitialized == true {
		isInit = isInitialized
		return
	}
	fmt.Println("Initialization a variables.")
	executionDir = getExecDir()

	fmt.Println("Auto Correct API")
	fmt.Println("Version: " + app_version)
	fmt.Println("Comment: " + app_comment)
	fmt.Println("Execution Path: " + executionDir)

	var api_server_ipaddress string = readINI("server", "api_server_ipaddress")
	var api_server_port string = readINI("server", "api_server_port")
	var serverrun string = api_server_ipaddress + ":" + api_server_port
	fmt.Println("serverrun: " + serverrun)

	isInitialized = true
	isInit = isInitialized
	return
}

//************ End: Function declaration and initialization ************//

// ************ Start: Main function ************//
func main() {
	initializeVariable()

	//var ireturnval int = 0
	//ireturnval = CountRowsExcelCompanyValue()
	//fmt.Println("ireturnval: ", ireturnval)
	//testGetExcel()

	//GetExcelCompanyValue("A2")
	//runTestDistance()
	runRestAPI()
}

//************ End: Main function ************//

// ************ Start: Test function ************//
func testGetExcel() string {
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

//************ End: Test function ************//

// ************ Start: Function declaration ************//
func getAutoCorrect(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, autoCorrects)
}

func postAutoCorrect(c *gin.Context) {
	var newInput autocorrectAllData

	if err := c.BindJSON(&newInput); err != nil {
		return
	}

	autoCorrects = append(autoCorrects, newInput)
	c.IndentedJSON(http.StatusCreated, newInput)

}

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, helloworld)
}

func postHello(c *gin.Context) {
	var newInput mappingValue

	if err := c.BindJSON(&newInput); err != nil {
		return
	}

	helloworld = append(helloworld, newInput)
	//c.IndentedJSON(http.StatusCreated, newInput)
	c.IndentedJSON(http.StatusCreated, helloworld)
}

func getExcelFileName(inputTopic string) string {
	original_excelCompareFileName := readINI("config", "excel_compare_filename")
	excel_filename := strings.Replace(original_excelCompareFileName, "{%TOPIC%}", inputTopic, -1)
	fmt.Println("Original Excel Compare Filename: " + original_excelCompareFileName)
	//fmt.Println("Excel Compare Filename: " + excelCompareFileName)
	return excel_filename
}

func GetExcelCompanyValue(excel_filename string, axis string) string {
	f, err := excelize.OpenFile(excel_filename)
	if err != nil {
		fmt.Println(err)
		return "N_A"
	}

	cellVal := f.GetCellValue("Sheet1", axis)
	s := "--------------------------------------------------------------------------------------\n" +
		"		Get cellvalue = " + cellVal
	printDebug(mode_deep_dive, s)
	return cellVal
}

func printDebug(mode debug_mode, a ...any) {
	if iDEBUG_MODE <= mode_none {
		return // do nothing
	}

	if iDEBUG_MODE >= mode {
		fmt.Println(a...)
	}
}

func printInfo(a ...any) {
	if isINFO_MODE {
		fmt.Println(a...)
	}
}

func CountRowsExcelCompanyValue(excel_filename string) int {
	f, err := excelize.OpenFile(excel_filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	//var i_loop int = 0
	var count_rows int = 0
	var axis = ""
	printInfo("Excel Compare Filename: " + excel_filename)
	printInfo("		Count all rows of excel")
	for i_loop := 1; i_loop <= max_excel_rows; i_loop++ {
		count_rows = i_loop - 1
		axis = fmt.Sprintf("A%d", i_loop)
		cellVal := f.GetCellValue("Sheet1", axis)
		if cellVal == "" {
			printInfo("		Count rows = ", count_rows)
			return count_rows
		}
	}

	return 0
}

func getDistanceCorrectLatest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, outputDistanceCorrectsLatest)
}

func getDistanceCorrect(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, outputDistanceCorrects)
}

// Distance Correct Version 0.9
func postDistanceCorrectV0_9(c *gin.Context) {
	var inputJson autocorrectAllData

	if err := c.BindJSON(&inputJson); err != nil {
		return
	}

	var inputData string = inputJson.Data
	var inputTopic string = inputJson.Topic
	var inputMinPercentMatch string = inputJson.MinPercentMatch
	floatMinPercentMatch, _ := strconv.ParseFloat(inputMinPercentMatch, 64)

	var result_distance int = 0
	var outputJson []autocorrectAllData
	var result_percent float32 = 0
	var result_percent_string string
	var result_message string = ""
	var local_excel_filename string = ""
	var isFindDistanceFound bool = false

	//source1 := "kitten"       //source data.
	source2 := inputJson.Data //input data from client req. //"sitting"
	i_count = i_count + 1
	i_count_loop_req = 0

	// Start ---- 1. Check database file or excel dictionary is Exists
	local_excel_filename = getExcelFileName(inputTopic)
	//fmt.Println("Check database file or excel dictionary is Exists")
	if fileExist(local_excel_filename) == false {
		result_message = "Error. Find Database/Excel not found!"
		outputJson = []autocorrectAllData{
			{Topic: inputTopic,
				Data:               inputData,
				MinPercentMatch:    inputMinPercentMatch,
				ResultCode:         -999, //Failure
				ResultMessage:      result_message,
				ResultData:         "",
				ResultPercentMatch: "0"},
		}
		fmt.Println("Error! " + result_message)
		outputDistanceCorrectsLatest = nil
		outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, inputJson)
		c.IndentedJSON(http.StatusCreated, outputJson)
		return //exit function
	}
	// Finish ---- 1. Check database file or excel dictionary is Exists

	// Start ---- 2. Compare data in  loop from source vs excel file
	var axis = ""
	printDebug(mode_default, "		First init = 0")
	var current_max_percent_result float32 = 0.0

	var max_loop int = CountRowsExcelCompanyValue(local_excel_filename)
	for i_excel_current_row := 1; i_excel_current_row <= max_loop; i_excel_current_row++ {
		axis = fmt.Sprintf("A%d", i_excel_current_row)
		source1 := GetExcelCompanyValue(local_excel_filename, axis)

		// Clear a variables
		result_message = ""
		outputJson = nil
		result_distance = 0
		result_percent = 0.0

		printDebug(mode_deep_dive, "Loop: [", i_excel_current_row, "]: DistanceCorrect: of ", source1, ", ", source2)
		fmt.Println()
		result_distance, result_percent = CalculateDistance(source1, source2)
		result_percent_string = fmt.Sprintf("%f", result_percent) //convert to string

		// Condition is pass
		if result_percent >= float32(floatMinPercentMatch) {
			outputDistanceCorrects = append(outputDistanceCorrects, inputJson)
			isFindDistanceFound = true
			result_message = fmt.Sprintf("Found: data distance = %d", result_distance)

			if result_percent > current_max_percent_result {
				current_max_percent_result = result_percent
				outputJson = []autocorrectAllData{
					{Topic: inputTopic,
						Data:               inputData,
						MinPercentMatch:    inputMinPercentMatch,
						ResultCode:         0, //Success
						ResultMessage:      result_message,
						ResultData:         source1,
						ResultPercentMatch: result_percent_string},
				}
				outputDistanceCorrectsLatest = nil
				outputDistanceCorrectsLatest = outputJson
				//c.IndentedJSON(http.StatusCreated, outputJson)
			}
		}
	}
	// Finish ---- 2. Compare data in  loop from source vs excel file

	if isFindDistanceFound { // find found
		c.IndentedJSON(http.StatusCreated, outputDistanceCorrectsLatest)
		return
	} else { // find not found
		result_message = "Find data mapping not found"
		outputJson = []autocorrectAllData{
			{Topic: inputTopic,
				Data:               inputData,
				MinPercentMatch:    inputMinPercentMatch,
				ResultCode:         -1, //Failure
				ResultMessage:      result_message,
				ResultData:         "",
				ResultPercentMatch: "0"},
		}
		outputDistanceCorrectsLatest = nil
		//outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, inputJson)
		c.IndentedJSON(http.StatusCreated, outputJson)
		return
	}
}

// Distance Correct Version 1
// Update: 2023-02-16
func postDistanceCorrect(c *gin.Context) {
	var InputJsonOriginal autocorrectInput
	var inputJsonAllData autocorrectAllData

	if err := c.BindJSON(&InputJsonOriginal); err != nil {
		return
	}

	var result_distance int = 0
	var outputJson []autocorrectAllData
	var result_percent float32 = 0
	var result_percent_string string
	var result_message string = ""
	var local_excel_filename string = ""
	var isFindDistanceFound bool = false

	var inputTopic string = InputJsonOriginal.Topic
	var inputData string = InputJsonOriginal.Data
	var inputMinPercentMatch string = InputJsonOriginal.MinPercentMatch
	//inputJsonAllData = `{"Topic": "", "Data": "", "MinPercentMatch": "", ResultCode: -999, ResultMessage: "", ResultData: "", ResultPercentMatch: ""}`
	inputJsonAllData.Topic = inputTopic
	inputJsonAllData.Data = inputData
	inputJsonAllData.MinPercentMatch = inputMinPercentMatch
	inputJsonAllData.ResultCode = -999
	inputJsonAllData.ResultMessage = ""
	inputJsonAllData.ResultPercentMatch = ""
	inputJsonAllData.ResultPercentMatch = "-999"

	//inputJsonAllData = append(inputJsonAllData, inputJsonAllDataBuffer)

	floatMinPercentMatch, _ := strconv.ParseFloat(inputMinPercentMatch, 64)

	//source1 := "kitten"       //source data.
	source2 := inputJsonAllData.Data //input data from client req. //"sitting"
	i_count = i_count + 1
	i_count_loop_req = 0

	// Start ---- 1. Check database file or excel dictionary is Exists
	local_excel_filename = getExcelFileName(inputTopic)
	//fmt.Println("Check database file or excel dictionary is Exists")
	if fileExist(local_excel_filename) == false {
		result_message = "Error. Find Database/Excel not found!"
		outputJson = []autocorrectAllData{
			{Topic: inputTopic,
				Data:               inputData,
				MinPercentMatch:    inputMinPercentMatch,
				ResultCode:         -999, //Failure
				ResultMessage:      result_message,
				ResultData:         "",
				ResultPercentMatch: "0"},
		}
		fmt.Println("Error! " + result_message)
		outputDistanceCorrectsLatest = nil
		outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, inputJsonAllData)
		c.IndentedJSON(http.StatusCreated, outputJson)
		return //exit function
	}
	// Finish ---- 1. Check database file or excel dictionary is Exists

	// Start ---- 2. Compare data in  loop from source vs excel file
	var axis = ""
	printInfo("		First init = 0")
	var current_max_percent_result float32 = 0.0

	var max_loop int = CountRowsExcelCompanyValue(local_excel_filename)
	for i_excel_current_row := 1; i_excel_current_row <= max_loop; i_excel_current_row++ {
		axis = fmt.Sprintf("A%d", i_excel_current_row)
		source1 := GetExcelCompanyValue(local_excel_filename, axis)

		// Clear a variables
		result_message = ""
		outputJson = nil
		result_distance = 0
		result_percent = 0.0

		result_distance, result_percent = CalculateDistance(source1, source2)
		result_percent_string = fmt.Sprintf("%f", result_percent) //convert to string
		//printDebug("		Loop: [", i_excel_current_row, "]: Distance Correct: of ", source1, ", ", source2, ", Result = ", result_percent_string, "%")

		// Condition is pass
		if result_percent >= float32(floatMinPercentMatch) {
			outputDistanceCorrects = append(outputDistanceCorrects, inputJsonAllData)
			isFindDistanceFound = true
			result_message = fmt.Sprintf("Found: data distance = %d", result_distance)
			printDebug(mode_deep_dive, "		FOUND! row = ", i_excel_current_row, result_message)

			if result_percent > current_max_percent_result {
				current_max_percent_result = result_percent
				outputJson = []autocorrectAllData{
					{Topic: inputTopic,
						Data:               inputData,
						MinPercentMatch:    inputMinPercentMatch,
						ResultCode:         0, //Success
						ResultMessage:      result_message,
						ResultData:         source1,
						ResultPercentMatch: result_percent_string},
				}
				outputDistanceCorrectsLatest = nil
				outputDistanceCorrectsLatest = outputJson
				//c.IndentedJSON(http.StatusCreated, outputJson)
			}
		}
	}
	// Finish ---- 2. Compare data in  loop from source vs excel file

	if isFindDistanceFound { // find found
		c.IndentedJSON(http.StatusCreated, outputDistanceCorrectsLatest)
		return
	} else { // find not found
		result_message = "Find data mapping not found"
		outputJson = []autocorrectAllData{
			{Topic: inputTopic,
				Data:               inputData,
				MinPercentMatch:    inputMinPercentMatch,
				ResultCode:         -1, //Failure
				ResultMessage:      result_message,
				ResultData:         "",
				ResultPercentMatch: "0"},
		}
		outputDistanceCorrectsLatest = nil
		//outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, inputJsonAllData)
		c.IndentedJSON(http.StatusCreated, outputJson)
		return
	}
}

func runRestAPI() {
	router := gin.Default()

	//Auto Correct
	router.GET("/autocorrect", getAutoCorrect)
	router.POST("/autocorrect", postAutoCorrect)

	//Auto Correct by Edit Distance
	router.GET("/distancecorrect", getDistanceCorrect)
	router.GET("/distancecorrectlastest", getDistanceCorrectLatest)
	router.POST("/distancecorrect", postDistanceCorrect)

	// TEST Hello
	router.GET("/hello", getHello)
	router.POST("/hello", postHello)

	// Version
	router.GET("/version", getVersion)

	var api_server_ipaddress string = readINI("server", "api_server_ipaddress")
	var api_server_port string = readINI("server", "api_server_port")
	var serverrun string = api_server_ipaddress + ":" + api_server_port
	router.Run(serverrun)
}

func runTestDistance() int {
	var ret_distance int = 0
	var ret_percent float32 = 0.0

	//source1 := "kitten"
	//source2 := "sitting"
	//distance := levenshtein.ComputeDistance(source1, source2)
	//fmt.Printf("The distance between %s and %s is %d.\n", source1, source2, distance)
	ret_distance, ret_percent = CalculateDistance("kitten", "sitting")
	fmt.Printf("The distance between %s and %s is %d and %f.\n", "kitten", "sitting", ret_distance, ret_percent)
	ret_distance, ret_percent = CalculateDistance("มกราคม", "มกรคม")
	fmt.Printf("The distance between %s and %s is %d and %f.\n", "มกราคม", "มกรคม", ret_distance, ret_percent)
	ret_distance, ret_percent = CalculateDistance("มกราคม", "มกคม")
	fmt.Printf("The distance between %s and %s is %d and %f.\n", "มกราคม", "มกคม", ret_distance, ret_percent)
	return 0
}

var i_count_loop_req int = 0

func CalculateDistance(source1 string, source2 string) (distance int, percent float32) {
	//var result_run int = 0
	var source1Length int
	var source2Length int
	var maxLength int = 0

	i_count_loop_req = i_count_loop_req + 1
	resultPercent = 0
	distance = 0

	printDebug(mode_default, "--------------------------------------------------------------------------------------")
	printDebug(mode_default, "[", i_count, "]", i_count_loop_req, " : Calculation UTF8 Distance of: source1 = ", source1, ", source2 = ", source2)

	// Length of UTF-8 encoded string
	source1Length = utf8.RuneCountInString(source1) //len(source1)
	source2Length = utf8.RuneCountInString(source2) //len(source2)
	//fmt.Printf("Length source1: %d\n", source1Length)
	//fmt.Printf("Length source2: %d\n", source2Length)

	if source1Length >= source2Length {
		maxLength = source1Length
	} else {
		maxLength = source2Length
	}

	distance = levenshtein.ComputeDistance(source1, source2)
	resultPercent = ToFloat32(maxLength-distance) / ToFloat32(maxLength) * 100

	printDebug(mode_default, "[", i_count, "]", i_count_loop_req, " :             distance = ", distance, ", maxchar = ", maxLength, ", percent = ", resultPercent, " %")

	return distance, resultPercent
}

//************ End: Function declaration ************//
