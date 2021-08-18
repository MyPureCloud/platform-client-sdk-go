package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licensebatchassignmentrequest
type Licensebatchassignmentrequest struct { 
	// Assignments - The list of license assignment updates to make.
	Assignments *[]Licenseassignmentrequest `json:"assignments,omitempty"`

}

func (u *Licensebatchassignmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licensebatchassignmentrequest

	

	return json.Marshal(&struct { 
		Assignments *[]Licenseassignmentrequest `json:"assignments,omitempty"`
		*Alias
	}{ 
		Assignments: u.Assignments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Licensebatchassignmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
