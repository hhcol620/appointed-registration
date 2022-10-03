package listaddress

import (
	"appointed-registration/dao"
	"appointed-registration/models/address"
	"fmt"
)

// 各个行政去代号
var AllAddressKey = map[string]string{
	"0":      "全部",
	"110101": "东城区",
	"110102": "西城区",
	"110105": "朝阳区",
	"110106": "丰台区",
	"110107": "石景山区",
	"110108": "海淀区",
	"110109": "门头沟区",
	"110111": "房山区",
	"110112": "通州区",
	"110113": "顺义区",
	"110114": "昌平区",
	"110115": "大兴区",
	"110116": "怀柔区",
	"110117": "平谷区",
	"110118": "密云区",
	"110119": "延庆区",
}

var AllAddressValue = map[string]string{
	"全部":   "0",
	"东城区":  "110101",
	"西城区":  "110102",
	"朝阳区":  "110105",
	"丰台区":  "110106",
	"石景山区": "110107",
	"海淀区":  "110108",
	"门头沟区": "110109",
	"房山区":  "110111",
	"通州区":  "110112",
	"顺义区":  "110113",
	"昌平区":  "110114",
	"大兴区":  "110115",
	"怀柔区":  "110116",
	"平谷区":  "110117",
	"密云区":  "110118",
	"延庆区":  "110119",
}
var AllGradeKey = map[int]string{
	0: "全部",
	1: "一级医院",
	2: "二级医院",
	3: "三级医院",
}
var AllGradeValue = map[string]int{
	"全部":   0,
	"一级医院": 1,
	"二级医院": 2,
	"三级医院": 3,
}

/**
* 代码描述: 行政区数据的插入和医院级别的数据插入
* 作者:小大白兔
* 创建时间:2022/10/01 21:59:32
 */
func InsertAddressAndGrade() {

	addresses, addressGrade := []address.AllAddress{}, []address.AllGrade{}

	for k, v := range AllAddressKey {
		a := address.AllAddress{
			AddressCode:  k,
			AddressValue: v,
		}
		addresses = append(addresses, a)
	}

	for k, v := range AllGradeKey {
		a := address.AllGrade{
			GradeCode:  uint(k),
			GradeValue: v,
		}
		addressGrade = append(addressGrade, a)
	}
	fmt.Println(addressGrade, addresses)
	dao.DB.Create(&addresses)

	dao.DB.Create(&addressGrade)
}
