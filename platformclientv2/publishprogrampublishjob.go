package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishprogrampublishjob
type Publishprogrampublishjob struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

func (o *Publishprogrampublishjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishprogrampublishjob
	
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

func (o *Publishprogrampublishjob) UnmarshalJSON(b []byte) error {
	var PublishprogrampublishjobMap map[string]interface{}
	err := json.Unmarshal(b, &PublishprogrampublishjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PublishprogrampublishjobMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := PublishprogrampublishjobMap["state"].(string); ok {
		o.State = &State
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Publishprogrampublishjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
