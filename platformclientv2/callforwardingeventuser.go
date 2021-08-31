package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwardingeventuser
type Callforwardingeventuser struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Callforwardingeventuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwardingeventuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Callforwardingeventuser) UnmarshalJSON(b []byte) error {
	var CallforwardingeventuserMap map[string]interface{}
	err := json.Unmarshal(b, &CallforwardingeventuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallforwardingeventuserMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callforwardingeventuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
