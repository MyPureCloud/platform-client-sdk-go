package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Textbotpromptsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
