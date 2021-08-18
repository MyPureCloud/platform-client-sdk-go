package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomainversiontrainingresponse
type Nludomainversiontrainingresponse struct { 
	// Message - A message indicating result of the action.
	Message *string `json:"message,omitempty"`


	// Version
	Version *Nludomainversion `json:"version,omitempty"`

}

func (u *Nludomainversiontrainingresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludomainversiontrainingresponse

	

	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Version *Nludomainversion `json:"version,omitempty"`
		*Alias
	}{ 
		Message: u.Message,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludomainversiontrainingresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
