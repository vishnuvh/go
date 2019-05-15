package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
)

func main() {
	var data interface{}
	if err := json.NewDecoder(strings.NewReader(src)).Decode(&data); err != nil {
		panic(err)
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	val, err := gabs.ParseJSON(dataJSON)
	prefix := val.Search("assetIdPrefix")

	//extract array
	arrays := val.Search("adiMap", "ItemBlob", "ADI", "Asset", "Metadata", "App_Data", "Category").Data()
	categories := arrayToString(arrays.([]interface{}))

	// custom date format
	layout := "2006-01-02T15:04:05"
	date, _ := time.Parse(layout, val.Search("adiMap", "ItemBlob", "ADI", "Asset", "Metadata", "App_Data", "Licensing_Window_End").Data().(string))

	//standard date format
	date2, _ := time.Parse(time.RFC3339, val.Search("adiMap", "itemArrivalDate").Data().(string))

	fmt.Println("prefix : ", prefix)
	fmt.Println("categories :", categories)
	fmt.Println("date :", date)
	fmt.Println("date2 :", date2)

}

func arrayToString(strArray []interface{}) string {
	values := []string{}
	for _, v := range strArray {
		values = append(values, v.(string))
	}
	return strings.Join(values, " ")
}

const src = `{
	"assetIdPrefix": "CCDN",
	"adiMap": {
		"itemArrivalDate": "2019-05-02T20:31:30.257Z",
		"ItemBlob": {
			"ADI": {
				"Asset": {
					"Metadata": {
						"App_Data": {
							"Category": ["cat1", "cat2", "cat3"],
							"Licensing_Window_End": "2019-07-09T23:59:59"
						}
					}
				}
			}

		}
	}
}`