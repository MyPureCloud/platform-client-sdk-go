package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperdate - An object to provide context to nullable fields in PATCH requests
type Valuewrapperdate struct { 
	// Value - The value for the associated field. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Value *time.Time `json:"value,omitempty"`

}

func (o *Valuewrapperdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Valuewrapperdate
	
	Value := new(string)
	if o.Value != nil {
		
		*Value = timeutil.Strftime(o.Value, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Value = nil
	}
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Value: Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Valuewrapperdate) UnmarshalJSON(b []byte) error {
	var ValuewrapperdateMap map[string]interface{}
	err := json.Unmarshal(b, &ValuewrapperdateMap)
	if err != nil {
		return err
	}
	
	if valueString, ok := ValuewrapperdateMap["value"].(string); ok {
		Value, _ := time.Parse("2006-01-02T15:04:05.999999Z", valueString)
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Valuewrapperdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
