package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lastattemptbycolumnconditionsettings
type Lastattemptbycolumnconditionsettings struct { 
	// EmailColumnName - The name of the contact column to evaluate for Email.
	EmailColumnName *string `json:"emailColumnName,omitempty"`


	// SmsColumnName - The name of the contact column to evaluate for SMS.
	SmsColumnName *string `json:"smsColumnName,omitempty"`


	// Operator - The operator to use when comparing values.
	Operator *string `json:"operator,omitempty"`


	// Value - The period value to compare against the contact's data.
	Value *string `json:"value,omitempty"`

}

func (o *Lastattemptbycolumnconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lastattemptbycolumnconditionsettings
	
	return json.Marshal(&struct { 
		EmailColumnName *string `json:"emailColumnName,omitempty"`
		
		SmsColumnName *string `json:"smsColumnName,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		EmailColumnName: o.EmailColumnName,
		
		SmsColumnName: o.SmsColumnName,
		
		Operator: o.Operator,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Lastattemptbycolumnconditionsettings) UnmarshalJSON(b []byte) error {
	var LastattemptbycolumnconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &LastattemptbycolumnconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if EmailColumnName, ok := LastattemptbycolumnconditionsettingsMap["emailColumnName"].(string); ok {
		o.EmailColumnName = &EmailColumnName
	}
    
	if SmsColumnName, ok := LastattemptbycolumnconditionsettingsMap["smsColumnName"].(string); ok {
		o.SmsColumnName = &SmsColumnName
	}
    
	if Operator, ok := LastattemptbycolumnconditionsettingsMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := LastattemptbycolumnconditionsettingsMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Lastattemptbycolumnconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
