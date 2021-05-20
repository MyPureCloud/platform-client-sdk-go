package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangefilterclause
type Dialercontactlistfilterconfigchangefilterclause struct { 
	// FilterType
	FilterType *string `json:"filterType,omitempty"`


	// Predicates
	Predicates *[]Dialercontactlistfilterconfigchangefilterpredicate `json:"predicates,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
