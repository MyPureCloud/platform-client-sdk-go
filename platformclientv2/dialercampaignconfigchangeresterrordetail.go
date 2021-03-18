package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangeresterrordetail
type Dialercampaignconfigchangeresterrordetail struct { 
	// VarError
	VarError *string `json:"error,omitempty"`


	// Details
	Details *string `json:"details,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangeresterrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
