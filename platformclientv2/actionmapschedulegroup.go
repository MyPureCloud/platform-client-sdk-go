package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapschedulegroup
type Actionmapschedulegroup struct { 
	// Id - The ID of the action maps's associated schedule group.
	Id *string `json:"id,omitempty"`

}

func (o *Actionmapschedulegroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionmapschedulegroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionmapschedulegroup) UnmarshalJSON(b []byte) error {
	var ActionmapschedulegroupMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapschedulegroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActionmapschedulegroupMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapschedulegroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
