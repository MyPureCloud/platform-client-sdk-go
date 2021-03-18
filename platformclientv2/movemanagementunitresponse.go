package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Movemanagementunitresponse
type Movemanagementunitresponse struct { 
	// BusinessUnit - The new business unit
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`


	// Status - The status of the move.  Will always be 'Processing' unless the Management Unit is already in the requested Business Unit in which case it will be 'Complete'
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Movemanagementunitresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
