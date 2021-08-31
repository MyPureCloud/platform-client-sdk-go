package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueobservationdatacontainer
type Queueobservationdatacontainer struct { 
	// Group - A mapping from dimension to value
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Observationmetricdata `json:"data,omitempty"`

}

func (o *Queueobservationdatacontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueobservationdatacontainer
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Observationmetricdata `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueobservationdatacontainer) UnmarshalJSON(b []byte) error {
	var QueueobservationdatacontainerMap map[string]interface{}
	err := json.Unmarshal(b, &QueueobservationdatacontainerMap)
	if err != nil {
		return err
	}
	
	if Group, ok := QueueobservationdatacontainerMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := QueueobservationdatacontainerMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueobservationdatacontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
