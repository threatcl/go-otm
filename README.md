# go-otm

`go-otm` is a golang package that implements a go schema of the Open Threat Model (OTM) Format. https://github.com/iriusrisk/OpenThreatModel

> **Note** This was primarily constructed to allow github.com/xntrik/hcltm to export OTM JSON schema exports.

## Manual update steps

Currently this spec is generated with the github.com/atombendor/go-jsonschema cli tool. I haven't yet bothered building this in github actions, so these are the manual steps I follow to update this repo.

1. Install github.com/atombender/go-jsonschema
2. Clone github.com/iriusrisk/OpenThreatModel somewhere
3. From the root of the go-otm repo run: `go-jsonschema -p otm -o pkg/otm/otm.go -v ../../iriusrisk/OpenThreatModel/otm_schema.json`
4. Check the diff
5. Copy the `otm_schema.json` into `pkg/otm/testdata/`
6. Make sure that the `make test` tests pass - you _may_ need to update the `pkg/otm/testdata/EXAMPLE.json` file from the `OpenThreatModel` repo
