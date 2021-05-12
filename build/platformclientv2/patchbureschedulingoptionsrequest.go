package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Patchbureschedulingoptionsrequest
type Patchbureschedulingoptionsrequest struct { 
	// ManagementUnits - Per-management unit rescheduling options to update
	ManagementUnits *[]Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
