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

func (u *Gkndocumentationresult) MarshalJSON() ([]byte, error) {
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
		Content: u.Content,
		
		Link: u.Link,
		
		Title: u.Title,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Gkndocumentationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
