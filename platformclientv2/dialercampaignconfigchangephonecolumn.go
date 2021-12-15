package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangephonecolumn
type Dialercampaignconfigchangephonecolumn struct { 
	// ColumnName - The name of the phone column
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - The type of the phone column, for example, 'cell' or 'home'
	VarType *string `json:"type,omitempty"`

}

func (o *Dialercampaignconfigchangephonecolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangephonecolumn
	
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

func (o *Dialercampaignconfigchangephonecolumn) UnmarshalJSON(b []byte) error {
	var DialercampaignconfigchangephonecolumnMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignconfigchangephonecolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := DialercampaignconfigchangephonecolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
	
	if VarType, ok := DialercampaignconfigchangephonecolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangephonecolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
