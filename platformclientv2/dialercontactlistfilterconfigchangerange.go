package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangerange
type Dialercontactlistfilterconfigchangerange struct { 
	// Min
	Min *string `json:"min,omitempty"`


	// Max
	Max *string `json:"max,omitempty"`


	// MinInclusive
	MinInclusive *bool `json:"minInclusive,omitempty"`


	// MaxInclusive
	MaxInclusive *bool `json:"maxInclusive,omitempty"`


	// InSet
	InSet *[]string `json:"inSet,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercontactlistfilterconfigchangerange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangerange

	

	return json.Marshal(&struct { 
		Min *string `json:"min,omitempty"`
		
		Max *string `json:"max,omitempty"`
		
		MinInclusive *bool `json:"minInclusive,omitempty"`
		
		MaxInclusive *bool `json:"maxInclusive,omitempty"`
		
		InSet *[]string `json:"inSet,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Min: u.Min,
		
		Max: u.Max,
		
		MinInclusive: u.MinInclusive,
		
		MaxInclusive: u.MaxInclusive,
		
		InSet: u.InSet,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangerange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
