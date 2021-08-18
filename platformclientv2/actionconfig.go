package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionconfig - Defines components of the Action Config.
type Actionconfig struct { 
	// Request - Configuration of outbound request.
	Request *Requestconfig `json:"request,omitempty"`


	// Response - Configuration of response processing.
	Response *Responseconfig `json:"response,omitempty"`

}

func (u *Actionconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionconfig

	

	return json.Marshal(&struct { 
		Request *Requestconfig `json:"request,omitempty"`
		
		Response *Responseconfig `json:"response,omitempty"`
		*Alias
	}{ 
		Request: u.Request,
		
		Response: u.Response,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actionconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
