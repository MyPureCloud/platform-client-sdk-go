package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingsettingrequestreference - Messaging Setting for messaging platform integrations
type Messagingsettingrequestreference struct { 
	// Id - The messaging Setting unique identifier associated with this integration
	Id *string `json:"id,omitempty"`

}

func (o *Messagingsettingrequestreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingsettingrequestreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagingsettingrequestreference) UnmarshalJSON(b []byte) error {
	var MessagingsettingrequestreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingsettingrequestreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingsettingrequestreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingsettingrequestreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
