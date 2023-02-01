package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainorgroledifference
type Domainorgroledifference struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RemovedPermissionPolicies
	RemovedPermissionPolicies *[]Domainpermissionpolicy `json:"removedPermissionPolicies,omitempty"`

	// AddedPermissionPolicies
	AddedPermissionPolicies *[]Domainpermissionpolicy `json:"addedPermissionPolicies,omitempty"`

	// SamePermissionPolicies
	SamePermissionPolicies *[]Domainpermissionpolicy `json:"samePermissionPolicies,omitempty"`

	// UserOrgRole
	UserOrgRole *Domainorganizationrole `json:"userOrgRole,omitempty"`

	// RoleFromDefault
	RoleFromDefault *Domainorganizationrole `json:"roleFromDefault,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainorgroledifference) SetField(field string, fieldValue interface{}) {
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

func (o Domainorgroledifference) MarshalJSON() ([]byte, error) {
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
	type Alias Domainorgroledifference
	
	return json.Marshal(&struct { 
		RemovedPermissionPolicies *[]Domainpermissionpolicy `json:"removedPermissionPolicies,omitempty"`
		
		AddedPermissionPolicies *[]Domainpermissionpolicy `json:"addedPermissionPolicies,omitempty"`
		
		SamePermissionPolicies *[]Domainpermissionpolicy `json:"samePermissionPolicies,omitempty"`
		
		UserOrgRole *Domainorganizationrole `json:"userOrgRole,omitempty"`
		
		RoleFromDefault *Domainorganizationrole `json:"roleFromDefault,omitempty"`
		Alias
	}{ 
		RemovedPermissionPolicies: o.RemovedPermissionPolicies,
		
		AddedPermissionPolicies: o.AddedPermissionPolicies,
		
		SamePermissionPolicies: o.SamePermissionPolicies,
		
		UserOrgRole: o.UserOrgRole,
		
		RoleFromDefault: o.RoleFromDefault,
		Alias:    (Alias)(o),
	})
}

func (o *Domainorgroledifference) UnmarshalJSON(b []byte) error {
	var DomainorgroledifferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DomainorgroledifferenceMap)
	if err != nil {
		return err
	}
	
	if RemovedPermissionPolicies, ok := DomainorgroledifferenceMap["removedPermissionPolicies"].([]interface{}); ok {
		RemovedPermissionPoliciesString, _ := json.Marshal(RemovedPermissionPolicies)
		json.Unmarshal(RemovedPermissionPoliciesString, &o.RemovedPermissionPolicies)
	}
	
	if AddedPermissionPolicies, ok := DomainorgroledifferenceMap["addedPermissionPolicies"].([]interface{}); ok {
		AddedPermissionPoliciesString, _ := json.Marshal(AddedPermissionPolicies)
		json.Unmarshal(AddedPermissionPoliciesString, &o.AddedPermissionPolicies)
	}
	
	if SamePermissionPolicies, ok := DomainorgroledifferenceMap["samePermissionPolicies"].([]interface{}); ok {
		SamePermissionPoliciesString, _ := json.Marshal(SamePermissionPolicies)
		json.Unmarshal(SamePermissionPoliciesString, &o.SamePermissionPolicies)
	}
	
	if UserOrgRole, ok := DomainorgroledifferenceMap["userOrgRole"].(map[string]interface{}); ok {
		UserOrgRoleString, _ := json.Marshal(UserOrgRole)
		json.Unmarshal(UserOrgRoleString, &o.UserOrgRole)
	}
	
	if RoleFromDefault, ok := DomainorgroledifferenceMap["roleFromDefault"].(map[string]interface{}); ok {
		RoleFromDefaultString, _ := json.Marshal(RoleFromDefault)
		json.Unmarshal(RoleFromDefaultString, &o.RoleFromDefault)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainorgroledifference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
