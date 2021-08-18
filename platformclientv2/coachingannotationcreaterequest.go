package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingannotationcreaterequest
type Coachingannotationcreaterequest struct { 
	// Text - The text of the annotation.
	Text *string `json:"text,omitempty"`


	// AccessType - Determines the permissions required to view this item.
	AccessType *string `json:"accessType,omitempty"`

}

func (u *Coachingannotationcreaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingannotationcreaterequest

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		AccessType *string `json:"accessType,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		AccessType: u.AccessType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coachingannotationcreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
