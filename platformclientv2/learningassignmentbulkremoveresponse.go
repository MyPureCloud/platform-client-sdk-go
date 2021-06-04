package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Learningassignmentbulkremoveresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
