package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Licensebatchassignmentrequest
type Licensebatchassignmentrequest struct { 
	// Assignments - The list of license assignment updates to make.
	Assignments *[]Licenseassignmentrequest `json:"assignments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Licensebatchassignmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
