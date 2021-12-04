package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Failedrecordingentitylisting
type Failedrecordingentitylisting struct { 
	// Entities
	Entities *[]Recordingjobfailedrecording `json:"entities,omitempty"`

}

func (o *Failedrecordingentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Failedrecordingentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Recordingjobfailedrecording `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Failedrecordingentitylisting) UnmarshalJSON(b []byte) error {
	var FailedrecordingentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &FailedrecordingentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := FailedrecordingentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Failedrecordingentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
