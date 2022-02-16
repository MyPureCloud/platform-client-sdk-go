package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishtopictopicsdefinitionsjob
type Publishtopictopicsdefinitionsjob struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

func (o *Publishtopictopicsdefinitionsjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishtopictopicsdefinitionsjob
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Publishtopictopicsdefinitionsjob) UnmarshalJSON(b []byte) error {
	var PublishtopictopicsdefinitionsjobMap map[string]interface{}
	err := json.Unmarshal(b, &PublishtopictopicsdefinitionsjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PublishtopictopicsdefinitionsjobMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := PublishtopictopicsdefinitionsjobMap["state"].(string); ok {
		o.State = &State
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Publishtopictopicsdefinitionsjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
