package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcardaction - CardAction Object
type Conversationcardaction struct { 
	// VarType - Describes the type of action.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - Text to be returned as the payload from a ButtonResponse when a button is clicked.
	Payload *string `json:"payload,omitempty"`


	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`

}

func (o *Conversationcardaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcardaction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcardaction) UnmarshalJSON(b []byte) error {
	var ConversationcardactionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcardactionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationcardactionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := ConversationcardactionMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Payload, ok := ConversationcardactionMap["payload"].(string); ok {
		o.Payload = &Payload
	}
	
	if Url, ok := ConversationcardactionMap["url"].(string); ok {
		o.Url = &Url
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcardaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
