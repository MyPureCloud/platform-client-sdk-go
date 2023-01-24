package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userauthorization
type Userauthorization struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Roles
	Roles *[]Domainrole `json:"roles,omitempty"`

	// UnusedRoles - A collection of the roles the user is not using
	UnusedRoles *[]Domainrole `json:"unusedRoles,omitempty"`

	// Permissions - A collection of the permissions granted by all assigned roles
	Permissions *[]string `json:"permissions,omitempty"`

	// PermissionPolicies - The policies configured for assigned permissions.
	PermissionPolicies *[]Resourcepermissionpolicy `json:"permissionPolicies,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userauthorization) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Userauthorization) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userauthorization
	
	return json.Marshal(&struct { 
		Roles *[]Domainrole `json:"roles,omitempty"`
		
		UnusedRoles *[]Domainrole `json:"unusedRoles,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		PermissionPolicies *[]Resourcepermissionpolicy `json:"permissionPolicies,omitempty"`
		Alias
	}{ 
		Roles: o.Roles,
		
		UnusedRoles: o.UnusedRoles,
		
		Permissions: o.Permissions,
		
		PermissionPolicies: o.PermissionPolicies,
		Alias:    (Alias)(o),
	})
}

func (o *Userauthorization) UnmarshalJSON(b []byte) error {
	var UserauthorizationMap map[string]interface{}
	err := json.Unmarshal(b, &UserauthorizationMap)
	if err != nil {
		return err
	}
	
	if Roles, ok := UserauthorizationMap["roles"].([]interface{}); ok {
		RolesString, _ := json.Marshal(Roles)
		json.Unmarshal(RolesString, &o.Roles)
	}
	
	if UnusedRoles, ok := UserauthorizationMap["unusedRoles"].([]interface{}); ok {
		UnusedRolesString, _ := json.Marshal(UnusedRoles)
		json.Unmarshal(UnusedRolesString, &o.UnusedRoles)
	}
	
	if Permissions, ok := UserauthorizationMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if PermissionPolicies, ok := UserauthorizationMap["permissionPolicies"].([]interface{}); ok {
		PermissionPoliciesString, _ := json.Marshal(PermissionPolicies)
		json.Unmarshal(PermissionPoliciesString, &o.PermissionPolicies)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
