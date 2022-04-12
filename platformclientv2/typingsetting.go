package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Typingsetting
type Typingsetting struct { 
	// On - Should typing indication Events be sent
	On *Settingdirection `json:"on,omitempty"`

}

func (o *Typingsetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Typingsetting
	
	return json.Marshal(&struct { 
		On *Settingdirection `json:"on,omitempty"`
		*Alias
	}{ 
		On: o.On,
		Alias:    (*Alias)(o),
	})
}

func (o *Typingsetting) UnmarshalJSON(b []byte) error {
	var TypingsettingMap map[string]interface{}
	err := json.Unmarshal(b, &TypingsettingMap)
	if err != nil {
		return err
	}
	
	if On, ok := TypingsettingMap["on"].(map[string]interface{}); ok {
		OnString, _ := json.Marshal(On)
		json.Unmarshal(OnString, &o.On)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Typingsetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
