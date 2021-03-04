package platformclientv2
import (
	"encoding/json"
)

// Programrequest
type Programrequest struct { 
	// Name - The program name
	Name *string `json:"name,omitempty"`


	// Description - The program description
	Description *string `json:"description,omitempty"`


	// TopicIds - The ids of topics associated to the program
	TopicIds *[]string `json:"topicIds,omitempty"`


	// Tags - The program tags
	Tags *[]string `json:"tags,omitempty"`

}

// String returns a JSON representation of the model
func (o *Programrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
