package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2patchoperation - Defines a SCIM PATCH operation. The path and value follow very specific rules based on operation types. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchoperation struct { 
	// Op - The PATCH operation to perform.
	Op *string `json:"op,omitempty"`


	// Path - The attribute path that describes the target of the operation. Required for a \"remove\" operation.
	Path *string `json:"path,omitempty"`


	// Value - The value to set in the path.
	Value *interface{} `json:"value,omitempty"`

}

func (o *Scimv2patchoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2patchoperation
	
	return json.Marshal(&struct { 
		Op *string `json:"op,omitempty"`
		
		Path *string `json:"path,omitempty"`
		
		Value *interface{} `json:"value,omitempty"`
		*Alias
	}{ 
		Op: o.Op,
		
		Path: o.Path,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimv2patchoperation) UnmarshalJSON(b []byte) error {
	var Scimv2patchoperationMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2patchoperationMap)
	if err != nil {
		return err
	}
	
	if Op, ok := Scimv2patchoperationMap["op"].(string); ok {
		o.Op = &Op
	}
    
	if Path, ok := Scimv2patchoperationMap["path"].(string); ok {
		o.Path = &Path
	}
    
	if Value, ok := Scimv2patchoperationMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2patchoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
