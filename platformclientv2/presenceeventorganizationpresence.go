package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presenceeventorganizationpresence
type Presenceeventorganizationpresence struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

}

func (o *Presenceeventorganizationpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presenceeventorganizationpresence
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SystemPresence: o.SystemPresence,
		Alias:    (*Alias)(o),
	})
}

func (o *Presenceeventorganizationpresence) UnmarshalJSON(b []byte) error {
	var PresenceeventorganizationpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &PresenceeventorganizationpresenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PresenceeventorganizationpresenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SystemPresence, ok := PresenceeventorganizationpresenceMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Presenceeventorganizationpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
