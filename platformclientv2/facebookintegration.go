package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookintegration
type Facebookintegration struct { 
	// Id - A unique Integration Id.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Facebook Integration
	Name *string `json:"name,omitempty"`


	// AppId - The App Id from Facebook messenger
	AppId *string `json:"appId,omitempty"`


	// PageId - The Page Id from Facebook messenger
	PageId *string `json:"pageId,omitempty"`


	// PageName - The name of the Facebook page
	PageName *string `json:"pageName,omitempty"`


	// PageProfileImageUrl - The url of the profile image of the Facebook page
	PageProfileImageUrl *string `json:"pageProfileImageUrl,omitempty"`


	// Status - The status of the Facebook Integration
	Status *string `json:"status,omitempty"`


	// Recipient - The recipient reference associated to the Facebook Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// CreateStatus - Status of asynchronous create operation
	CreateStatus *string `json:"createStatus,omitempty"`


	// CreateError - Error information returned, if createStatus is set to Error
	CreateError *Errorbody `json:"createError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
