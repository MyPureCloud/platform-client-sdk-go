package platformclientv2
import (
	"time"
	"encoding/json"
)

// Outcome
type Outcome struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - Whether or not the outcome is active.
	IsActive *bool `json:"isActive,omitempty"`


	// DisplayName - The display name of the outcome.
	DisplayName *string `json:"displayName,omitempty"`


	// Version - The version of the outcome.
	Version *int `json:"version,omitempty"`


	// Description - A description of the outcome.
	Description *string `json:"description,omitempty"`


	// IsPositive - Whether or not the outcome is positive.
	IsPositive *bool `json:"isPositive,omitempty"`


	// Context - The context of the outcome.
	Context *Context `json:"context,omitempty"`


	// Journey - The pattern of rules defining the filter of the outcome.
	Journey *Journey `json:"journey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Timestamp indicating when the outcome was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Timestamp indicating when the outcome was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outcome) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
