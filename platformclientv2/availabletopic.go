package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletopic
type Availabletopic struct { 
	// Description
	Description *string `json:"description,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// PermissionDetails - Full detailed permissions required to subscribe to the topic
	PermissionDetails *[]Permissiondetails `json:"permissionDetails,omitempty"`


	// RequiresPermissions - Permissions required to subscribe to the topic
	RequiresPermissions *[]string `json:"requiresPermissions,omitempty"`


	// RequiresDivisionPermissions - True if the subscribing user must belong to the same division as the topic object ID
	RequiresDivisionPermissions *bool `json:"requiresDivisionPermissions,omitempty"`


	// Enforced - Whether or not the permissions on this topic are enforced
	Enforced *bool `json:"enforced,omitempty"`


	// Visibility - Visibility of this topic (Public or Preview)
	Visibility *string `json:"visibility,omitempty"`


	// Schema
	Schema *map[string]interface{} `json:"schema,omitempty"`


	// RequiresCurrentUser - True if the topic user ID is required to match the subscribing user ID
	RequiresCurrentUser *bool `json:"requiresCurrentUser,omitempty"`


	// RequiresCurrentUserOrPermission - True if permissions are only required when the topic user ID does not match the subscribing user ID
	RequiresCurrentUserOrPermission *bool `json:"requiresCurrentUserOrPermission,omitempty"`


	// Transports - Transports that support events for the topic
	Transports *[]string `json:"transports,omitempty"`


	// PublicApiTemplateUriPaths
	PublicApiTemplateUriPaths *[]string `json:"publicApiTemplateUriPaths,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availabletopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
