/*
* @Time    : 2020-08-21 15:07
* @Author  : CoderCharm
* @File    : tools.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package utils

import (
	"encoding/json"
	"fmt"
)

/*
字符串转Map

**/
func StringToMap(content string) map[string]interface{} {

	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		fmt.Println("string转map失败", err)
	}
	return resMap
}
