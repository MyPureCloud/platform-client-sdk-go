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

func (u *Helplink) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Helplink

	

	return json.Marshal(&struct { 
		Uri *string `json:"uri,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Uri: u.Uri,
		
		Title: u.Title,
		
		Description: u.Description,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Helplink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
