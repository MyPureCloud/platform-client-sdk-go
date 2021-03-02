package platformclientv2
import (
	"time"
	"encoding/json"
)

// Trustor
type Trustor struct { 
	// Id - Organization Id for this trust.
	Id *string `json:"id,omitempty"`


	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// DateCreated - Date Trust was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - User that created trust.
	CreatedBy *Orguser `json:"createdBy,omitempty"`


	// Organization - Organization associated with this trust.
	Organization *Organization `json:"organization,omitempty"`


	// Authorization - Authorization for the trustee user has in this trustor organization
	Authorization *Trusteeauthorization `json:"authorization,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustor) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
