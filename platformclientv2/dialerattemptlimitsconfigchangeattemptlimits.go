package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialerattemptlimitsconfigchangeattemptlimits
type Dialerattemptlimitsconfigchangeattemptlimits struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`


	// MaxAttemptsPerContact
	MaxAttemptsPerContact *int32 `json:"maxAttemptsPerContact,omitempty"`


	// MaxAttemptsPerNumber
	MaxAttemptsPerNumber *int32 `json:"maxAttemptsPerNumber,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// ResetPeriod
	ResetPeriod *string `json:"resetPeriod,omitempty"`


	// RecallEntries
	RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`


	// BreadthFirstRecalls
	BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangeattemptlimits) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
