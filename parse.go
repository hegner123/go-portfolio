package main
import "fmt"
func parse(data interface{}, keys []string, results *[]map[string]interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			newKeys := append(keys, key)
			parse(value, newKeys, results)
		}
	case []interface{}:
		for i, value := range v {
			newKeys := append(keys, fmt.Sprintf("[%d]", i))
			parse(value, newKeys, results)
		}
	default:
		result := make(map[string]interface{})
		for i, key := range keys {
			if i == len(keys)-1 {
				result[key] = v
			} else {
				if _, ok := result[key]; !ok {
					result[key] = make(map[string]interface{})
				}
				result = result[key].(map[string]interface{})
			}
		}
		*results = append(*results, result)
	}
}
