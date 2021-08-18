package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangecontactsort
type Dialercampaignconfigchangecontactsort struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Numeric
	Numeric *bool `json:"numeric,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercampaignconfigchangecontactsort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangecontactsort

	

	return json.Marshal(&struct { 
		FieldName *string `json:"fieldName,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Numeric *bool `json:"numeric,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		FieldName: u.FieldName,
		
		Direction: u.Direction,
		
		Numeric: u.Numeric,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecontactsort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
