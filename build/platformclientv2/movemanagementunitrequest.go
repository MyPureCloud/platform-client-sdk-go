package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Movemanagementunitrequest
type Movemanagementunitrequest struct { 
	// BusinessUnitId - The ID of the business unit to which to move the management unit
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Movemanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
