package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Reschedulingmanagementunitresponse
type Reschedulingmanagementunitresponse struct { 
	// ManagementUnit - The management unit
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// Applied - Whether the rescheduling run is applied for the given management unit
	Applied *bool `json:"applied,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reschedulingmanagementunitresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
