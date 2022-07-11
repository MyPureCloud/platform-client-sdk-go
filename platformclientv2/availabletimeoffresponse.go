package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletimeoffresponse
type Availabletimeoffresponse struct { 
	// Values
	Values *[]Availabletimeoffrange `json:"values,omitempty"`

}

func (o *Availabletimeoffresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletimeoffresponse
	
	return json.Marshal(&struct { 
		Values *[]Availabletimeoffrange `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletimeoffresponse) UnmarshalJSON(b []byte) error {
	var AvailabletimeoffresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletimeoffresponseMap)
	if err != nil {
		return err
	}
	
	if Values, ok := AvailabletimeoffresponseMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletimeoffresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
