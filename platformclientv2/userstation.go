package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstation
type Userstation struct { 
	// Id - A globally unique identifier for this station
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AssociatedUser
	AssociatedUser *User `json:"associatedUser,omitempty"`


	// AssociatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssociatedDate *time.Time `json:"associatedDate,omitempty"`


	// DefaultUser
	DefaultUser *User `json:"defaultUser,omitempty"`


	// ProviderInfo - Provider-specific info for this station, e.g. { \"edgeGroupId\": \"ffe7b15c-a9cc-4f4c-88f5-781327819a49\" }
	ProviderInfo *map[string]string `json:"providerInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
