package platformclientv2
import (
	"time"
	"encoding/json"
)

// Attemptlimits
type Attemptlimits struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// MaxAttemptsPerContact - The maximum number of times a contact can be called within the resetPeriod. Required if maxAttemptsPerNumber is not defined.
	MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`


	// MaxAttemptsPerNumber - The maximum number of times a phone number can be called within the resetPeriod. Required if maxAttemptsPerContact is not defined.
	MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`


	// TimeZoneId - If the resetPeriod is TODAY, this specifies the timezone in which TODAY occurs. Required if the resetPeriod is TODAY.
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// ResetPeriod - After how long the number of attempts will be set back to 0. Defaults to NEVER.
	ResetPeriod *string `json:"resetPeriod,omitempty"`


	// RecallEntries - Configuration for recall attempts.
	RecallEntries *map[string]Recallentry `json:"recallEntries,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Attemptlimits) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
