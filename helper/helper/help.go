package helper

import (
	"appointed-registration/global"
	"encoding/json"
	"errors"
)

/**处理字符串为map类型的数据
* 代码描述:
* 作者:小大白兔
* 创建时间:2022/10/06 19:47:21
 */
func FormatString(str string) ([]interface{}, error) {
	arr, result := []byte(str), map[string]interface{}{}

	if err := json.Unmarshal(arr, &result); err != nil {
		global.LogSuger.Errorf("数据转换失败: " + err.Error())
		return nil, errors.New("数据转换失败: " + err.Error())
	}

	res := result["data"].(map[string]interface{})["list"].([]interface{})

	return res, nil
}
