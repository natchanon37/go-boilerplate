package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func ConvertInterfaceToStruct(data interface{}, obj interface{}) {
	// Convert map to json string
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, obj); err != nil {
		fmt.Printf("%v", err)
	}
}

func ConvertInt64SliceToUintSlice(int64Slice []int64) []uint {
	uintSlice := make([]uint, len(int64Slice))
	for i, val := range int64Slice {
		if val >= 0 {
			uintSlice[i] = uint(val)
		}
	}
	return uintSlice
}

func StringToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}

func StringToInt(str string) int {
	b, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return b
}

func ConvertUintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func ConvertInterfaceToBytes(i interface{}) []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Error converting interface to bytes: %v", err)
	}
	return data
}

func ConvertByteArrayToInterface(i []byte, o interface{}) {
	if err := json.Unmarshal(i, &o); err != nil {
		log.Fatalf("Error converting byte array to interface: %v", err)
	}
}

func ConvertStructToMap(obj interface{}, isCheckZero bool) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := structToMapRecursive(result, reflect.ValueOf(obj), "", 0, 3, isCheckZero)
	return result, err
}

func ConvertTo2Decimal(value float64) float64 {
	return math.Trunc(math.Round(value*100)) / 100
}

func isZeroValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return value.Len() == 0
	default:
		zeroValue := reflect.Zero(value.Type())
		return reflect.DeepEqual(value.Interface(), zeroValue.Interface())
	}
}

func structToMapRecursive(
	result map[string]interface{},
	val reflect.Value,
	prefix string,
	depth, maxDepth int,
	isCheckZero bool,
) error {
	if depth > maxDepth {
		return nil
	}

	switch val.Kind() {
	case reflect.Struct:
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name

			// If there's a prefix, add it to the field name
			if prefix != "" {
				fieldName = fmt.Sprintf("%s.%s", prefix, fieldName)
			}

			if err := structToMapRecursive(result, field, fieldName, depth+1, maxDepth, isCheckZero); err != nil {
				return err
			}
		}

	case reflect.Ptr:
		if !val.IsNil() {
			return structToMapRecursive(result, val.Elem(), prefix, depth, maxDepth, isCheckZero)
		}

	case reflect.Interface:
		if !val.IsNil() {
			return structToMapRecursive(result, val.Elem(), prefix, depth, maxDepth, isCheckZero)
		}

	default:
		if isZeroValue(val) && isCheckZero {
			return nil
		}
		result[prefix] = val.Interface()
	}

	return nil
}

func SplitName(fullName string, allowEmpty bool) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		if allowEmpty {
			return "", ""
		} else {
			return "-", "-"
		}
	}
	firstName := parts[0]
	lastName := ""
	if len(parts) > 1 {
		lastName = strings.Join(parts[1:], " ")
	}

	if !allowEmpty && firstName == "" {
		firstName = "-"
	}

	if !allowEmpty && lastName == "" {
		lastName = "-"
	}

	return firstName, lastName
}

func PrintJsonFromStruct(obj interface{}) {
	marshaled, _ := json.MarshalIndent(obj, "", "   ")

	fmt.Println(string(marshaled))
}
