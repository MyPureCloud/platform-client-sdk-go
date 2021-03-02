package platformclientv2
import (
	"encoding/json"
)

// Atzmtimeslotwithtimezone
type Atzmtimeslotwithtimezone struct { 
	// EarliestCallableTime - The earliest time to dial a contact. Valid format is HH:mm
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime - The latest time to dial a contact. Valid format is HH:mm
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`


	// TimeZoneId - The time zone to use for contacts that cannot be mapped.
	TimeZoneId *string `json:"timeZoneId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Atzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
