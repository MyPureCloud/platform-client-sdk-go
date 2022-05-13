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

func (o *Objectivezone) MarshalJSON() ([]byte, error) {
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
		Label: o.Label,
		
		DirectionType: o.DirectionType,
		
		ZoneType: o.ZoneType,
		
		UpperLimitPoints: o.UpperLimitPoints,
		
		LowerLimitPoints: o.LowerLimitPoints,
		
		UpperLimitValue: o.UpperLimitValue,
		
		LowerLimitValue: o.LowerLimitValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Objectivezone) UnmarshalJSON(b []byte) error {
	var ObjectivezoneMap map[string]interface{}
	err := json.Unmarshal(b, &ObjectivezoneMap)
	if err != nil {
		return err
	}
	
	if Label, ok := ObjectivezoneMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if DirectionType, ok := ObjectivezoneMap["directionType"].(string); ok {
		o.DirectionType = &DirectionType
	}
    
	if ZoneType, ok := ObjectivezoneMap["zoneType"].(string); ok {
		o.ZoneType = &ZoneType
	}
    
	if UpperLimitPoints, ok := ObjectivezoneMap["upperLimitPoints"].(float64); ok {
		UpperLimitPointsInt := int(UpperLimitPoints)
		o.UpperLimitPoints = &UpperLimitPointsInt
	}
	
	if LowerLimitPoints, ok := ObjectivezoneMap["lowerLimitPoints"].(float64); ok {
		LowerLimitPointsInt := int(LowerLimitPoints)
		o.LowerLimitPoints = &LowerLimitPointsInt
	}
	
	if UpperLimitValue, ok := ObjectivezoneMap["upperLimitValue"].(float64); ok {
		UpperLimitValueInt := int(UpperLimitValue)
		o.UpperLimitValue = &UpperLimitValueInt
	}
	
	if LowerLimitValue, ok := ObjectivezoneMap["lowerLimitValue"].(float64); ok {
		LowerLimitValueInt := int(LowerLimitValue)
		o.LowerLimitValue = &LowerLimitValueInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Objectivezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
