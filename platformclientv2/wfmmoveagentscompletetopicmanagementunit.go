package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmoveagentscompletetopicmanagementunit
type Wfmmoveagentscompletetopicmanagementunit struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmmoveagentscompletetopicmanagementunit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmoveagentscompletetopicmanagementunit
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmmoveagentscompletetopicmanagementunit) UnmarshalJSON(b []byte) error {
	var WfmmoveagentscompletetopicmanagementunitMap map[string]interface{}
	err := json.Unmarshal(b, &WfmmoveagentscompletetopicmanagementunitMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmmoveagentscompletetopicmanagementunitMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicmanagementunit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
