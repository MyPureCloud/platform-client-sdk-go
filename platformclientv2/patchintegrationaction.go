package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchintegrationaction
type Patchintegrationaction struct { 
	// Id - ID of the integration action to be invoked.
	Id *string `json:"id,omitempty"`

}

func (o *Patchintegrationaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchintegrationaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchintegrationaction) UnmarshalJSON(b []byte) error {
	var PatchintegrationactionMap map[string]interface{}
	err := json.Unmarshal(b, &PatchintegrationactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PatchintegrationactionMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchintegrationaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
