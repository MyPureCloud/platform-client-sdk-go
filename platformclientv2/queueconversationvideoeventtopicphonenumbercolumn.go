package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicphonenumbercolumn
type Queueconversationvideoeventtopicphonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Queueconversationvideoeventtopicphonenumbercolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicphonenumbercolumn
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicphonenumbercolumn) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicphonenumbercolumnMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicphonenumbercolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := QueueconversationvideoeventtopicphonenumbercolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
	
	if VarType, ok := QueueconversationvideoeventtopicphonenumbercolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
