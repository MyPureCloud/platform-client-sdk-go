package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsetext - Contains information about the text associated with a response.
type Responsetext struct { 
	// Content - Response text content.
	Content *string `json:"content,omitempty"`


	// ContentType - Response text content type.
	ContentType *string `json:"contentType,omitempty"`

}

func (o *Responsetext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsetext
	
	return json.Marshal(&struct { 
		Content *string `json:"content,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		*Alias
	}{ 
		Content: o.Content,
		
		ContentType: o.ContentType,
		Alias:    (*Alias)(o),
	})
}

func (o *Responsetext) UnmarshalJSON(b []byte) error {
	var ResponsetextMap map[string]interface{}
	err := json.Unmarshal(b, &ResponsetextMap)
	if err != nil {
		return err
	}
	
	if Content, ok := ResponsetextMap["content"].(string); ok {
		o.Content = &Content
	}
	
	if ContentType, ok := ResponsetextMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responsetext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
