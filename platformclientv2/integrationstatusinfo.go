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

func (u *Integrationstatusinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationstatusinfo

	
	LastUpdated := new(string)
	if u.LastUpdated != nil {
		
		*LastUpdated = timeutil.Strftime(u.LastUpdated, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Code: u.Code,
		
		Effective: u.Effective,
		
		Detail: u.Detail,
		
		LastUpdated: LastUpdated,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Integrationstatusinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
