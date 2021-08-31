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

func (o *Licensebatchassignmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licensebatchassignmentrequest
	
	return json.Marshal(&struct { 
		Assignments *[]Licenseassignmentrequest `json:"assignments,omitempty"`
		*Alias
	}{ 
		Assignments: o.Assignments,
		Alias:    (*Alias)(o),
	})
}

func (o *Licensebatchassignmentrequest) UnmarshalJSON(b []byte) error {
	var LicensebatchassignmentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LicensebatchassignmentrequestMap)
	if err != nil {
		return err
	}
	
	if Assignments, ok := LicensebatchassignmentrequestMap["assignments"].([]interface{}); ok {
		AssignmentsString, _ := json.Marshal(Assignments)
		json.Unmarshal(AssignmentsString, &o.Assignments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Licensebatchassignmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
