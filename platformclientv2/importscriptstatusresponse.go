package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Importscriptstatusresponse
type Importscriptstatusresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Succeeded
	Succeeded *bool `json:"succeeded,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`

}

func (u *Importscriptstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importscriptstatusresponse

	

	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Succeeded *bool `json:"succeeded,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		Url: u.Url,
		
		Succeeded: u.Succeeded,
		
		Message: u.Message,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Importscriptstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
