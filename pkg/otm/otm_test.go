package otm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/wI2L/jsondiff"
	"github.com/xeipuuv/gojsonschema"
)

// This is used for validating JSON string against the schema
type SchemaValidator struct {
	Schema *gojsonschema.Schema
}

func NewSchemaValidator(specBytes []byte) (*SchemaValidator, error) {
	schemaFile := gojsonschema.NewStringLoader(string(specBytes))
	schema, err := gojsonschema.NewSchema(schemaFile)
	if err != nil {
		return nil, err
	}
	newLoader := &SchemaValidator{
		Schema: schema,
	}
	return newLoader, nil
}

func (sv *SchemaValidator) Validate(jsonBytes []byte) (bool, []string, error) {
	fileToVal := gojsonschema.NewStringLoader(string(jsonBytes))

	result, err := sv.Schema.Validate(fileToVal)
	if err != nil {
		return false, []string{}, err
	}

	if !result.Valid() {
		var errorString []string
		for _, s := range result.Errors() {
			errorString = append(errorString, s.String())
		}
		return false, errorString, nil
	}

	return true, []string{}, nil
}

func removeNullFields(data []byte) ([]byte, error) {
	var obj interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	cleanedObj := clean(obj)
	return json.Marshal(cleanedObj)
}

func clean(v interface{}) interface{} {
	switch x := v.(type) {
	case map[string]interface{}:
		for k, v := range x {
			if v == nil {
				delete(x, k) // Remove the field if the value is nil
			} else {
				x[k] = clean(v) // Recurse into the value
			}
		}
	case []interface{}:
		for i, v := range x {
			x[i] = clean(v) // Recurse into each element
		}
	}
	return v
}

func TestOtm(t *testing.T) {

	// Open the otm_schema.json from the upstream OpenThreatModel repo
	specFile, err := os.Open("./testdata/otm_schema.json")
	if err != nil {
		t.Fatalf("Error loading existing file: %s", err)
	}

	defer specFile.Close()

	// Read the file into a byte slice
	specBytes, err := ioutil.ReadAll(specFile)
	if err != nil {
		t.Fatalf("Error loading existing file: %s", err)
	}

	jsonValidator, err := NewSchemaValidator(specBytes)
	if err != nil {
		t.Fatalf("Error preparing schema validator: %s", err)
	}

	cases := []struct {
		name         string
		testFile     string
		expectedDiff int
	}{
		{
			"Original EXAMPLE file",
			"./testdata/EXAMPLE.json",
			3,
		},
		{
			"Modified EXAMPLE file",
			"./testdata/EXAMPLE2.json",
			0,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var newBlep OtmSchemaJson

			// Open the EXAMPLE.json files from the upstream OpenThreatModel repo
			// Although we have copies of these in our testdata
			jsonFile, err := os.Open(tc.testFile)
			if err != nil {
				t.Fatalf("Error loading existing file: %s", err)
			}

			defer jsonFile.Close()

			// Read the file into a byte slice
			byteVals, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				t.Fatalf("Error loading existing file: %s", err)
			}

			// Unmarshal the raw json bytes into the otm spec
			newBlep.UnmarshalJSON(byteVals)

			// Marshall this back into a byte slice
			jsonOut, err := json.Marshal(newBlep)
			if err != nil {
				t.Fatalf("Error loading existing file: %s", err)
			}

			// We remove the `null` attributes from the original JSON blob
			// When these get marshalled into the otmSchema the null entries are
			// ommitted
			newByteVals, err := removeNullFields(byteVals)
			if err != nil {
				t.Fatalf("Error loading existing file: %s", err)
			}

			patch, err := jsondiff.CompareJSON(newByteVals, jsonOut, jsondiff.Equivalent())
			if err != nil {
				t.Fatalf("Error loading existing file: %s", err)
			}

			if len(patch) != tc.expectedDiff {
				t.Errorf("There were %d items found in the patch when there should be %d", len(patch), tc.expectedDiff)

			}

			// Validate the resultant JSON against the schema
			targetJsonBytes := [][]byte{newByteVals, jsonOut}
			for _, bits := range targetJsonBytes {
				result, verboseErrors, err := jsonValidator.Validate(bits)
				if err != nil {
					t.Fatalf("Error validating the schema: %s", err)
				}

				if !result {
					t.Errorf("The json didn't match the schema: %+v", verboseErrors)
				}
			}

		})
	}
}
