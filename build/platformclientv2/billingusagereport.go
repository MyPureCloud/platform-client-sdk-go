package platformclientv2
import (
	"time"
	"encoding/json"
)

// Billingusagereport
type Billingusagereport struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// StartDate - The period start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The period end date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// Status - Generation status of report
	Status *string `json:"status,omitempty"`


	// Usages - The usages for the given period.
	Usages *[]Billingusage `json:"usages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Billingusagereport) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
