package structtomap

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/homingos/campaign-svc/utils"
)

type Address struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
}

type Campaign struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Age         *int     `json:"age,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

func main() {
	// age := 6
	campaign := Campaign{
		Name:        "Aman",
		Description: "This is an example campaign",
		// Age:         &age,
	}

	updateMap := utils.StructToMap(campaign)
	updateMap["last_name"] = "singh"
	// updateMap["address.city"]="Banglore"
	// delete(updateMap, "address")
	fmt.Println(updateMap)
}

func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := strings.Split(v.Field(i).Tag.Get("json"), ",")[0]
		field := reflectValue.Field(i).Interface()
		fieldType := v.Field(i).Type.Kind()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructToMap(field)
			} else if !IsZeroVal(field) {
				res[tag] = field
			} else if fieldType == reflect.Bool {
				res[tag] = field
			}
		}
	}
	return res
}

func IsZeroVal(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
