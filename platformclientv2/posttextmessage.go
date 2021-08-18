package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Posttextmessage
type Posttextmessage struct { 
	// VarType - Message type
	VarType *string `json:"type,omitempty"`


	// Text - Message text. If type is structured, used as fallback for clients that do not support particular structured content
	Text *string `json:"text,omitempty"`


	// Content - A list of content elements in message
	Content *[]Messagecontent `json:"content,omitempty"`

}

func (u *Posttextmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Posttextmessage

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Content *[]Messagecontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Text: u.Text,
		
		Content: u.Content,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Posttextmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
