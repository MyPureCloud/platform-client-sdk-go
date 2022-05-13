package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wrapupcodereference
type Wrapupcodereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

}

func (o *Wrapupcodereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wrapupcodereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wrapupcodereference) UnmarshalJSON(b []byte) error {
	var WrapupcodereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WrapupcodereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WrapupcodereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wrapupcodereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
