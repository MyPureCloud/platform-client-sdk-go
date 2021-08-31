package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishtopicpublishjob
type Publishtopicpublishjob struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

func (o *Publishtopicpublishjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishtopicpublishjob
	
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

func (o *Publishtopicpublishjob) UnmarshalJSON(b []byte) error {
	var PublishtopicpublishjobMap map[string]interface{}
	err := json.Unmarshal(b, &PublishtopicpublishjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PublishtopicpublishjobMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := PublishtopicpublishjobMap["state"].(string); ok {
		o.State = &State
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Publishtopicpublishjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
