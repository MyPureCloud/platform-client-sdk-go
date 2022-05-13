package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediautilization
type Mediautilization struct { 
	// MaximumCapacity - Defines the maximum number of conversations of this type that an agent can handle at one time.
	MaximumCapacity *int `json:"maximumCapacity,omitempty"`


	// InterruptableMediaTypes - Defines the list of other media types that can interrupt a conversation of this media type.  Values include call, chat, email, callback, and message.
	InterruptableMediaTypes *[]string `json:"interruptableMediaTypes,omitempty"`


	// IncludeNonAcd - If true, then track non-ACD conversations against utilization
	IncludeNonAcd *bool `json:"includeNonAcd,omitempty"`

}

func (o *Mediautilization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediautilization
	
	return json.Marshal(&struct { 
		MaximumCapacity *int `json:"maximumCapacity,omitempty"`
		
		InterruptableMediaTypes *[]string `json:"interruptableMediaTypes,omitempty"`
		
		IncludeNonAcd *bool `json:"includeNonAcd,omitempty"`
		*Alias
	}{ 
		MaximumCapacity: o.MaximumCapacity,
		
		InterruptableMediaTypes: o.InterruptableMediaTypes,
		
		IncludeNonAcd: o.IncludeNonAcd,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediautilization) UnmarshalJSON(b []byte) error {
	var MediautilizationMap map[string]interface{}
	err := json.Unmarshal(b, &MediautilizationMap)
	if err != nil {
		return err
	}
	
	if MaximumCapacity, ok := MediautilizationMap["maximumCapacity"].(float64); ok {
		MaximumCapacityInt := int(MaximumCapacity)
		o.MaximumCapacity = &MaximumCapacityInt
	}
	
	if InterruptableMediaTypes, ok := MediautilizationMap["interruptableMediaTypes"].([]interface{}); ok {
		InterruptableMediaTypesString, _ := json.Marshal(InterruptableMediaTypes)
		json.Unmarshal(InterruptableMediaTypesString, &o.InterruptableMediaTypes)
	}
	
	if IncludeNonAcd, ok := MediautilizationMap["includeNonAcd"].(bool); ok {
		o.IncludeNonAcd = &IncludeNonAcd
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mediautilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
