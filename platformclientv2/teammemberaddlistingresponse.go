package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Teammemberaddlistingresponse
type Teammemberaddlistingresponse struct { 
	// Entities
	Entities *[]Userreference `json:"entities,omitempty"`


	// Failures - List of any user ids that were not added.
	Failures *[]Teamaddmemberfailure `json:"failures,omitempty"`

}

func (o *Teammemberaddlistingresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Teammemberaddlistingresponse
	
	return json.Marshal(&struct { 
		Entities *[]Userreference `json:"entities,omitempty"`
		
		Failures *[]Teamaddmemberfailure `json:"failures,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		Failures: o.Failures,
		Alias:    (*Alias)(o),
	})
}

func (o *Teammemberaddlistingresponse) UnmarshalJSON(b []byte) error {
	var TeammemberaddlistingresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TeammemberaddlistingresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := TeammemberaddlistingresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Failures, ok := TeammemberaddlistingresponseMap["failures"].([]interface{}); ok {
		FailuresString, _ := json.Marshal(Failures)
		json.Unmarshal(FailuresString, &o.Failures)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Teammemberaddlistingresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
