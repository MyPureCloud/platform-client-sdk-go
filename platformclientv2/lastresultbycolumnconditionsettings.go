package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lastresultbycolumnconditionsettings
type Lastresultbycolumnconditionsettings struct { 
	// EmailColumnName - The name of the contact column to evaluate for Email.
	EmailColumnName *string `json:"emailColumnName,omitempty"`


	// EmailWrapupCodes - A list of wrapup code identifiers to match for Email.
	EmailWrapupCodes *[]string `json:"emailWrapupCodes,omitempty"`


	// SmsColumnName - The name of the contact column to evaluate for SMS.
	SmsColumnName *string `json:"smsColumnName,omitempty"`


	// SmsWrapupCodes - A list of wrapup code identifiers to match for SMS.
	SmsWrapupCodes *[]string `json:"smsWrapupCodes,omitempty"`

}

func (o *Lastresultbycolumnconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lastresultbycolumnconditionsettings
	
	return json.Marshal(&struct { 
		EmailColumnName *string `json:"emailColumnName,omitempty"`
		
		EmailWrapupCodes *[]string `json:"emailWrapupCodes,omitempty"`
		
		SmsColumnName *string `json:"smsColumnName,omitempty"`
		
		SmsWrapupCodes *[]string `json:"smsWrapupCodes,omitempty"`
		*Alias
	}{ 
		EmailColumnName: o.EmailColumnName,
		
		EmailWrapupCodes: o.EmailWrapupCodes,
		
		SmsColumnName: o.SmsColumnName,
		
		SmsWrapupCodes: o.SmsWrapupCodes,
		Alias:    (*Alias)(o),
	})
}

func (o *Lastresultbycolumnconditionsettings) UnmarshalJSON(b []byte) error {
	var LastresultbycolumnconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &LastresultbycolumnconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if EmailColumnName, ok := LastresultbycolumnconditionsettingsMap["emailColumnName"].(string); ok {
		o.EmailColumnName = &EmailColumnName
	}
    
	if EmailWrapupCodes, ok := LastresultbycolumnconditionsettingsMap["emailWrapupCodes"].([]interface{}); ok {
		EmailWrapupCodesString, _ := json.Marshal(EmailWrapupCodes)
		json.Unmarshal(EmailWrapupCodesString, &o.EmailWrapupCodes)
	}
	
	if SmsColumnName, ok := LastresultbycolumnconditionsettingsMap["smsColumnName"].(string); ok {
		o.SmsColumnName = &SmsColumnName
	}
    
	if SmsWrapupCodes, ok := LastresultbycolumnconditionsettingsMap["smsWrapupCodes"].([]interface{}); ok {
		SmsWrapupCodesString, _ := json.Marshal(SmsWrapupCodes)
		json.Unmarshal(SmsWrapupCodesString, &o.SmsWrapupCodes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lastresultbycolumnconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
