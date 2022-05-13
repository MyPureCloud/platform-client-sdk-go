package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mailfromresult
type Mailfromresult struct { 
	// Status - The verification status.
	Status *string `json:"status,omitempty"`


	// Records - The list of DNS records that pertain that need to exist for verification.
	Records *[]Record `json:"records,omitempty"`


	// MailFromDomain - The custom MAIL FROM domain.
	MailFromDomain *string `json:"mailFromDomain,omitempty"`

}

func (o *Mailfromresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mailfromresult
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Records *[]Record `json:"records,omitempty"`
		
		MailFromDomain *string `json:"mailFromDomain,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Records: o.Records,
		
		MailFromDomain: o.MailFromDomain,
		Alias:    (*Alias)(o),
	})
}

func (o *Mailfromresult) UnmarshalJSON(b []byte) error {
	var MailfromresultMap map[string]interface{}
	err := json.Unmarshal(b, &MailfromresultMap)
	if err != nil {
		return err
	}
	
	if Status, ok := MailfromresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Records, ok := MailfromresultMap["records"].([]interface{}); ok {
		RecordsString, _ := json.Marshal(Records)
		json.Unmarshal(RecordsString, &o.Records)
	}
	
	if MailFromDomain, ok := MailfromresultMap["mailFromDomain"].(string); ok {
		o.MailFromDomain = &MailFromDomain
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mailfromresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
