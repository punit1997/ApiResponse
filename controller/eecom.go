package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/punit1997/ApiResponse/dbconn"
	"github.com/punit1997/ApiResponse/models/eecom"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func JobOrder(r *gin.Context){

	body := r.Request.Body
	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8000/capital/eds/v1/job_order", body)
	request.SetBasicAuth("api", "api")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyByte)
	results := gjson.GetMany(bodyString, "data._raw", "name.last", "age")  //to be change

	var JobOrderResponse eecom.JobOrderResponse
	err = json.Unmarshal(bodyByte, &JobOrderResponse)


	dbconn.DB.Create(&JobOrderResponse)

	r.JSON(200, JobOrderResponse)

}