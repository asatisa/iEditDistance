package main

import (
	"fmt"
	"net/http"

	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"
)

// Structure
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
	ResultCode         int32  `json:"ResultCode"`
	ResultMessage      string `json:"ResultMessage"`
	ResultData         string `json:"ResultData"`
	ResultPercentMatch string `json:"ResultPercentMatch"`
}
type autocorrectAllData struct {
	Topic              string `json:"Topic"`
	Data               string `json:"Data"`
	MinPercentMatch    string `json:"MinPercentMatch"`
	ResultCode         int32  `json:"ResultCode"`
	ResultMessage      string `json:"ResultMessage"`
	ResultData         string `json:"ResultData"`
	ResultPercentMatch string `json:"ResultPercentMatch"`
}

// Constant of Application
const app_version string = "0.3"

// Variable of Application
var i_count int = 0
var percent float32

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

// Function declaration
func getVersion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, version)
}

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

func getDistanceCorrectLatest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, outputDistanceCorrectsLatest)
}
func getDistanceCorrect(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, outputDistanceCorrects)
}
func postDistanceCorrect(c *gin.Context) {
	var newInput autocorrectAllData

	if err := c.BindJSON(&newInput); err != nil {
		return
	}

	var resultval int
	s1 := "kitten"
	s2 := "sitting"

	resultval = Calculate(newInput.Data, "sitting")

	fmt.Println("postDistanceCorrect: ", s1, ", ", s2, ", ", resultval)

	outputDistanceCorrects = append(outputDistanceCorrects, newInput)
	outputDistanceCorrectsLatest = append(outputDistanceCorrectsLatest, newInput)
	c.IndentedJSON(http.StatusCreated, newInput)
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

	//TEST
	router.GET("/hello", getHello)
	router.POST("/hello", postHello)

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
	var return_int int = 0
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

	percent = 0
	return_int = levenshtein.ComputeDistance(source1, source2)

	percent = ToFloat32(maxLength-return_int) / ToFloat32(maxLength) * 100

	fmt.Println(i_count, " : Result = ", return_int, ", ", maxLength, ", ", percent, " % #")
	//fmt.Printf("%d : Result = %d, %d, %f %% #", i_count, return_int, maxLength, percent)

	return return_int
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////
// Main function
// //////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	var return_path string

	return_path = getExecDir()
	fmt.Println("Application Version: " + app_version)
	fmt.Println("Execution Path: " + return_path)

	var api_server_ipaddress string = readINI("server", "api_server_ipaddress")
	var api_server_port string = readINI("server", "api_server_port")
	var serverrun string = api_server_ipaddress + ":" + api_server_port
	fmt.Println("serverrun: " + serverrun)
	runRestAPI()
	//runTestDistance()
}
