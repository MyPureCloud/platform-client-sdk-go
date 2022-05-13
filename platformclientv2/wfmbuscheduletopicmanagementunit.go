package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicmanagementunit
type Wfmbuscheduletopicmanagementunit struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmbuscheduletopicmanagementunit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicmanagementunit
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicmanagementunit) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicmanagementunitMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicmanagementunitMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduletopicmanagementunitMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicmanagementunit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
