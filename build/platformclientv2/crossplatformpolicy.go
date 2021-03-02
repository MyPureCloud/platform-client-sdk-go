package platformclientv2
import (
	"time"
	"encoding/json"
)

// Crossplatformpolicy
type Crossplatformpolicy struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Order
	Order *int `json:"order,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// MediaPolicies - Conditions and actions per media type
	MediaPolicies *Crossplatformmediapolicies `json:"mediaPolicies,omitempty"`


	// Conditions - Conditions
	Conditions *Policyconditions `json:"conditions,omitempty"`


	// Actions - Actions
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`


	// PolicyErrors
	PolicyErrors *Policyerrors `json:"policyErrors,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Crossplatformpolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
