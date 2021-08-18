package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsproperty
type Analyticsproperty struct { 
	// Property - User-defined rather than intrinsic system-observed values. These are tagged onto segments by other components within PureCloud or by API users directly.  This is the name of the user-defined property.
	Property *string `json:"property,omitempty"`


	// PropertyType - Indicates what the data type is (e.g. integer vs string) and therefore how to evaluate what would constitute a match
	PropertyType *string `json:"propertyType,omitempty"`


	// Value - What property value to match against
	Value *string `json:"value,omitempty"`

}

func (u *Analyticsproperty) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsproperty

	

	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Property: u.Property,
		
		PropertyType: u.PropertyType,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsproperty) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
