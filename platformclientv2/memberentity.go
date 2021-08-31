package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Memberentity
type Memberentity struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Memberentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Memberentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Memberentity) UnmarshalJSON(b []byte) error {
	var MemberentityMap map[string]interface{}
	err := json.Unmarshal(b, &MemberentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MemberentityMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Memberentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
