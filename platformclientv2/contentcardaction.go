package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentcardaction - A card action that a user can take.
type Contentcardaction struct { 
	// VarType - Describes the type of action.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - Text to be returned as the payload from a ButtonResponse when a button is clicked. The payload and text are a combination which will have to be unique across each card and carousel in order to determine which button was clicked in that card or carousel.
	Payload *string `json:"payload,omitempty"`


	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`

}

func (o *Contentcardaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentcardaction
	
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

func (o *Contentcardaction) UnmarshalJSON(b []byte) error {
	var ContentcardactionMap map[string]interface{}
	err := json.Unmarshal(b, &ContentcardactionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ContentcardactionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := ContentcardactionMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Payload, ok := ContentcardactionMap["payload"].(string); ok {
		o.Payload = &Payload
	}
	
	if Url, ok := ContentcardactionMap["url"].(string); ok {
		o.Url = &Url
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentcardaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
