package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluutterancesegment
type Nluutterancesegment struct { 
	// Text - The text of the segment.
	Text *string `json:"text,omitempty"`


	// Entity - The entity annotation of the segment.
	Entity *Namedentityannotation `json:"entity,omitempty"`

}

func (u *Nluutterancesegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluutterancesegment

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Entity *Namedentityannotation `json:"entity,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Entity: u.Entity,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluutterancesegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
