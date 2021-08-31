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

func (o *Actionconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionconfig
	
	return json.Marshal(&struct { 
		Request *Requestconfig `json:"request,omitempty"`
		
		Response *Responseconfig `json:"response,omitempty"`
		*Alias
	}{ 
		Request: o.Request,
		
		Response: o.Response,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionconfig) UnmarshalJSON(b []byte) error {
	var ActionconfigMap map[string]interface{}
	err := json.Unmarshal(b, &ActionconfigMap)
	if err != nil {
		return err
	}
	
	if Request, ok := ActionconfigMap["request"].(map[string]interface{}); ok {
		RequestString, _ := json.Marshal(Request)
		json.Unmarshal(RequestString, &o.Request)
	}
	
	if Response, ok := ActionconfigMap["response"].(map[string]interface{}); ok {
		ResponseString, _ := json.Marshal(Response)
		json.Unmarshal(ResponseString, &o.Response)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
