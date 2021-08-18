package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentupdate
type Learningassignmentupdate struct { 
	// State - The Learning Assignment state
	State *string `json:"state,omitempty"`


	// Assessment - An updated Assessment
	Assessment *Learningassessment `json:"assessment,omitempty"`

}

func (u *Learningassignmentupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentupdate

	

	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		Assessment *Learningassessment `json:"assessment,omitempty"`
		*Alias
	}{ 
		State: u.State,
		
		Assessment: u.Assessment,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
