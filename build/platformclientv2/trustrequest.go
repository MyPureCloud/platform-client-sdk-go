package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustrequest
type Trustrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// CreatedBy - User who created this request.
	CreatedBy *Orguser `json:"createdBy,omitempty"`


	// DateCreated - Date request was created. There is a 48 hour expiration on all requests. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// Trustee - Trustee organization who generated this request.
	Trustee *Organization `json:"trustee,omitempty"`


	// Users - The list of trustee users that are requesting access.
	Users *[]Orguser `json:"users,omitempty"`


	// Groups - The list of trustee groups that are requesting access.
	Groups *[]Trustgroup `json:"groups,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
