package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeidnamepair
type Edgeidnamepair struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Edgeidnamepair) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeidnamepair
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeidnamepair) UnmarshalJSON(b []byte) error {
	var EdgeidnamepairMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeidnamepairMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgeidnamepairMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := EdgeidnamepairMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeidnamepair) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
