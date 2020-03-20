package platformclientv2
import (
	"encoding/json"
)

// Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping
type Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping struct { 
	// ContactColumnName
	ContactColumnName *string `json:"contactColumnName,omitempty"`


	// DataActionField
	DataActionField *string `json:"dataActionField,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
