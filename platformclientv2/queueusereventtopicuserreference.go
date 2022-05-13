package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueusereventtopicuserreference
type Queueusereventtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Queueusereventtopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueusereventtopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueusereventtopicuserreference) UnmarshalJSON(b []byte) error {
	var QueueusereventtopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &QueueusereventtopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueusereventtopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueusereventtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
