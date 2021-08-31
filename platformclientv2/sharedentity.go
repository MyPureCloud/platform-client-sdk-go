package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sharedentity
type Sharedentity struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Sharedentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sharedentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Sharedentity) UnmarshalJSON(b []byte) error {
	var SharedentityMap map[string]interface{}
	err := json.Unmarshal(b, &SharedentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SharedentityMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sharedentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
