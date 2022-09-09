package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setcontenttemplateactionsettings
type Setcontenttemplateactionsettings struct { 
	// SmsContentTemplateId - A string of sms contentTemplateId.
	SmsContentTemplateId *string `json:"smsContentTemplateId,omitempty"`


	// EmailContentTemplateId - A string of email contentTemplateId.
	EmailContentTemplateId *string `json:"emailContentTemplateId,omitempty"`

}

func (o *Setcontenttemplateactionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setcontenttemplateactionsettings
	
	return json.Marshal(&struct { 
		SmsContentTemplateId *string `json:"smsContentTemplateId,omitempty"`
		
		EmailContentTemplateId *string `json:"emailContentTemplateId,omitempty"`
		*Alias
	}{ 
		SmsContentTemplateId: o.SmsContentTemplateId,
		
		EmailContentTemplateId: o.EmailContentTemplateId,
		Alias:    (*Alias)(o),
	})
}

func (o *Setcontenttemplateactionsettings) UnmarshalJSON(b []byte) error {
	var SetcontenttemplateactionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SetcontenttemplateactionsettingsMap)
	if err != nil {
		return err
	}
	
	if SmsContentTemplateId, ok := SetcontenttemplateactionsettingsMap["smsContentTemplateId"].(string); ok {
		o.SmsContentTemplateId = &SmsContentTemplateId
	}
    
	if EmailContentTemplateId, ok := SetcontenttemplateactionsettingsMap["emailContentTemplateId"].(string); ok {
		o.EmailContentTemplateId = &EmailContentTemplateId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Setcontenttemplateactionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
