package platformclientv2
import (
	"encoding/json"
)

// Amazonlexrequest
type Amazonlexrequest struct { 
	// RequestAttributes - AttributeName/AttributeValue pairs of User Defined Request Attributes to be sent to the amazon bot See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs
	RequestAttributes *map[string]string `json:"requestAttributes,omitempty"`


	// SessionAttributes - AttributeName/AttributeValue pairs of Session Attributes to be sent to the amazon bot. See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs
	SessionAttributes *map[string]string `json:"sessionAttributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Amazonlexrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
