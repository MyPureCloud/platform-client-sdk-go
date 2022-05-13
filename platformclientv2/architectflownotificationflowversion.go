package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationflowversion - A bare-bones flow version object
type Architectflownotificationflowversion struct { 
	// Id - The version ID
	Id *string `json:"id,omitempty"`

}

func (o *Architectflownotificationflowversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationflowversion
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflownotificationflowversion) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationflowversionMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationflowversionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflownotificationflowversionMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationflowversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
