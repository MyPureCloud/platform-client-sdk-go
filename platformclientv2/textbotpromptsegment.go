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


	// Content - Details to display Rich Media content. This is only populated when the segment 'type' is 'Rich Media'.
	Content *[]Messagecontent `json:"content,omitempty"`

}

func (o *Textbotpromptsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotpromptsegment
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Format *Format `json:"format,omitempty"`
		
		Content *[]Messagecontent `json:"content,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		VarType: o.VarType,
		
		Format: o.Format,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotpromptsegment) UnmarshalJSON(b []byte) error {
	var TextbotpromptsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotpromptsegmentMap)
	if err != nil {
		return err
	}
	
	if Text, ok := TextbotpromptsegmentMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if VarType, ok := TextbotpromptsegmentMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Format, ok := TextbotpromptsegmentMap["format"].(map[string]interface{}); ok {
		FormatString, _ := json.Marshal(Format)
		json.Unmarshal(FormatString, &o.Format)
	}
	
	if Content, ok := TextbotpromptsegmentMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotpromptsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
