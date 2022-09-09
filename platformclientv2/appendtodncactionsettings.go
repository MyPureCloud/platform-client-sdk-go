package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Appendtodncactionsettings
type Appendtodncactionsettings struct { 
	// Expire - Whether to expire the record appended to the DNC list.
	Expire *bool `json:"expire,omitempty"`


	// ExpirationDuration - If 'expire' is set to true, how long to keep the record.
	ExpirationDuration *string `json:"expirationDuration,omitempty"`

}

func (o *Appendtodncactionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Appendtodncactionsettings
	
	return json.Marshal(&struct { 
		Expire *bool `json:"expire,omitempty"`
		
		ExpirationDuration *string `json:"expirationDuration,omitempty"`
		*Alias
	}{ 
		Expire: o.Expire,
		
		ExpirationDuration: o.ExpirationDuration,
		Alias:    (*Alias)(o),
	})
}

func (o *Appendtodncactionsettings) UnmarshalJSON(b []byte) error {
	var AppendtodncactionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AppendtodncactionsettingsMap)
	if err != nil {
		return err
	}
	
	if Expire, ok := AppendtodncactionsettingsMap["expire"].(bool); ok {
		o.Expire = &Expire
	}
    
	if ExpirationDuration, ok := AppendtodncactionsettingsMap["expirationDuration"].(string); ok {
		o.ExpirationDuration = &ExpirationDuration
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Appendtodncactionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
