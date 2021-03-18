package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Lineintegration
type Lineintegration struct { 
	// Id - A unique Integration Id
	Id *string `json:"id,omitempty"`


	// Name - The name of the LINE Integration
	Name *string `json:"name,omitempty"`


	// ChannelId - The Channel Id from LINE messenger
	ChannelId *string `json:"channelId,omitempty"`


	// WebhookUri - The Webhook URI to be updated in LINE platform
	WebhookUri *string `json:"webhookUri,omitempty"`


	// Status - The status of the LINE Integration
	Status *string `json:"status,omitempty"`


	// Recipient - The recipient associated to the Line Integration. This recipient is used to associate a flow to an integration
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
func (o *Lineintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
