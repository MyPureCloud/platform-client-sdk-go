package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Batcheventresponse
type Batcheventresponse struct { 
	// Errors - A list of validation or server errors that occurred for posted events.
	Errors *[]Eventerror `json:"errors,omitempty"`

}

func (o *Batcheventresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batcheventresponse
	
	return json.Marshal(&struct { 
		Errors *[]Eventerror `json:"errors,omitempty"`
		*Alias
	}{ 
		Errors: o.Errors,
		Alias:    (*Alias)(o),
	})
}

func (o *Batcheventresponse) UnmarshalJSON(b []byte) error {
	var BatcheventresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BatcheventresponseMap)
	if err != nil {
		return err
	}
	
	if Errors, ok := BatcheventresponseMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batcheventresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
