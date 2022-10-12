package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Verificationresult
type Verificationresult struct { 
	// Status - The verification status.
	Status *string `json:"status,omitempty"`


	// Records - The list of DNS records that pertain that need to exist for verification.
	Records *[]Record `json:"records,omitempty"`

}

func (o *Verificationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Verificationresult
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Records *[]Record `json:"records,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Records: o.Records,
		Alias:    (*Alias)(o),
	})
}

func (o *Verificationresult) UnmarshalJSON(b []byte) error {
	var VerificationresultMap map[string]interface{}
	err := json.Unmarshal(b, &VerificationresultMap)
	if err != nil {
		return err
	}
	
	if Status, ok := VerificationresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Records, ok := VerificationresultMap["records"].([]interface{}); ok {
		RecordsString, _ := json.Marshal(Records)
		json.Unmarshal(RecordsString, &o.Records)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Verificationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
