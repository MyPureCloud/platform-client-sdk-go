package platformclientv2
import (
	"time"
	"encoding/json"
)

// Billingusageresource
type Billingusageresource struct { 
	// Name - Identifies the resource (e.g. license user, device).
	Name *string `json:"name,omitempty"`


	// Date - The date that the usage was first observed by the billing subsystem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Date *time.Time `json:"date,omitempty"`

}

// String returns a JSON representation of the model
func (o *Billingusageresource) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
