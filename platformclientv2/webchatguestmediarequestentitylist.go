package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatguestmediarequestentitylist
type Webchatguestmediarequestentitylist struct { 
	// Entities
	Entities *[]Webchatguestmediarequest `json:"entities,omitempty"`

}

func (o *Webchatguestmediarequestentitylist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatguestmediarequestentitylist
	
	return json.Marshal(&struct { 
		Entities *[]Webchatguestmediarequest `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatguestmediarequestentitylist) UnmarshalJSON(b []byte) error {
	var WebchatguestmediarequestentitylistMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatguestmediarequestentitylistMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WebchatguestmediarequestentitylistMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequestentitylist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
