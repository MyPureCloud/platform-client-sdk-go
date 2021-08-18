package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Propertychange
type Propertychange struct { 
	// Property - The property that was changed
	Property *string `json:"property,omitempty"`


	// OldValues - Previous values for the property.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - New values for the property.
	NewValues *[]string `json:"newValues,omitempty"`

}

func (u *Propertychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Propertychange

	

	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		Property: u.Property,
		
		OldValues: u.OldValues,
		
		NewValues: u.NewValues,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Propertychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
