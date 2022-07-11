package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sessionenddetails
type Sessionenddetails struct { 
	// VarType - The type of termination handling that resulted in the session end. It can be either Exit or Disconnect
	VarType *string `json:"type,omitempty"`


	// Reason - The reason for termination action. It can be due to an error or normal flow execution
	Reason *string `json:"reason,omitempty"`

}

func (o *Sessionenddetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sessionenddetails
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Reason *string `json:"reason,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Reason: o.Reason,
		Alias:    (*Alias)(o),
	})
}

func (o *Sessionenddetails) UnmarshalJSON(b []byte) error {
	var SessionenddetailsMap map[string]interface{}
	err := json.Unmarshal(b, &SessionenddetailsMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SessionenddetailsMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Reason, ok := SessionenddetailsMap["reason"].(string); ok {
		o.Reason = &Reason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sessionenddetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
