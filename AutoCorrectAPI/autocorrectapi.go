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

	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"

	//"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/360EntSecGroup-Skylar/excelize"
	//"github.com/xuri/excelize"
)

// ************ Start: Constant of Application ************//
const app_version string = "0.7"

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
var resultPercent float32 = 0.0               //Result oof search in percent
var original_excelCompareFileName string = "" //read from INI
var excelCompareFileName string = ""
var isInitialized = false // Is initailized
var executionDir string   // Execution directory

var version = []mappingValue{
	{KEY: "Version", Value: app_version},
}

var helloworld = []mappingValue{
	{KEY: "Greeting", Value: "Hello! World."},
}

var autoCorrects = []autocorrectAllData{
	{Topic: "ชื่อบริษัท", Data: "APPLA", MinPercentMatch: "80"},
	{Topic: "TAXID-สัญญารายย่อย", Data: "1999892001002", MinPercentMatch: "90"},
	{Topic: "เลขที่บัตรประชาชน", Data: "1999892001002", MinPercentMatch: "80"},
}

var inputDistanceCorrects = []autocorrectInput{
	{Topic: "ชื่อบริษัท", Data: "sitting", MinPercentMatch: "80"},
}

var outputDistanceCorrects = []autocorrectAllData{
	{Topic: "ชื่อบริษัท", Data: "sitting", MinPercentMatch: "80", ResultCode: 0, ResultMessage: "Success", ResultData: "sitting", ResultPercentMatch: "80"},
}

var outputDistanceCorrectsLatest = []autocorrectAllData{
	{Topic: "ชื่อบริษัท", Data: "sitting", MinPercentMatch: "80", ResultCode: 0, ResultMessage: "Success", ResultData: "sitting", ResultPercentMatch: "80"},
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
	original_excelCompareFileName := readINI("config", "excel_compare_filename")
	fmt.Println("Original Excel Compare Filename: " + original_excelCompareFileName)
	//excelCompareFileName := strings.Replace(original_excelCompareFileName, "{%TOPIC%}", TopicName, -1)
	fmt.Println("Excel Compare Filename: " + excelCompareFileName)
	executionDir = getExecDir()

	fmt.Println("Auto Correct API")
	fmt.Println("Version: " + app_version)
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
	testGetExcel()
	//runRestAPI()

	//GetExcelCompanyValue("A2")
	//runTestDistance()
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
	excel_filename := strings.Replace(original_excelCompareFileName, "{%TOPIC%}", inputTopic, -1)
	return excel_filename
}

func GetExcelCompanyValue(excel_filename string, axis string) string {
	f, err := excelize.OpenFile(excel_filename)
	if err != nil {
		fmt.Println(err)
		return "N_A"
	}

	// cellVal := f.GetCellValue("Sheet1", "A1")
	cellVal := f.GetCellValue("Sheet1", axis)
	fmt.Println("cell value = " + cellVal)
	return cellVal
}

func CountRowsExcelCompanyValue(excel_filename string) int {
	f, err := excelize.OpenFile(excel_filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	var i_loop int = 0
	var axis = ""
	fmt.Println("First init = ", i_loop)
	for i_loop := 1; i_loop <= 200; i_loop++ {
		axis = fmt.Sprintf("A%d", i_loop)
		cellVal := f.GetCellValue("Sheet1", axis)
		if cellVal == "" {
			return i_loop - 1
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

func postDistanceCorrect(c *gin.Context) {
	var inputJson autocorrectAllData
	//var outputJson autocorrectAllData
	//outputJson = &inputJson
	//outputJson.Data = inputJson.Data

	if err := c.BindJSON(&inputJson); err != nil {
		return
	}

	var inputData string = inputJson.Data
	var inputTopic string = inputJson.Topic
	var inputMinPercentMatch string = inputJson.MinPercentMatch
	//var floatMinPercentMatch float64
	floatMinPercentMatch, _ := strconv.ParseFloat(inputMinPercentMatch, 64)

	//var _ error
	//value, _ := sjson.Set("", "Topic", inputTopic)
	//fmt.Println("value = ", value)

	var result_run int
	var outputJson []autocorrectAllData
	var result_percent float32 = 0
	var result_percent_string string
	var result_message string = ""

	//source1 := "kitten"       //source data.
	source2 := inputJson.Data //input data from client req. //"sitting"
	var source1Length int
	var source2Length int
	var maxLength int = 0
	var local_excel_filename string = ""
	i_count = i_count + 1

	// Start ---- 1. Check database file or excel dictionary is Exists
	local_excel_filename = getExcelFileName(inputTopic)
	fmt.Println("Check database file or excel dictionary is Exists")
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
	var i_loop int = 0
	var axis = ""
	fmt.Println("First init = ", i_loop)
	var max_loop int = CountRowsExcelCompanyValue(local_excel_filename)
	for i_loop := 1; i_loop <= max_loop; i_loop++ {
		axis = fmt.Sprintf("A%d", i_loop)
		source1 := GetExcelCompanyValue(local_excel_filename, axis)
		//source1 := "kitten" //source data.
		source1Length = len(source1)
		source2Length = len(source2)
		if source1Length >= source2Length {
			maxLength = source1Length
		} else {
			maxLength = source2Length
		}

		result_run = Calculate(source1, source2)

		result_percent = ToFloat32(maxLength-result_run) / ToFloat32(maxLength) * 100
		result_percent_string = fmt.Sprintf("%f", result_percent)
		fmt.Println("[", i_loop, "] DistanceCorrect: of ", source1, ", ", source2, " = ", result_run)

		outputDistanceCorrects = append(outputDistanceCorrects, inputJson)

		outputJson = nil

		// Pass
		result_message = ""
		if result_percent >= float32(floatMinPercentMatch) {
			result_message = fmt.Sprintf("Success, result = %d ", result_run)
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
			outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, inputJson)
			c.IndentedJSON(http.StatusCreated, outputJson)
			return
		} else { //Not pass or not found
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
		}
	}
	// Finish ---- 2. Compare data in  loop from source vs excel file

	// find not found
	outputDistanceCorrectsLatest = nil
	outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, inputJson)
	c.IndentedJSON(http.StatusCreated, outputJson)
	return
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

func runTestDistance() {
	var resultval int
	s1 := "kitten"
	s2 := "sitting"
	distance := levenshtein.ComputeDistance(s1, s2)
	fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
	resultval = Calculate("kitten", "sitting")
	resultval = Calculate("มกราคม", "มกรคม")
	resultval = Calculate("มกราคม", "มกคม")
	resultval = 0
	//return resultval
	fmt.Println("runTestDistance : ", s1, ", ", resultval)

	// Output:
	// The distance between kitten and sitting is 3.
}

func Calculate(source1 string, source2 string) int {
	var result_run int = 0
	var source1Length int
	var source2Length int
	var maxLength int = 0

	source1Length = len(source1)
	source2Length = len(source2)

	i_count = i_count + 1
	if source1Length >= source2Length {
		maxLength = source1Length
	} else {
		maxLength = source2Length
	}

	resultPercent = 0
	result_run = levenshtein.ComputeDistance(source1, source2)

	resultPercent = ToFloat32(maxLength-result_run) / ToFloat32(maxLength) * 100

	fmt.Println(i_count, " : When data : source1 = ", source1, ", source2 = ", source2)
	fmt.Println(i_count, " :              Result = ", result_run, ", ", maxLength, ", ", resultPercent, " %")
	//fmt.Printf("%d : Result = %d, %d, %f %% #", i_count, return_int, maxLength, percent)

	return result_run
}

//************ End: Function declaration ************//
