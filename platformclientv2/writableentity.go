package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Writableentity
type Writableentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

}

func (o *Writableentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Writableentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Writableentity) UnmarshalJSON(b []byte) error {
	var WritableentityMap map[string]interface{}
	err := json.Unmarshal(b, &WritableentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WritableentityMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Writableentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
