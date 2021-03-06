package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json: "-"`        //此字段不会输出的屏幕上
	Subjects []string `json: "subjects"` // 二次编码
	IsOk     bool     `json: ",string"`
	Price    float64  `json: "price"`
}

func main() {
	// -------通过struct来生成json ------------
	s := IT{"iscast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	// 编码, 根据内容生成json文本
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))

	//格式化编码
	buf, err = json.MarshalIndent(s, "", " ")
	fmt.Println("buf = ", string(buf))

	// -------通过map来生成json ------------
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 66.666

	// 编码成json
	res, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(res))

	// ------用json存到结构体 --------
	jsonBuf := `
	{
		"company" : "itcast",
		"subjects":[
			"Go",
			"C++",
			"Python",
			"Test"
		],
		"isok" : true,
		"price": 666.666
	}`

	m2 := make(map[string]interface{}, 4)
	err = json.Unmarshal([]byte(jsonBuf), &m2)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m2)

	// 下面两行是不work的, 因为m["company"]是interface{}类型,
	// 不能强转为string类型
	// var str string
	// str = m["company"]

	//用类型断言
	for k, v := range m {
		switch value := v.(type) {
		case string:
			fmt.Printf("map[%s] ==> type is string, v = %s\n", k, value)

		}
	}
}
