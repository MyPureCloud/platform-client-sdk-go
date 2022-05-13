package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Importstatusrequest
type Importstatusrequest struct { 
	// Status - New status for existing import operation
	Status *string `json:"status,omitempty"`

}

func (o *Importstatusrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importstatusrequest
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Importstatusrequest) UnmarshalJSON(b []byte) error {
	var ImportstatusrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ImportstatusrequestMap)
	if err != nil {
		return err
	}
	
	if Status, ok := ImportstatusrequestMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Importstatusrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
