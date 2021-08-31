package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buttonresponse
type Buttonresponse struct { 
	// VarType - Button response type that captures Button and QuickReply type responses
	VarType *string `json:"type,omitempty"`


	// Text - Text to show inside the Button reply. This is also used as the response text after clicking on the Button.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a button
	Payload *string `json:"payload,omitempty"`

}

func (o *Buttonresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buttonresponse
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		Alias:    (*Alias)(o),
	})
}

func (o *Buttonresponse) UnmarshalJSON(b []byte) error {
	var ButtonresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ButtonresponseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ButtonresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := ButtonresponseMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Payload, ok := ButtonresponseMap["payload"].(string); ok {
		o.Payload = &Payload
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
