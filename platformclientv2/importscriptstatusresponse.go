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

func (o *Importscriptstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importscriptstatusresponse
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Succeeded *bool `json:"succeeded,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		Succeeded: o.Succeeded,
		
		Message: o.Message,
		Alias:    (*Alias)(o),
	})
}

func (o *Importscriptstatusresponse) UnmarshalJSON(b []byte) error {
	var ImportscriptstatusresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ImportscriptstatusresponseMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ImportscriptstatusresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Succeeded, ok := ImportscriptstatusresponseMap["succeeded"].(bool); ok {
		o.Succeeded = &Succeeded
	}
    
	if Message, ok := ImportscriptstatusresponseMap["message"].(string); ok {
		o.Message = &Message
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Importscriptstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
