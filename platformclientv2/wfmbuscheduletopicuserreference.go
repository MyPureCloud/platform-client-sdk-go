package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicuserreference
type Wfmbuscheduletopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmbuscheduletopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicuserreference) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduletopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
