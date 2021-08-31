package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentpostback - Postback response object representing the click of a rich media button (Deprecated).
type Contentpostback struct { 
	// Id - An ID assigned to the button response.
	Id *string `json:"id,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - The response payload associated with the clicked button.
	Payload *string `json:"payload,omitempty"`

}

func (o *Contentpostback) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentpostback
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		Payload: o.Payload,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentpostback) UnmarshalJSON(b []byte) error {
	var ContentpostbackMap map[string]interface{}
	err := json.Unmarshal(b, &ContentpostbackMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentpostbackMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Text, ok := ContentpostbackMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Payload, ok := ContentpostbackMap["payload"].(string); ok {
		o.Payload = &Payload
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentpostback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
