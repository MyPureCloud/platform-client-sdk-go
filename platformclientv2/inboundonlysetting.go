package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Inboundonlysetting
type Inboundonlysetting struct { 
	// Inbound
	Inbound *string `json:"inbound,omitempty"`

}

func (o *Inboundonlysetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Inboundonlysetting
	
	return json.Marshal(&struct { 
		Inbound *string `json:"inbound,omitempty"`
		*Alias
	}{ 
		Inbound: o.Inbound,
		Alias:    (*Alias)(o),
	})
}

func (o *Inboundonlysetting) UnmarshalJSON(b []byte) error {
	var InboundonlysettingMap map[string]interface{}
	err := json.Unmarshal(b, &InboundonlysettingMap)
	if err != nil {
		return err
	}
	
	if Inbound, ok := InboundonlysettingMap["inbound"].(string); ok {
		o.Inbound = &Inbound
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Inboundonlysetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
