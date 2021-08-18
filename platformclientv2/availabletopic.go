package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// RequiresAnyValidator - If multiple permissions are required for this topic, such as both requiresCurrentUser and requiresDivisionPermissions, then true here indicates that meeting any one condition will satisfy the requirements; false indicates all conditions must be met.
	RequiresAnyValidator *bool `json:"requiresAnyValidator,omitempty"`


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

func (u *Availabletopic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletopic

	

	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		PermissionDetails *[]Permissiondetails `json:"permissionDetails,omitempty"`
		
		RequiresPermissions *[]string `json:"requiresPermissions,omitempty"`
		
		RequiresDivisionPermissions *bool `json:"requiresDivisionPermissions,omitempty"`
		
		RequiresAnyValidator *bool `json:"requiresAnyValidator,omitempty"`
		
		Enforced *bool `json:"enforced,omitempty"`
		
		Visibility *string `json:"visibility,omitempty"`
		
		Schema *map[string]interface{} `json:"schema,omitempty"`
		
		RequiresCurrentUser *bool `json:"requiresCurrentUser,omitempty"`
		
		RequiresCurrentUserOrPermission *bool `json:"requiresCurrentUserOrPermission,omitempty"`
		
		Transports *[]string `json:"transports,omitempty"`
		
		PublicApiTemplateUriPaths *[]string `json:"publicApiTemplateUriPaths,omitempty"`
		*Alias
	}{ 
		Description: u.Description,
		
		Id: u.Id,
		
		PermissionDetails: u.PermissionDetails,
		
		RequiresPermissions: u.RequiresPermissions,
		
		RequiresDivisionPermissions: u.RequiresDivisionPermissions,
		
		RequiresAnyValidator: u.RequiresAnyValidator,
		
		Enforced: u.Enforced,
		
		Visibility: u.Visibility,
		
		Schema: u.Schema,
		
		RequiresCurrentUser: u.RequiresCurrentUser,
		
		RequiresCurrentUserOrPermission: u.RequiresCurrentUserOrPermission,
		
		Transports: u.Transports,
		
		PublicApiTemplateUriPaths: u.PublicApiTemplateUriPaths,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Availabletopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
