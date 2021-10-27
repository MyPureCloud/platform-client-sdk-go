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


	// TopicParameters - Parameters in the topic name that can be substituted, in the order they appear in the topic name
	TopicParameters *[]string `json:"topicParameters,omitempty"`

}

func (o *Availabletopic) MarshalJSON() ([]byte, error) {
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
		
		TopicParameters *[]string `json:"topicParameters,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		Id: o.Id,
		
		PermissionDetails: o.PermissionDetails,
		
		RequiresPermissions: o.RequiresPermissions,
		
		RequiresDivisionPermissions: o.RequiresDivisionPermissions,
		
		RequiresAnyValidator: o.RequiresAnyValidator,
		
		Enforced: o.Enforced,
		
		Visibility: o.Visibility,
		
		Schema: o.Schema,
		
		RequiresCurrentUser: o.RequiresCurrentUser,
		
		RequiresCurrentUserOrPermission: o.RequiresCurrentUserOrPermission,
		
		Transports: o.Transports,
		
		PublicApiTemplateUriPaths: o.PublicApiTemplateUriPaths,
		
		TopicParameters: o.TopicParameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletopic) UnmarshalJSON(b []byte) error {
	var AvailabletopicMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletopicMap)
	if err != nil {
		return err
	}
	
	if Description, ok := AvailabletopicMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Id, ok := AvailabletopicMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if PermissionDetails, ok := AvailabletopicMap["permissionDetails"].([]interface{}); ok {
		PermissionDetailsString, _ := json.Marshal(PermissionDetails)
		json.Unmarshal(PermissionDetailsString, &o.PermissionDetails)
	}
	
	if RequiresPermissions, ok := AvailabletopicMap["requiresPermissions"].([]interface{}); ok {
		RequiresPermissionsString, _ := json.Marshal(RequiresPermissions)
		json.Unmarshal(RequiresPermissionsString, &o.RequiresPermissions)
	}
	
	if RequiresDivisionPermissions, ok := AvailabletopicMap["requiresDivisionPermissions"].(bool); ok {
		o.RequiresDivisionPermissions = &RequiresDivisionPermissions
	}
	
	if RequiresAnyValidator, ok := AvailabletopicMap["requiresAnyValidator"].(bool); ok {
		o.RequiresAnyValidator = &RequiresAnyValidator
	}
	
	if Enforced, ok := AvailabletopicMap["enforced"].(bool); ok {
		o.Enforced = &Enforced
	}
	
	if Visibility, ok := AvailabletopicMap["visibility"].(string); ok {
		o.Visibility = &Visibility
	}
	
	if Schema, ok := AvailabletopicMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if RequiresCurrentUser, ok := AvailabletopicMap["requiresCurrentUser"].(bool); ok {
		o.RequiresCurrentUser = &RequiresCurrentUser
	}
	
	if RequiresCurrentUserOrPermission, ok := AvailabletopicMap["requiresCurrentUserOrPermission"].(bool); ok {
		o.RequiresCurrentUserOrPermission = &RequiresCurrentUserOrPermission
	}
	
	if Transports, ok := AvailabletopicMap["transports"].([]interface{}); ok {
		TransportsString, _ := json.Marshal(Transports)
		json.Unmarshal(TransportsString, &o.Transports)
	}
	
	if PublicApiTemplateUriPaths, ok := AvailabletopicMap["publicApiTemplateUriPaths"].([]interface{}); ok {
		PublicApiTemplateUriPathsString, _ := json.Marshal(PublicApiTemplateUriPaths)
		json.Unmarshal(PublicApiTemplateUriPathsString, &o.PublicApiTemplateUriPaths)
	}
	
	if TopicParameters, ok := AvailabletopicMap["topicParameters"].([]interface{}); ok {
		TopicParametersString, _ := json.Marshal(TopicParameters)
		json.Unmarshal(TopicParametersString, &o.TopicParameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
