package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentbulkaddresponse
type Learningassignmentbulkaddresponse struct { 
	// Entities - The learning assignments that were assigned correctly
	Entities *[]Learningassignment `json:"entities,omitempty"`


	// DisallowedEntities - The items that were not allowed to be assigned
	DisallowedEntities *[]Disallowedentitylearningassignmentitem `json:"disallowedEntities,omitempty"`

}

func (u *Learningassignmentbulkaddresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentbulkaddresponse

	

	return json.Marshal(&struct { 
		Entities *[]Learningassignment `json:"entities,omitempty"`
		
		DisallowedEntities *[]Disallowedentitylearningassignmentitem `json:"disallowedEntities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		DisallowedEntities: u.DisallowedEntities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentbulkaddresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
