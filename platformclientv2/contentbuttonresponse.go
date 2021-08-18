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

func (u *Contentbuttonresponse) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		VarType: u.VarType,
		
		Text: u.Text,
		
		Payload: u.Payload,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentbuttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
