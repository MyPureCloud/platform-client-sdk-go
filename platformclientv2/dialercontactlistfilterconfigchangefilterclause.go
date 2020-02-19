package platformclientv2
import (
	"encoding/json"
)

// Dialercontactlistfilterconfigchangefilterclause
type Dialercontactlistfilterconfigchangefilterclause struct { 
	// FilterType
	FilterType *string `json:"filterType,omitempty"`


	// Predicates
	Predicates *[]Dialercontactlistfilterconfigchangefilterpredicate `json:"predicates,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
