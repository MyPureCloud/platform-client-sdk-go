package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Dialercontactlistfilterconfigchangefilterclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangefilterclause

	

	return json.Marshal(&struct { 
		FilterType *string `json:"filterType,omitempty"`
		
		Predicates *[]Dialercontactlistfilterconfigchangefilterpredicate `json:"predicates,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		FilterType: u.FilterType,
		
		Predicates: u.Predicates,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
