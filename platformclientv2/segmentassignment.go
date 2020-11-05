package platformclientv2
import (
	"time"
	"encoding/json"
)

// Segmentassignment
type Segmentassignment struct { 
	// Id - Unique identifier for the segment assignment.
	Id *string `json:"id,omitempty"`


	// State - State of the segment assignment.
	State *string `json:"state,omitempty"`


	// DateAssigned - Date when the segment was assigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssigned *time.Time `json:"dateAssigned,omitempty"`


	// DateUnassigned - Date when the segment was unassigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateUnassigned *time.Time `json:"dateUnassigned,omitempty"`


	// DateModified - Date when the segment assignment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Segment - The segment the assignment is for.
	Segment *Segmentassignmentsegment `json:"segment,omitempty"`


	// CustomerId - ID of the customer to which the segment is assigned.
	CustomerId *string `json:"customerId,omitempty"`


	// CustomerIdType - Type of customer ID (e.g. cookie, email, phone).
	CustomerIdType *string `json:"customerIdType,omitempty"`


	// Session - For session-scoped segments, the session for which the segment was assigned.
	Session *Segmentassignmentsession `json:"session,omitempty"`


	// ExternalContact - External contact of the customer to which the segment is assigned.
	ExternalContact *Addressableentityref `json:"externalContact,omitempty"`

}

// String returns a JSON representation of the model
func (o *Segmentassignment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
