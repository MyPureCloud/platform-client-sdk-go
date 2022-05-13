package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usageitem
type Usageitem struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// TotalDocumentByteCount
	TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`


	// TotalDocumentCount
	TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`

}

func (o *Usageitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usageitem
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`
		
		TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		TotalDocumentByteCount: o.TotalDocumentByteCount,
		
		TotalDocumentCount: o.TotalDocumentCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Usageitem) UnmarshalJSON(b []byte) error {
	var UsageitemMap map[string]interface{}
	err := json.Unmarshal(b, &UsageitemMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := UsageitemMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if TotalDocumentByteCount, ok := UsageitemMap["totalDocumentByteCount"].(float64); ok {
		TotalDocumentByteCountInt := int(TotalDocumentByteCount)
		o.TotalDocumentByteCount = &TotalDocumentByteCountInt
	}
	
	if TotalDocumentCount, ok := UsageitemMap["totalDocumentCount"].(float64); ok {
		TotalDocumentCountInt := int(TotalDocumentCount)
		o.TotalDocumentCount = &TotalDocumentCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Usageitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
