package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fileuploadsettings - File upload settings for messenger
type Fileuploadsettings struct { 
	// Modes - The list of supported file upload modes
	Modes *[]Fileuploadmode `json:"modes,omitempty"`

}

func (u *Fileuploadsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fileuploadsettings

	

	return json.Marshal(&struct { 
		Modes *[]Fileuploadmode `json:"modes,omitempty"`
		*Alias
	}{ 
		Modes: u.Modes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Fileuploadsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
