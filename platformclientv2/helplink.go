package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Helplink - Link to a help or support resource
type Helplink struct { 
	// Uri - URI of the help resource
	Uri *string `json:"uri,omitempty"`


	// Title - Link text of the resource
	Title *string `json:"title,omitempty"`


	// Description - Description of the document or resource
	Description *string `json:"description,omitempty"`

}

func (o *Helplink) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Helplink
	
	return json.Marshal(&struct { 
		Uri *string `json:"uri,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Uri: o.Uri,
		
		Title: o.Title,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Helplink) UnmarshalJSON(b []byte) error {
	var HelplinkMap map[string]interface{}
	err := json.Unmarshal(b, &HelplinkMap)
	if err != nil {
		return err
	}
	
	if Uri, ok := HelplinkMap["uri"].(string); ok {
		o.Uri = &Uri
	}
	
	if Title, ok := HelplinkMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := HelplinkMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Helplink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
