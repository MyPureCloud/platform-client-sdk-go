package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulingsettingsresponse
type Buschedulingsettingsresponse struct { 
	// MessageSeverities - Schedule generation message severity configuration
	MessageSeverities *[]Schedulermessagetypeseverity `json:"messageSeverities,omitempty"`

}

func (o *Buschedulingsettingsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulingsettingsresponse
	
	return json.Marshal(&struct { 
		MessageSeverities *[]Schedulermessagetypeseverity `json:"messageSeverities,omitempty"`
		*Alias
	}{ 
		MessageSeverities: o.MessageSeverities,
		Alias:    (*Alias)(o),
	})
}

func (o *Buschedulingsettingsresponse) UnmarshalJSON(b []byte) error {
	var BuschedulingsettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulingsettingsresponseMap)
	if err != nil {
		return err
	}
	
	if MessageSeverities, ok := BuschedulingsettingsresponseMap["messageSeverities"].([]interface{}); ok {
		MessageSeveritiesString, _ := json.Marshal(MessageSeverities)
		json.Unmarshal(MessageSeveritiesString, &o.MessageSeverities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulingsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
