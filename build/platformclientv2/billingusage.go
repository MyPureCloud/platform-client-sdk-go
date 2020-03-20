package platformclientv2
import (
	"encoding/json"
)

// Billingusage
type Billingusage struct { 
	// Name - Identifies the billable usage.
	Name *string `json:"name,omitempty"`


	// TotalUsage - The total amount of usage, expressed as a decimal number in string format.
	TotalUsage *string `json:"totalUsage,omitempty"`


	// Resources - The resources for which usage was observed (e.g. license users, devices).
	Resources *[]Billingusageresource `json:"resources,omitempty"`

}

// String returns a JSON representation of the model
func (o *Billingusage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
