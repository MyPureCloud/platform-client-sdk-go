package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Recipient
type Recipient struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Flow - An automate flow object which defines the set of actions to be taken, when a message is received by this provisioned phone number.
	Flow *Flow `json:"flow,omitempty"`


	// DateCreated - Date this recipient was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this recipient was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User that created this recipient
	CreatedBy *User `json:"createdBy,omitempty"`


	// ModifiedBy - User that modified this recipient
	ModifiedBy *User `json:"modifiedBy,omitempty"`


	// MessengerType - The messenger type for this recipient
	MessengerType *string `json:"messengerType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
