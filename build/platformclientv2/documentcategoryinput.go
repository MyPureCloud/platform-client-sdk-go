package platformclientv2
import (
	"encoding/json"
)

// Documentcategoryinput
type Documentcategoryinput struct { 
	// Id - KnowledgeBase Category ID
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentcategoryinput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
