package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningassignmentrulerunnotification
type Learningassignmentruleruntopiclearningassignmentrulerunnotification struct { 
	// Entities
	Entities *[]Learningassignmentruleruntopiclearningassignmentscreated `json:"entities,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`

}

func (u *Learningassignmentruleruntopiclearningassignmentrulerunnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentruleruntopiclearningassignmentrulerunnotification

	

	return json.Marshal(&struct { 
		Entities *[]Learningassignmentruleruntopiclearningassignmentscreated `json:"entities,omitempty"`
		
		Total *int `json:"total,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		Total: u.Total,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningassignmentrulerunnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
