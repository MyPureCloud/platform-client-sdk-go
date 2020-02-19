package platformclientv2
import (
	"encoding/json"
)

// Answeroption
type Answeroption struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Value
	Value *int32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
