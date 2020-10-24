package platformclientv2
import (
	"time"
	"encoding/json"
)

// Propertyindexrequest
type Propertyindexrequest struct { 
	// SessionId - Attach properties to a segment in the indicated session
	SessionId *string `json:"sessionId,omitempty"`


	// TargetDate - Attach properties to a segment covering a specific point in time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TargetDate *time.Time `json:"targetDate,omitempty"`


	// Properties - The list of properties to index
	Properties *[]Analyticsproperty `json:"properties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Propertyindexrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
