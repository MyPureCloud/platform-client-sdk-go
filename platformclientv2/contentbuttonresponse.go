package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentbuttonresponse - Button response object representing the click of a structured message button, such as a quick reply.
type Contentbuttonresponse struct { 
	// Id - An ID assigned to the button response (Deprecated).
	Id *string `json:"id,omitempty"`


	// VarType - Describes the button that resulted in the Button Response.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - The response payload associated with the clicked button.
	Payload *string `json:"payload,omitempty"`

}

func (o *Contentbuttonresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentbuttonresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentbuttonresponse) UnmarshalJSON(b []byte) error {
	var ContentbuttonresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ContentbuttonresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentbuttonresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := ContentbuttonresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := ContentbuttonresponseMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Payload, ok := ContentbuttonresponseMap["payload"].(string); ok {
		o.Payload = &Payload
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentbuttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
