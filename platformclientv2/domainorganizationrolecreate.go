package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainorganizationrolecreate
type Domainorganizationrolecreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - role id
	Id *string `json:"id,omitempty"`

	// Name - The role name
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// DefaultRoleId
	DefaultRoleId *string `json:"defaultRoleId,omitempty"`

	// Permissions
	Permissions *[]string `json:"permissions,omitempty"`

	// UnusedPermissions - A collection of the permissions the role is not using
	UnusedPermissions *[]string `json:"unusedPermissions,omitempty"`

	// PermissionPolicies
	PermissionPolicies *[]Domainpermissionpolicy `json:"permissionPolicies,omitempty"`

	// UserCount
	UserCount *int `json:"userCount,omitempty"`

	// RoleNeedsUpdate - Optional unless patch operation.
	RoleNeedsUpdate *bool `json:"roleNeedsUpdate,omitempty"`

	// BaseLicense
	BaseLicense *string `json:"baseLicense,omitempty"`

	// AddonLicenses
	AddonLicenses *[]string `json:"addonLicenses,omitempty"`

	// DateLicenseLastUpdated - The time that this role licenses were most recently updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLicenseLastUpdated *time.Time `json:"dateLicenseLastUpdated,omitempty"`

	// VarDefault
	VarDefault *bool `json:"default,omitempty"`

	// Base
	Base *bool `json:"base,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainorganizationrolecreate) SetField(field string, fieldValue interface{}) {
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

func (o Domainorganizationrolecreate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLicenseLastUpdated", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Domainorganizationrolecreate
	
	DateLicenseLastUpdated := new(string)
	if o.DateLicenseLastUpdated != nil {
		
		*DateLicenseLastUpdated = timeutil.Strftime(o.DateLicenseLastUpdated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLicenseLastUpdated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultRoleId *string `json:"defaultRoleId,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		UnusedPermissions *[]string `json:"unusedPermissions,omitempty"`
		
		PermissionPolicies *[]Domainpermissionpolicy `json:"permissionPolicies,omitempty"`
		
		UserCount *int `json:"userCount,omitempty"`
		
		RoleNeedsUpdate *bool `json:"roleNeedsUpdate,omitempty"`
		
		BaseLicense *string `json:"baseLicense,omitempty"`
		
		AddonLicenses *[]string `json:"addonLicenses,omitempty"`
		
		DateLicenseLastUpdated *string `json:"dateLicenseLastUpdated,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		
		Base *bool `json:"base,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DefaultRoleId: o.DefaultRoleId,
		
		Permissions: o.Permissions,
		
		UnusedPermissions: o.UnusedPermissions,
		
		PermissionPolicies: o.PermissionPolicies,
		
		UserCount: o.UserCount,
		
		RoleNeedsUpdate: o.RoleNeedsUpdate,
		
		BaseLicense: o.BaseLicense,
		
		AddonLicenses: o.AddonLicenses,
		
		DateLicenseLastUpdated: DateLicenseLastUpdated,
		
		VarDefault: o.VarDefault,
		
		Base: o.Base,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Domainorganizationrolecreate) UnmarshalJSON(b []byte) error {
	var DomainorganizationrolecreateMap map[string]interface{}
	err := json.Unmarshal(b, &DomainorganizationrolecreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainorganizationrolecreateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DomainorganizationrolecreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := DomainorganizationrolecreateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DefaultRoleId, ok := DomainorganizationrolecreateMap["defaultRoleId"].(string); ok {
		o.DefaultRoleId = &DefaultRoleId
	}
    
	if Permissions, ok := DomainorganizationrolecreateMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if UnusedPermissions, ok := DomainorganizationrolecreateMap["unusedPermissions"].([]interface{}); ok {
		UnusedPermissionsString, _ := json.Marshal(UnusedPermissions)
		json.Unmarshal(UnusedPermissionsString, &o.UnusedPermissions)
	}
	
	if PermissionPolicies, ok := DomainorganizationrolecreateMap["permissionPolicies"].([]interface{}); ok {
		PermissionPoliciesString, _ := json.Marshal(PermissionPolicies)
		json.Unmarshal(PermissionPoliciesString, &o.PermissionPolicies)
	}
	
	if UserCount, ok := DomainorganizationrolecreateMap["userCount"].(float64); ok {
		UserCountInt := int(UserCount)
		o.UserCount = &UserCountInt
	}
	
	if RoleNeedsUpdate, ok := DomainorganizationrolecreateMap["roleNeedsUpdate"].(bool); ok {
		o.RoleNeedsUpdate = &RoleNeedsUpdate
	}
    
	if BaseLicense, ok := DomainorganizationrolecreateMap["baseLicense"].(string); ok {
		o.BaseLicense = &BaseLicense
	}
    
	if AddonLicenses, ok := DomainorganizationrolecreateMap["addonLicenses"].([]interface{}); ok {
		AddonLicensesString, _ := json.Marshal(AddonLicenses)
		json.Unmarshal(AddonLicensesString, &o.AddonLicenses)
	}
	
	if dateLicenseLastUpdatedString, ok := DomainorganizationrolecreateMap["dateLicenseLastUpdated"].(string); ok {
		DateLicenseLastUpdated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLicenseLastUpdatedString)
		o.DateLicenseLastUpdated = &DateLicenseLastUpdated
	}
	
	if VarDefault, ok := DomainorganizationrolecreateMap["default"].(bool); ok {
		o.VarDefault = &VarDefault
	}
    
	if Base, ok := DomainorganizationrolecreateMap["base"].(bool); ok {
		o.Base = &Base
	}
    
	if SelfUri, ok := DomainorganizationrolecreateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainorganizationrolecreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
