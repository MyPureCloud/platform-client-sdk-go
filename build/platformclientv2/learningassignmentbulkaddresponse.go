package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Learningassignmentbulkaddresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
