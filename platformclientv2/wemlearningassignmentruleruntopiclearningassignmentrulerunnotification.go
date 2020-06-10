package platformclientv2
import (
	"encoding/json"
)

// Wemlearningassignmentruleruntopiclearningassignmentrulerunnotification
type Wemlearningassignmentruleruntopiclearningassignmentrulerunnotification struct { 
	// Entities
	Entities *[]Wemlearningassignmentruleruntopicwemlearningassignmentscreated `json:"entities,omitempty"`


	// Total
	Total *int32 `json:"total,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmentruleruntopiclearningassignmentrulerunnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
