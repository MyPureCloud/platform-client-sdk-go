package platformclientv2
import (
	"time"
	"encoding/json"
)

// Integrationstatusinfo - Status information for an Integration.
type Integrationstatusinfo struct { 
	// Code - Machine-readable status as reported by the integration.
	Code *string `json:"code,omitempty"`


	// Effective - Localized, human-readable, effective status of the integration.
	Effective *string `json:"effective,omitempty"`


	// Detail - Localizable status details for the integration.
	Detail *Messageinfo `json:"detail,omitempty"`


	// LastUpdated - Date and time (in UTC) when the integration status (i.e. the code field) was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integrationstatusinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
