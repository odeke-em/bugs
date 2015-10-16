package main

import (
	"fmt"
	"reflect"
	"strings"
)

type pushCmd struct {
	Hidden         *bool   `json:"hidden"`
	Source         *string `json:"source"`
	Convert        *bool   `json:"convert"`
	Ocr            *bool   `json:"ocr"`
	IgnoreChecksum *bool   `json:"ignore-checksum"`
}

func siftCliTags(from interface{}, alreadyDefined map[string]bool, defaults map[string]interface{}) string {
	fromValue := reflect.ValueOf(from)
	fromType := reflect.TypeOf(from)

	mapping := map[string]string{}

	for i, n := 0, fromType.NumField(); i < n; i++ {
		fromFieldT := fromType.Field(i)
		fromTag := fromFieldT.Tag.Get("json")

		fmt.Println("fromTag", fromTag, alreadyDefined)
		if fromTag == "" {
			continue
		}

		fromFieldV := fromValue.Field(i)

		elem := fromFieldV.Elem()

		if _, defined := alreadyDefined[fromTag]; !defined {
			if retr, defaultSet := defaults[fromTag]; defaultSet {
				elem = reflect.ValueOf(retr)
			}
		}

		stringified := ""
		switch elem.Kind() {
		case reflect.String:
			stringified = fmt.Sprintf("%q", elem)
		default:
			stringified = fmt.Sprintf("%v", elem)
		}

		fmt.Println("stringified", stringified)
		mapping[fromTag] = stringified
	}

	joined := []string{}

	for k, v := range mapping {
		joined = append(joined, fmt.Sprintf("%q:%v", k, v))
	}

	stringified := strings.Join(joined, ",")

	return fmt.Sprintf("{%v}", stringified)
}

func main() {
	hidden := true
	ocr := false
	ignoreChecksum := true
	convert := true
	source := "/home/emmanuel/drivemount"
	p := &pushCmd{
		Ocr:            &ocr,
		Hidden:         &hidden,
		Convert:        &convert,
		IgnoreChecksum: &ignoreChecksum,
		Source:         &source,
	}

	alreadyDefined := map[string]bool{
		"Hidden": true,
	}
	defaults := map[string]interface{}{}

	stringified := siftCliTags(*p, alreadyDefined, defaults)
	fmt.Println(stringified)
}
