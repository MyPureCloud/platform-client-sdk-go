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

func (u *Responsetext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsetext

	

	return json.Marshal(&struct { 
		Content *string `json:"content,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		*Alias
	}{ 
		Content: u.Content,
		
		ContentType: u.ContentType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Responsetext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
