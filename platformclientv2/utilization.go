package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Utilization
type Utilization struct { 
	// Utilization - Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
	Utilization *map[string]Mediautilization `json:"utilization,omitempty"`

}

func (o *Utilization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Utilization
	
	return json.Marshal(&struct { 
		Utilization *map[string]Mediautilization `json:"utilization,omitempty"`
		*Alias
	}{ 
		Utilization: o.Utilization,
		Alias:    (*Alias)(o),
	})
}

func (o *Utilization) UnmarshalJSON(b []byte) error {
	var UtilizationMap map[string]interface{}
	err := json.Unmarshal(b, &UtilizationMap)
	if err != nil {
		return err
	}
	
	if Utilization, ok := UtilizationMap["utilization"].(map[string]interface{}); ok {
		UtilizationString, _ := json.Marshal(Utilization)
		json.Unmarshal(UtilizationString, &o.Utilization)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Utilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
