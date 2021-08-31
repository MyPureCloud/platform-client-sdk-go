package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Useraggregatedatacontainer
type Useraggregatedatacontainer struct { 
	// Group - A mapping from dimension to value
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Statisticalresponse `json:"data,omitempty"`

}

func (o *Useraggregatedatacontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useraggregatedatacontainer
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Statisticalresponse `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Useraggregatedatacontainer) UnmarshalJSON(b []byte) error {
	var UseraggregatedatacontainerMap map[string]interface{}
	err := json.Unmarshal(b, &UseraggregatedatacontainerMap)
	if err != nil {
		return err
	}
	
	if Group, ok := UseraggregatedatacontainerMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := UseraggregatedatacontainerMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Useraggregatedatacontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
