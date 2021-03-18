package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Encryptionkey
type Encryptionkey struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CreateDate - create date of the key pair. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`


	// KeydataSummary - key data summary (base 64 encoded public key)
	KeydataSummary *string `json:"keydataSummary,omitempty"`


	// User - user that requested generation of public key
	User *User `json:"user,omitempty"`


	// LocalEncryptionConfiguration - Local configuration
	LocalEncryptionConfiguration *Localencryptionconfiguration `json:"localEncryptionConfiguration,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Encryptionkey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
