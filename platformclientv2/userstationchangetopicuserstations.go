package platformclientv2
import (
	"encoding/json"
)

// Userstationchangetopicuserstations
type Userstationchangetopicuserstations struct { 
	// AssociatedStation
	AssociatedStation *Userstationchangetopicuserstation `json:"associatedStation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstations) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
