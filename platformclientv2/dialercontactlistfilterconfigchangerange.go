package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangerange - FilterRange is one of the attributes of a FilterPredicate
type Dialercontactlistfilterconfigchangerange struct { 
	// Min - Minimum end of the range
	Min *string `json:"min,omitempty"`


	// Max - Maximum end of the range
	Max *string `json:"max,omitempty"`


	// MinInclusive - Whether or not to include the minimum in the range
	MinInclusive *bool `json:"minInclusive,omitempty"`


	// MaxInclusive - Whether or not to include the maximum in the range
	MaxInclusive *bool `json:"maxInclusive,omitempty"`


	// InSet - Elements that apply to the IN operator
	InSet *[]string `json:"inSet,omitempty"`

}

func (o *Dialercontactlistfilterconfigchangerange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangerange
	
	return json.Marshal(&struct { 
		Min *string `json:"min,omitempty"`
		
		Max *string `json:"max,omitempty"`
		
		MinInclusive *bool `json:"minInclusive,omitempty"`
		
		MaxInclusive *bool `json:"maxInclusive,omitempty"`
		
		InSet *[]string `json:"inSet,omitempty"`
		*Alias
	}{ 
		Min: o.Min,
		
		Max: o.Max,
		
		MinInclusive: o.MinInclusive,
		
		MaxInclusive: o.MaxInclusive,
		
		InSet: o.InSet,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangerange) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangerangeMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangerangeMap)
	if err != nil {
		return err
	}
	
	if Min, ok := DialercontactlistfilterconfigchangerangeMap["min"].(string); ok {
		o.Min = &Min
	}
	
	if Max, ok := DialercontactlistfilterconfigchangerangeMap["max"].(string); ok {
		o.Max = &Max
	}
	
	if MinInclusive, ok := DialercontactlistfilterconfigchangerangeMap["minInclusive"].(bool); ok {
		o.MinInclusive = &MinInclusive
	}
	
	if MaxInclusive, ok := DialercontactlistfilterconfigchangerangeMap["maxInclusive"].(bool); ok {
		o.MaxInclusive = &MaxInclusive
	}
	
	if InSet, ok := DialercontactlistfilterconfigchangerangeMap["inSet"].([]interface{}); ok {
		InSetString, _ := json.Marshal(InSet)
		json.Unmarshal(InSetString, &o.InSet)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangerange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
