package platformclientv2
import (
	"time"
	"encoding/json"
)

// Outboundsettings
type Outboundsettings struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int32 `json:"version,omitempty"`


	// MaxCallsPerAgent - The maximum number of calls that can be placed per agent on any campaign
	MaxCallsPerAgent *int32 `json:"maxCallsPerAgent,omitempty"`


	// MaxConfigurableCallsPerAgent - The maximum number of calls that can be configured to be placed per agent on any campaign
	MaxConfigurableCallsPerAgent *int32 `json:"maxConfigurableCallsPerAgent,omitempty"`


	// MaxLineUtilization - The maximum percentage of lines that should be used for Outbound, expressed as a decimal in the range [0.0, 1.0]
	MaxLineUtilization *float64 `json:"maxLineUtilization,omitempty"`


	// AbandonSeconds - The number of seconds used to determine if a call is abandoned
	AbandonSeconds *float64 `json:"abandonSeconds,omitempty"`


	// ComplianceAbandonRateDenominator - The denominator to be used in determining the compliance abandon rate
	ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
