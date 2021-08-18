package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Change
type Change struct { 
	// Entity
	Entity *Auditentity `json:"entity,omitempty"`


	// Property - The property that was changed
	Property *string `json:"property,omitempty"`


	// OldValues - The old values which were modified and/or removed by this action.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - The new values which were modified and/or added by this action.
	NewValues *[]string `json:"newValues,omitempty"`

}

func (u *Change) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Change

	

	return json.Marshal(&struct { 
		Entity *Auditentity `json:"entity,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		Entity: u.Entity,
		
		Property: u.Property,
		
		OldValues: u.OldValues,
		
		NewValues: u.NewValues,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Change) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
