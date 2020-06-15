package atif

// EncodeData encodes the data for an AT message.
func EncodeData(datetime string, values []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"datetime": datetime, "values": values}
}

func EncodeOutput(entity string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"entity": entity, "data": data}
}

func DecodeOutput(output map[string]interface{}) (string, map[string]interface{}) {
	return output["entity"].(string), output["data"].(map[string]interface{})
}

func DecodeData(message map[string]interface{}) (string, []map[string]interface{}) {
	return message["datetime"].(string), message["values"].([]map[string]interface{})
}
