package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchexternalsegment
type Patchexternalsegment struct { 
	// Name - Name for the external segment in the system where it originates from.
	Name *string `json:"name,omitempty"`

}

func (o *Patchexternalsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchexternalsegment
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchexternalsegment) UnmarshalJSON(b []byte) error {
	var PatchexternalsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &PatchexternalsegmentMap)
	if err != nil {
		return err
	}
	
	if Name, ok := PatchexternalsegmentMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchexternalsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
