package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/threatcl/go-otm/pkg/otm"
)

func main() {

	// The below is an example github.com/xntrik/hcltm threat model

	// spec_version = "0.1.6"
	//
	// threatmodel "Tower of London" {
	//   description = "A historic castle"
	//   author = "@xntrik"
	//
	//   attributes {
	//     new_initiative = "true"
	//     internet_facing = "true"
	//     initiative_size = "Small"
	//   }
	//
	//   additional_attribute "network_segment" {
	//     value = "dmz"
	//   }
	//
	//   information_asset "crown jewels" {
	//     description = "including the imperial state crown"
	//     information_classification = "Confidential"
	//   }
	//
	//   usecase {
	//     description = "The Queen can fetch the crown"
	//   }
	//
	//   third_party_dependency "community watch" {
	//     description = "The community watch helps guard the premise"
	//     uptime_dependency = "degraded"
	//   }
	//
	//   threat {
	//     description = "Someone who isn't the Queen steals the crown"
	//     impacts = ["Confidentiality"]
	//
	//     expanded_control "Lots of Guards" {
	//       implemented = true
	//       description = "Lots of guards patrol the area"
	//       implementation_notes = "They are trained to be guards as well"
	//       risk_reduction = 80
	//     }
	//   }
	//
	// }

	// Now lets use the github.com/threatcl/go-otm/pkg/otm to create an OTM struct
	myOtm := otm.OtmSchemaJson{}

	myOtm.OtmVersion = "0.2.0"

	myOtm.Project.Name = "Tower of London"
	myOtm.Project.Id = "tower-of-london"
	description := "A historic castle"
	myOtm.Project.Description = &description
	owner := "@xntrik"
	myOtm.Project.Owner = &owner
	myOtm.Project.Attributes = map[string]interface{}{
		"new_initiative":  "true",
		"internet_facing": "true",
		"initiative_size": "Small",
		"network_segment": "dmz",
	}

	assetDescription := "including the imperial state crown"
	asset := otm.OtmSchemaJsonAssetsElem{
		Name:        "crown jewels",
		Id:          "crown-jewels",
		Description: &assetDescription,
		Attributes: map[string]interface{}{
			"information_classification": "Confidential",
		},
	}

	myOtm.Assets = append(myOtm.Assets, asset)

	infoDisclosure := "Information Disclosure"
	threatDescription := "Someone who isn't the Queen steals the Crown"
	threat := otm.OtmSchemaJsonThreatsElem{
		Description: &threatDescription,
		Id:          "crown-theft",
		Name:        "Crown theft",
		Categories: []*string{
			&infoDisclosure,
		},
	}

	myOtm.Threats = append(myOtm.Threats, threat)

	mitigationDescription := "They are trained to be guards as well"
	mitigation := otm.OtmSchemaJsonMitigationsElem{
		Name:          "Lots of Guards",
		Id:            "lots-of-guards",
		Description:   &mitigationDescription,
		RiskReduction: 80,
		Attributes: map[string]interface{}{
			"implemented": "true",
		},
	}

	myOtm.Mitigations = append(myOtm.Mitigations, mitigation)

	// Marshal into a JSON byte slice
	jsonOut, err := json.Marshal(myOtm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Finally we print it out
	fmt.Printf("%s\n", jsonOut)
}
