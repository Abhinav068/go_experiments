package structtomap

import (
	"fmt"
	"reflect"
	"strings"
)

type Address struct {
	AddressMeta `json:",inline"`
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
}
type AddressMeta struct {
	HouseNo string `json:"house_no"`
	StNo    string `json:"st_no"`
}

type Campaign struct {
	Name        string  `json:"name,omitempty"`
	LastName    string  `json:"last_name,omitempty"`
	Description string  `json:"description,omitempty"`
	Age         *int    `json:"age,omitempty"`
	Address     Address `json:"address,omitempty"`
}

func main() {
	// age := 6
	campaign := Campaign{
		Name:        "Aman",
		LastName:    "abc",
		Description: "This is an example campaign",
		// Age:         &age,
		Address: Address{
			AddressMeta: AddressMeta{
				HouseNo: "5",
				StNo:    "12",
			},
			Street: "hsr",
			City:   "blr",
		},
	}

	updateMap := StructToMap(campaign)
	updateMap["last_name"] = "singh"
	// updateMap["address.city"]="Banglore"
	// delete(updateMap, "address")
	fmt.Println(updateMap)
	fmt.Printf("address: %+v\n", updateMap["address"])

	updateMap = StructToMapV2(campaign)
	fmt.Printf("%+v\n", updateMap)
	fmt.Printf("address: %+v\n", updateMap["address"])
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

func StructToMapV2(item interface{}) map[string]interface{} {
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
		if fieldType == reflect.Struct {
			nested := StructToMapV2(field) // <-- use recursive self
			if tag == "" {                 // inline struct
				for k, val := range nested {
					res[k] = val
				}
			} else {
				res[tag] = nested
			}
			continue // only to avoid setting struct (as it is) in the result map
		}
		if tag == "" || tag == "-" {
			continue
		}
		if !IsZeroVal(field) || fieldType == reflect.Bool {
			res[tag] = field
		}
	}
	return res
}
