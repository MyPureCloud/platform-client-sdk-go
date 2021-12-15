package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangefilterclause
type Dialercontactlistfilterconfigchangefilterclause struct { 
	// FilterType - Contact list filter type
	FilterType *string `json:"filterType,omitempty"`


	// Predicates - The list of predicates in that clause
	Predicates *[]Dialercontactlistfilterconfigchangefilterpredicate `json:"predicates,omitempty"`

}

func (o *Dialercontactlistfilterconfigchangefilterclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangefilterclause
	
	return json.Marshal(&struct { 
		FilterType *string `json:"filterType,omitempty"`
		
		Predicates *[]Dialercontactlistfilterconfigchangefilterpredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		FilterType: o.FilterType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangefilterclause) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangefilterclauseMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangefilterclauseMap)
	if err != nil {
		return err
	}
	
	if FilterType, ok := DialercontactlistfilterconfigchangefilterclauseMap["filterType"].(string); ok {
		o.FilterType = &FilterType
	}
	
	if Predicates, ok := DialercontactlistfilterconfigchangefilterclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
