package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangeemailcolumn
type Dialercontactlistconfigchangeemailcolumn struct { 
	// ColumnName - The name of the email address column
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - The type of the email address column, for example, 'work' or 'home'
	VarType *string `json:"type,omitempty"`

}

func (o *Dialercontactlistconfigchangeemailcolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangeemailcolumn
	
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

func (o *Dialercontactlistconfigchangeemailcolumn) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangeemailcolumnMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangeemailcolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := DialercontactlistconfigchangeemailcolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if VarType, ok := DialercontactlistconfigchangeemailcolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangeemailcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
