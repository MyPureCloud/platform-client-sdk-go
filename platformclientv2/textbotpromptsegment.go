package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotpromptsegment - Data for a single bot flow prompt segment.
type Textbotpromptsegment struct { 
	// Text - The text of this prompt segment.
	Text *string `json:"text,omitempty"`


	// VarType - The segment type which describes any semantics about the 'text' and also indicates which other field might include additional relevant info.
	VarType *string `json:"type,omitempty"`


	// Format - Additional details describing the segment’s contents, which the client should honour where possible.
	Format *Format `json:"format,omitempty"`

}

func (u *Textbotpromptsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotpromptsegment

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Format *Format `json:"format,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		VarType: u.VarType,
		
		Format: u.Format,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotpromptsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
