package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//var json_str = `{"bool":true,"str":"hello world!","num":123}`
	//var data map[string]interface{}
	//
	//json.Unmarshal([]byte(json_str), &data)
	//fmt.Println(data)
	//fmt.Println(data["str"])
	//cds
	var json_str = `{"RequestId":"D8B4F807-76E0-5B53-9A5C-FDBD7A88F7E3","IsTruncated":false,"UserBasicInfos":{"UserBasicInfo":[{"UserId":"241092357633425764","DisplayName":"ops","UserPrincipalName":"ops01@1179890457900088.
	onaliyun.com"},{"UserId":"203028657631752387","DisplayName":"yan.iu","UserPrincipalName":"yan_test@1179890457900088.onaliyun.com"}]}}`
	var data map[string]interface{}

	json.Unmarshal([]byte(json_str), &data)
	fmt.Println(data)
	fmt.Println(data["DisplayName"])

}
