package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationstatusinfo - Status information for an Integration.
type Integrationstatusinfo struct { 
	// Code - Machine-readable status as reported by the integration.
	Code *string `json:"code,omitempty"`


	// Effective - Localized, human-readable, effective status of the integration.
	Effective *string `json:"effective,omitempty"`


	// Detail - Localizable status details for the integration.
	Detail *Messageinfo `json:"detail,omitempty"`


	// LastUpdated - Date and time (in UTC) when the integration status (i.e. the code field) was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`

}

func (o *Integrationstatusinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationstatusinfo
	
	LastUpdated := new(string)
	if o.LastUpdated != nil {
		
		*LastUpdated = timeutil.Strftime(o.LastUpdated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastUpdated = nil
	}
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Effective *string `json:"effective,omitempty"`
		
		Detail *Messageinfo `json:"detail,omitempty"`
		
		LastUpdated *string `json:"lastUpdated,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Effective: o.Effective,
		
		Detail: o.Detail,
		
		LastUpdated: LastUpdated,
		Alias:    (*Alias)(o),
	})
}

func (o *Integrationstatusinfo) UnmarshalJSON(b []byte) error {
	var IntegrationstatusinfoMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationstatusinfoMap)
	if err != nil {
		return err
	}
	
	if Code, ok := IntegrationstatusinfoMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Effective, ok := IntegrationstatusinfoMap["effective"].(string); ok {
		o.Effective = &Effective
	}
	
	if Detail, ok := IntegrationstatusinfoMap["detail"].(map[string]interface{}); ok {
		DetailString, _ := json.Marshal(Detail)
		json.Unmarshal(DetailString, &o.Detail)
	}
	
	if lastUpdatedString, ok := IntegrationstatusinfoMap["lastUpdated"].(string); ok {
		LastUpdated, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastUpdatedString)
		o.LastUpdated = &LastUpdated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationstatusinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
