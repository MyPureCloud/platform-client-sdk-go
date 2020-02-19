package platformclientv2
import (
	"encoding/json"
)

// Userstationchangetopicuser
type Userstationchangetopicuser struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
