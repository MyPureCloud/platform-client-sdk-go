package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentbulkremoveresponse
type Learningassignmentbulkremoveresponse struct { 
	// Entities - The learning assignments that were removed successfully
	Entities *[]Learningassignmententity `json:"entities,omitempty"`


	// DisallowedEntities - The learning assignments that were not removed due to missing permissions
	DisallowedEntities *[]Disallowedentitylearningassignmentreference `json:"disallowedEntities,omitempty"`

}

func (u *Learningassignmentbulkremoveresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentbulkremoveresponse

	

	return json.Marshal(&struct { 
		Entities *[]Learningassignmententity `json:"entities,omitempty"`
		
		DisallowedEntities *[]Disallowedentitylearningassignmentreference `json:"disallowedEntities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		DisallowedEntities: u.DisallowedEntities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentbulkremoveresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
