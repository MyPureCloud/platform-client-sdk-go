package platformclientv2
import (
	"time"
	"encoding/json"
)

// Locationaddressverificationdetails
type Locationaddressverificationdetails struct { 
	// Status - Status of address verification process
	Status *string `json:"status,omitempty"`


	// DateFinished - Finished time of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateFinished *time.Time `json:"dateFinished,omitempty"`


	// DateStarted - Time started of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`


	// Service - Third party service used for address verification
	Service *string `json:"service,omitempty"`

}

// String returns a JSON representation of the model
func (o *Locationaddressverificationdetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
