package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Objectivezone
type Objectivezone struct { 
	// Label - label
	Label *string `json:"label,omitempty"`


	// DirectionType - direction type
	DirectionType *string `json:"directionType,omitempty"`


	// ZoneType - zone type
	ZoneType *string `json:"zoneType,omitempty"`


	// UpperLimitPoints - upper limit points
	UpperLimitPoints *int `json:"upperLimitPoints,omitempty"`


	// LowerLimitPoints - lower limit points
	LowerLimitPoints *int `json:"lowerLimitPoints,omitempty"`


	// UpperLimitValue - upper limit value
	UpperLimitValue *int `json:"upperLimitValue,omitempty"`


	// LowerLimitValue - lower limit value
	LowerLimitValue *int `json:"lowerLimitValue,omitempty"`

}

func (u *Objectivezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Objectivezone

	

	return json.Marshal(&struct { 
		Label *string `json:"label,omitempty"`
		
		DirectionType *string `json:"directionType,omitempty"`
		
		ZoneType *string `json:"zoneType,omitempty"`
		
		UpperLimitPoints *int `json:"upperLimitPoints,omitempty"`
		
		LowerLimitPoints *int `json:"lowerLimitPoints,omitempty"`
		
		UpperLimitValue *int `json:"upperLimitValue,omitempty"`
		
		LowerLimitValue *int `json:"lowerLimitValue,omitempty"`
		*Alias
	}{ 
		Label: u.Label,
		
		DirectionType: u.DirectionType,
		
		ZoneType: u.ZoneType,
		
		UpperLimitPoints: u.UpperLimitPoints,
		
		LowerLimitPoints: u.LowerLimitPoints,
		
		UpperLimitValue: u.UpperLimitValue,
		
		LowerLimitValue: u.LowerLimitValue,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Objectivezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
