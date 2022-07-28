package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func main() {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential("LTAI5tNf7LTGTQqZscSh16Qe", "QSFUY4NQIC1etEPpOdGg0FncJdFhND")
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := sdk.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "ims.aliyuncs.com"
	request.Version = "2019-08-15"
	request.ApiName = "ListUserBasicInfos"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())

	rejson = response.json

	//fmt.Print(response)

	//dew
	//var data map[string]interface{}
	//json.Unmarshal([]byte(response), &data)
	//fmt.Println(data)

}
