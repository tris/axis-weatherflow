package axis

import (
	"reflect"
	"strconv"
)

func structToPairs(s interface{}) []string {
	var pairs []string
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.IsNil() {
			continue
		}

		switch field.Kind() {
		case reflect.Ptr:
			field = field.Elem()
			switch field.Kind() {
			case reflect.String:
				pairs = append(pairs, fieldType.Name, field.String())
			case reflect.Int:
				pairs = append(pairs, fieldType.Name, strconv.Itoa(int(field.Int())))
			case reflect.Bool:
				pairs = append(pairs, fieldType.Name, strconv.FormatBool(field.Bool()))
			}
		}
	}

	return pairs
}

func structToArrayOfPairs(s interface{}) map[string]string {
	pairs := structToPairs(s)

	data := make(map[string]string)
	for i := 0; i < len(pairs); i += 2 {
		data[pairs[i]] = pairs[i+1]
	}

	return data
}
