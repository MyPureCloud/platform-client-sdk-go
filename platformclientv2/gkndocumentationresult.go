package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gkndocumentationresult
type Gkndocumentationresult struct { 
	// Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
	Content *string `json:"content,omitempty"`


	// Link - URL link for the documentation entity. Will be returned in responses for certain entities.
	Link *string `json:"link,omitempty"`


	// Title - The title of the documentation entity. Will be returned in responses for certain entities.
	Title *string `json:"title,omitempty"`


	// VarType - The search type. Will be returned in responses for certain entities.
	VarType *string `json:"_type,omitempty"`

}

func (o *Gkndocumentationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gkndocumentationresult
	
	return json.Marshal(&struct { 
		Content *string `json:"content,omitempty"`
		
		Link *string `json:"link,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		VarType *string `json:"_type,omitempty"`
		*Alias
	}{ 
		Content: o.Content,
		
		Link: o.Link,
		
		Title: o.Title,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Gkndocumentationresult) UnmarshalJSON(b []byte) error {
	var GkndocumentationresultMap map[string]interface{}
	err := json.Unmarshal(b, &GkndocumentationresultMap)
	if err != nil {
		return err
	}
	
	if Content, ok := GkndocumentationresultMap["content"].(string); ok {
		o.Content = &Content
	}
    
	if Link, ok := GkndocumentationresultMap["link"].(string); ok {
		o.Link = &Link
	}
    
	if Title, ok := GkndocumentationresultMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if VarType, ok := GkndocumentationresultMap["_type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Gkndocumentationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
