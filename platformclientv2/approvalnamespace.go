package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Approvalnamespace
type Approvalnamespace struct { 
	// Namespace - The namespace of the associated approvers.
	Namespace *string `json:"namespace,omitempty"`


	// Status - The current namespace approval status.
	Status *string `json:"status,omitempty"`


	// VarType - The type of namespace approval.
	VarType *string `json:"type,omitempty"`

}

func (o *Approvalnamespace) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Approvalnamespace
	
	return json.Marshal(&struct { 
		Namespace *string `json:"namespace,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Namespace: o.Namespace,
		
		Status: o.Status,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Approvalnamespace) UnmarshalJSON(b []byte) error {
	var ApprovalnamespaceMap map[string]interface{}
	err := json.Unmarshal(b, &ApprovalnamespaceMap)
	if err != nil {
		return err
	}
	
	if Namespace, ok := ApprovalnamespaceMap["namespace"].(string); ok {
		o.Namespace = &Namespace
	}
    
	if Status, ok := ApprovalnamespaceMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarType, ok := ApprovalnamespaceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Approvalnamespace) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
