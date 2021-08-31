package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Movemanagementunitrequest
type Movemanagementunitrequest struct { 
	// BusinessUnitId - The ID of the business unit to which to move the management unit
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

}

func (o *Movemanagementunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Movemanagementunitrequest
	
	return json.Marshal(&struct { 
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		*Alias
	}{ 
		BusinessUnitId: o.BusinessUnitId,
		Alias:    (*Alias)(o),
	})
}

func (o *Movemanagementunitrequest) UnmarshalJSON(b []byte) error {
	var MovemanagementunitrequestMap map[string]interface{}
	err := json.Unmarshal(b, &MovemanagementunitrequestMap)
	if err != nil {
		return err
	}
	
	if BusinessUnitId, ok := MovemanagementunitrequestMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Movemanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
