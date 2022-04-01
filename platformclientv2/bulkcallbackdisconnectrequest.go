package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkcallbackdisconnectrequest
type Bulkcallbackdisconnectrequest struct { 
	// CallbackDisconnectIdentifiers - The list of requests to disconnect callbacks in bulk
	CallbackDisconnectIdentifiers *[]Callbackdisconnectidentifier `json:"callbackDisconnectIdentifiers,omitempty"`

}

func (o *Bulkcallbackdisconnectrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkcallbackdisconnectrequest
	
	return json.Marshal(&struct { 
		CallbackDisconnectIdentifiers *[]Callbackdisconnectidentifier `json:"callbackDisconnectIdentifiers,omitempty"`
		*Alias
	}{ 
		CallbackDisconnectIdentifiers: o.CallbackDisconnectIdentifiers,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkcallbackdisconnectrequest) UnmarshalJSON(b []byte) error {
	var BulkcallbackdisconnectrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkcallbackdisconnectrequestMap)
	if err != nil {
		return err
	}
	
	if CallbackDisconnectIdentifiers, ok := BulkcallbackdisconnectrequestMap["callbackDisconnectIdentifiers"].([]interface{}); ok {
		CallbackDisconnectIdentifiersString, _ := json.Marshal(CallbackDisconnectIdentifiers)
		json.Unmarshal(CallbackDisconnectIdentifiersString, &o.CallbackDisconnectIdentifiers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkcallbackdisconnectrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
