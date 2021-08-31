package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilterrange
type Contactlistfilterrange struct { 
	// Min - The minimum value of the range. Required for the operator BETWEEN.
	Min *string `json:"min,omitempty"`


	// Max - The maximum value of the range. Required for the operator BETWEEN.
	Max *string `json:"max,omitempty"`


	// MinInclusive - Whether or not to include the minimum in the range.
	MinInclusive *bool `json:"minInclusive,omitempty"`


	// MaxInclusive - Whether or not to include the maximum in the range.
	MaxInclusive *bool `json:"maxInclusive,omitempty"`


	// InSet - A set of values that the contact data should be in. Required for the IN operator.
	InSet *[]string `json:"inSet,omitempty"`

}

func (o *Contactlistfilterrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistfilterrange
	
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

func (o *Contactlistfilterrange) UnmarshalJSON(b []byte) error {
	var ContactlistfilterrangeMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistfilterrangeMap)
	if err != nil {
		return err
	}
	
	if Min, ok := ContactlistfilterrangeMap["min"].(string); ok {
		o.Min = &Min
	}
	
	if Max, ok := ContactlistfilterrangeMap["max"].(string); ok {
		o.Max = &Max
	}
	
	if MinInclusive, ok := ContactlistfilterrangeMap["minInclusive"].(bool); ok {
		o.MinInclusive = &MinInclusive
	}
	
	if MaxInclusive, ok := ContactlistfilterrangeMap["maxInclusive"].(bool); ok {
		o.MaxInclusive = &MaxInclusive
	}
	
	if InSet, ok := ContactlistfilterrangeMap["inSet"].([]interface{}); ok {
		InSetString, _ := json.Marshal(InSet)
		json.Unmarshal(InSetString, &o.InSet)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistfilterrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
