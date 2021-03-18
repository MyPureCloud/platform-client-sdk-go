package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
