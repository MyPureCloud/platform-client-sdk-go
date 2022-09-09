package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lastresultoverallconditionsettings
type Lastresultoverallconditionsettings struct { 
	// EmailWrapupCodes - A list of wrapup code identifiers to match for Email.
	EmailWrapupCodes *[]string `json:"emailWrapupCodes,omitempty"`


	// SmsWrapupCodes - A list of wrapup code identifiers to match for SMS.
	SmsWrapupCodes *[]string `json:"smsWrapupCodes,omitempty"`

}

func (o *Lastresultoverallconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lastresultoverallconditionsettings
	
	return json.Marshal(&struct { 
		EmailWrapupCodes *[]string `json:"emailWrapupCodes,omitempty"`
		
		SmsWrapupCodes *[]string `json:"smsWrapupCodes,omitempty"`
		*Alias
	}{ 
		EmailWrapupCodes: o.EmailWrapupCodes,
		
		SmsWrapupCodes: o.SmsWrapupCodes,
		Alias:    (*Alias)(o),
	})
}

func (o *Lastresultoverallconditionsettings) UnmarshalJSON(b []byte) error {
	var LastresultoverallconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &LastresultoverallconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if EmailWrapupCodes, ok := LastresultoverallconditionsettingsMap["emailWrapupCodes"].([]interface{}); ok {
		EmailWrapupCodesString, _ := json.Marshal(EmailWrapupCodes)
		json.Unmarshal(EmailWrapupCodesString, &o.EmailWrapupCodes)
	}
	
	if SmsWrapupCodes, ok := LastresultoverallconditionsettingsMap["smsWrapupCodes"].([]interface{}); ok {
		SmsWrapupCodesString, _ := json.Marshal(SmsWrapupCodes)
		json.Unmarshal(SmsWrapupCodesString, &o.SmsWrapupCodes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lastresultoverallconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
