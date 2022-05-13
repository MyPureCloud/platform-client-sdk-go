package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicuserreference
type Wfmagentscheduleupdatetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmagentscheduleupdatetopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicuserreference) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmagentscheduleupdatetopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
