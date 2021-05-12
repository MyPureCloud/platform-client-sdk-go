package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Change) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
