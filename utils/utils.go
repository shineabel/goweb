package utils

import (
	"log"
	"reflect"
	"fmt"
)

func setField(o interface{}, name string, value interface{}) error  {

	structData := reflect.ValueOf(o).Elem()
	fieldValue := structData.FieldByName(name)
	if !fieldValue.IsValid(){
		fmt.Errorf("no such field %s in obj", name)
	}
	if !fieldValue.CanSet() {
		fmt.Errorf("can not set %s field value ", name)
	}

	//fieldType := fieldValue.Type()
	//val := reflect.ValueOf(value)
	//
	//
	//
	//valTypeStr := value.
	return nil

}

func parse(obj interface{}, m map[string]interface{}) error  {

	for k, v := range m {

		if err := setField(obj, k, v); err != nil {
			log.Printf("error:", err)
			return err
		}
	}
	return nil
}
