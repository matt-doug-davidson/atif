package atif

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	values := []map[string]interface{}{}
	value := map[string]interface{}{}
	value["field"] = "NO2"
	value["amount"] = 32.21
	values = append(values, value)
	value["field"] = "SO2"
	value["amount"] = 43.32
	values = append(values, value)

	message := EncodeData("2020-06-14T16:15:00.000Z", values)
	fmt.Println(message)

	output := EncodeOutput("/A/B/C", message)
	fmt.Println(output)

	entity, rMessage := DecodeOutput(output)
	fmt.Println("entity   ", entity)
	fmt.Println("rMessage ", rMessage)

	datetime, values := DecodeData(rMessage)
	fmt.Println(datetime)
	fmt.Println(values)

	for i, v := range values {
		fmt.Println(i)
		fmt.Println(v)
		for ii, vv := range v {
			fmt.Println(ii)
			fmt.Println(vv)
		}
	}

}
