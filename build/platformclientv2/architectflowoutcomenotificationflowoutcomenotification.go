package platformclientv2
import (
	"encoding/json"
)

// Architectflowoutcomenotificationflowoutcomenotification
type Architectflowoutcomenotificationflowoutcomenotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectflowoutcomenotificationarchitectoperation `json:"currentOperation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationflowoutcomenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
