package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Detectednamedentityvalue
type Detectednamedentityvalue struct { 
	// Raw - The raw value of the detected named entity.
	Raw *string `json:"raw,omitempty"`


	// Resolved - The resolved value of the detected named entity.
	Resolved *string `json:"resolved,omitempty"`

}

func (u *Detectednamedentityvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Detectednamedentityvalue

	

	return json.Marshal(&struct { 
		Raw *string `json:"raw,omitempty"`
		
		Resolved *string `json:"resolved,omitempty"`
		*Alias
	}{ 
		Raw: u.Raw,
		
		Resolved: u.Resolved,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Detectednamedentityvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
