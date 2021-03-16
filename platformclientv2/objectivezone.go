package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Objectivezone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
