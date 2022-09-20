package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulingsettingsrequest
type Buschedulingsettingsrequest struct { 
	// MessageSeverities - Schedule generation message severity configuration
	MessageSeverities *[]Schedulermessagetypeseverity `json:"messageSeverities,omitempty"`

}

func (o *Buschedulingsettingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulingsettingsrequest
	
	return json.Marshal(&struct { 
		MessageSeverities *[]Schedulermessagetypeseverity `json:"messageSeverities,omitempty"`
		*Alias
	}{ 
		MessageSeverities: o.MessageSeverities,
		Alias:    (*Alias)(o),
	})
}

func (o *Buschedulingsettingsrequest) UnmarshalJSON(b []byte) error {
	var BuschedulingsettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulingsettingsrequestMap)
	if err != nil {
		return err
	}
	
	if MessageSeverities, ok := BuschedulingsettingsrequestMap["messageSeverities"].([]interface{}); ok {
		MessageSeveritiesString, _ := json.Marshal(MessageSeverities)
		json.Unmarshal(MessageSeveritiesString, &o.MessageSeverities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulingsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
