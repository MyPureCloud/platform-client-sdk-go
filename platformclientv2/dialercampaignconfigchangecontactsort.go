package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangecontactsort
type Dialercampaignconfigchangecontactsort struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Numeric
	Numeric *bool `json:"numeric,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecontactsort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
