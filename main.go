package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wI2L/jsondiff"
	"github.com/xntrik/go-otm/pkg/otm"
)

func main() {
	// blep := otm.OpenThreatModelSpecification{}

	// fmt.Printf("%+v\n", blep)

	// var newBlep otm.OpenThreatModelSpecification
	var newBlep otm.OtmSchemaJson

	jsonFile, err := os.Open("EXAMPLE.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	byteVals, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	newBlep.UnmarshalJSON(byteVals)

	// fmt.Printf("%+v\n=======\n", newBlep)

	jsonOut, err := json.Marshal(newBlep)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", jsonOut)

	// fileOut, err := newBlep.MarshalJSON()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	_ = ioutil.WriteFile("test.json", jsonOut, 0644)
	//
	// os.Stdout.Write(fileOut)

	// patch, err := jsondiff.Compare(jsonFile, fileOut)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// // b, err := json.MarshalIndent(patch, "", "  ")
	// _, err = json.MarshalIndent(patch, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%+v\n", b)
	// os.Stdout.Write(b)

	patch, err := jsondiff.CompareJSON(byteVals, jsonOut, jsondiff.Equivalent())
	// patch, err = jsondiff.CompareJSON(byteVals, fileOut)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("=====\n%+v\n", patch)

	b, err := json.MarshalIndent(patch, "", "  ")
	// _, err = json.MarshalIndent(patch, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Printf("%+v\n", b)
	os.Stdout.Write(b)

}
