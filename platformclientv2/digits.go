package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Digits
type Digits struct { 
	// Digits - A string representing the digits pressed on phone.
	Digits *string `json:"digits,omitempty"`

}

func (o *Digits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Digits
	
	return json.Marshal(&struct { 
		Digits *string `json:"digits,omitempty"`
		*Alias
	}{ 
		Digits: o.Digits,
		Alias:    (*Alias)(o),
	})
}

func (o *Digits) UnmarshalJSON(b []byte) error {
	var DigitsMap map[string]interface{}
	err := json.Unmarshal(b, &DigitsMap)
	if err != nil {
		return err
	}
	
	if Digits, ok := DigitsMap["digits"].(string); ok {
		o.Digits = &Digits
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Digits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
