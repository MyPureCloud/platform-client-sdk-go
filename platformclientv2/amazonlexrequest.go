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

func (o *Amazonlexrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Amazonlexrequest
	
	return json.Marshal(&struct { 
		RequestAttributes *map[string]string `json:"requestAttributes,omitempty"`
		
		SessionAttributes *map[string]string `json:"sessionAttributes,omitempty"`
		*Alias
	}{ 
		RequestAttributes: o.RequestAttributes,
		
		SessionAttributes: o.SessionAttributes,
		Alias:    (*Alias)(o),
	})
}

func (o *Amazonlexrequest) UnmarshalJSON(b []byte) error {
	var AmazonlexrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AmazonlexrequestMap)
	if err != nil {
		return err
	}
	
	if RequestAttributes, ok := AmazonlexrequestMap["requestAttributes"].(map[string]interface{}); ok {
		RequestAttributesString, _ := json.Marshal(RequestAttributes)
		json.Unmarshal(RequestAttributesString, &o.RequestAttributes)
	}
	
	if SessionAttributes, ok := AmazonlexrequestMap["sessionAttributes"].(map[string]interface{}); ok {
		SessionAttributesString, _ := json.Marshal(SessionAttributes)
		json.Unmarshal(SessionAttributesString, &o.SessionAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Amazonlexrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
