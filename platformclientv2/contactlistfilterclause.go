package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilterclause
type Contactlistfilterclause struct { 
	// FilterType - How to join predicates together.
	FilterType *string `json:"filterType,omitempty"`


	// Predicates - Conditions to filter the contacts by.
	Predicates *[]Contactlistfilterpredicate `json:"predicates,omitempty"`

}

func (o *Contactlistfilterclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistfilterclause
	
	return json.Marshal(&struct { 
		FilterType *string `json:"filterType,omitempty"`
		
		Predicates *[]Contactlistfilterpredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		FilterType: o.FilterType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactlistfilterclause) UnmarshalJSON(b []byte) error {
	var ContactlistfilterclauseMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistfilterclauseMap)
	if err != nil {
		return err
	}
	
	if FilterType, ok := ContactlistfilterclauseMap["filterType"].(string); ok {
		o.FilterType = &FilterType
	}
    
	if Predicates, ok := ContactlistfilterclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistfilterclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
