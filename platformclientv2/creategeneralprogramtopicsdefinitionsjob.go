package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Creategeneralprogramtopicsdefinitionsjob
type Creategeneralprogramtopicsdefinitionsjob struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

func (o *Creategeneralprogramtopicsdefinitionsjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Creategeneralprogramtopicsdefinitionsjob
	
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

func (o *Creategeneralprogramtopicsdefinitionsjob) UnmarshalJSON(b []byte) error {
	var CreategeneralprogramtopicsdefinitionsjobMap map[string]interface{}
	err := json.Unmarshal(b, &CreategeneralprogramtopicsdefinitionsjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreategeneralprogramtopicsdefinitionsjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := CreategeneralprogramtopicsdefinitionsjobMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Creategeneralprogramtopicsdefinitionsjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
