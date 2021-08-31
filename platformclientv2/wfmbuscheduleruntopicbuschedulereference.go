package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduleruntopicbuschedulereference
type Wfmbuscheduleruntopicbuschedulereference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmbuscheduleruntopicbuschedulereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduleruntopicbuschedulereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduleruntopicbuschedulereference) UnmarshalJSON(b []byte) error {
	var WfmbuscheduleruntopicbuschedulereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduleruntopicbuschedulereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduleruntopicbuschedulereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
