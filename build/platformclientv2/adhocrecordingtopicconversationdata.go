package platformclientv2
import (
	"encoding/json"
)

// Adhocrecordingtopicconversationdata
type Adhocrecordingtopicconversationdata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicconversationdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
