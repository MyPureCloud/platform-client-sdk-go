package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangecontactphonenumbercolumn
type Dialercontactlistconfigchangecontactphonenumbercolumn struct { 
	// ColumnName - name of the phone column
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - type of the phone column, for example, 'cell' or 'home'
	VarType *string `json:"type,omitempty"`


	// CallableTimeColumn - name of the column indicating the timezone to be considered for determining callable times
	CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangecontactphonenumbercolumn
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		VarType: o.VarType,
		
		CallableTimeColumn: o.CallableTimeColumn,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangecontactphonenumbercolumnMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangecontactphonenumbercolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if VarType, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if CallableTimeColumn, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["callableTimeColumn"].(string); ok {
		o.CallableTimeColumn = &CallableTimeColumn
	}
    
	if AdditionalProperties, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
