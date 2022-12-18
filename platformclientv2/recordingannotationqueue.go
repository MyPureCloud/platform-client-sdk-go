package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingannotationqueue
type Recordingannotationqueue struct { 
	// Name - The queue name
	Name *string `json:"name,omitempty"`


	// Id - The queue Id
	Id *string `json:"id,omitempty"`

}

func (o *Recordingannotationqueue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingannotationqueue
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingannotationqueue) UnmarshalJSON(b []byte) error {
	var RecordingannotationqueueMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingannotationqueueMap)
	if err != nil {
		return err
	}
	
	if Name, ok := RecordingannotationqueueMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Id, ok := RecordingannotationqueueMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingannotationqueue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
