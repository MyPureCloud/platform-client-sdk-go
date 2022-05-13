package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationeventaction
type Journeywebactioneventsnotificationeventaction struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Journeywebactioneventsnotificationeventaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationeventaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationeventaction) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationeventactionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationeventactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebactioneventsnotificationeventactionMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationeventaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
