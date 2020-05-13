package platformclientv2
import (
	"encoding/json"
)

// Messagingtemplaterequest
type Messagingtemplaterequest struct { 
	// ResponseId - A Response Management response identifier for a messaging template defined response
	ResponseId *string `json:"responseId,omitempty"`


	// Parameters - A list of Response Management response substitutions for the response's messaging template
	Parameters *[]Templateparameter `json:"parameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagingtemplaterequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
