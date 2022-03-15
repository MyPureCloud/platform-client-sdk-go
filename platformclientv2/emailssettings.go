package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailssettings
type Emailssettings struct { 
	// SendingSizeLimit
	SendingSizeLimit *int `json:"sendingSizeLimit,omitempty"`

}

func (o *Emailssettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailssettings
	
	return json.Marshal(&struct { 
		SendingSizeLimit *int `json:"sendingSizeLimit,omitempty"`
		*Alias
	}{ 
		SendingSizeLimit: o.SendingSizeLimit,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailssettings) UnmarshalJSON(b []byte) error {
	var EmailssettingsMap map[string]interface{}
	err := json.Unmarshal(b, &EmailssettingsMap)
	if err != nil {
		return err
	}
	
	if SendingSizeLimit, ok := EmailssettingsMap["sendingSizeLimit"].(float64); ok {
		SendingSizeLimitInt := int(SendingSizeLimit)
		o.SendingSizeLimit = &SendingSizeLimitInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emailssettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
