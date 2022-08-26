package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentcontentblock
type Documentcontentblock struct { 
	// VarType - The type of the paragraph block.
	VarType *string `json:"type,omitempty"`


	// Text - Text. It must contain a value if the type of the block is Text.
	Text *Documenttext `json:"text,omitempty"`


	// Image - Image. It must contain a value if the type of the block is Image.
	Image *Documentbodyimage `json:"image,omitempty"`

}

func (o *Documentcontentblock) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentcontentblock
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *Documenttext `json:"text,omitempty"`
		
		Image *Documentbodyimage `json:"image,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Image: o.Image,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentcontentblock) UnmarshalJSON(b []byte) error {
	var DocumentcontentblockMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentcontentblockMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DocumentcontentblockMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := DocumentcontentblockMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	
	if Image, ok := DocumentcontentblockMap["image"].(map[string]interface{}); ok {
		ImageString, _ := json.Marshal(Image)
		json.Unmarshal(ImageString, &o.Image)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentcontentblock) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
