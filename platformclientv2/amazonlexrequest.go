package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Amazonlexrequest
type Amazonlexrequest struct { 
	// RequestAttributes - AttributeName/AttributeValue pairs of User Defined Request Attributes to be sent to the amazon bot See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs
	RequestAttributes *map[string]string `json:"requestAttributes,omitempty"`


	// SessionAttributes - AttributeName/AttributeValue pairs of Session Attributes to be sent to the amazon bot. See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs
	SessionAttributes *map[string]string `json:"sessionAttributes,omitempty"`

}

func (u *Amazonlexrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Amazonlexrequest

	

	return json.Marshal(&struct { 
		RequestAttributes *map[string]string `json:"requestAttributes,omitempty"`
		
		SessionAttributes *map[string]string `json:"sessionAttributes,omitempty"`
		*Alias
	}{ 
		RequestAttributes: u.RequestAttributes,
		
		SessionAttributes: u.SessionAttributes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Amazonlexrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
