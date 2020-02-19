package platformclientv2
import (
	"encoding/json"
)

// Analyticsproperty
type Analyticsproperty struct { 
	// PropertyType - Indicates what the data type is (e.g. integer vs string) and therefore how to evaluate what would constitute a match
	PropertyType *string `json:"propertyType,omitempty"`


	// Property - User-defined rather than intrinsic system-observed values. These are tagged onto segments by other components within PureCloud or by API users directly.  This is the name of the user-defined property.
	Property *string `json:"property,omitempty"`


	// Value - What property value to match against
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsproperty) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
