package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2group - Defines a SCIM group.
type Scimv2group struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`

	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`

	// DisplayName - The display name of the group.
	DisplayName *string `json:"displayName,omitempty"`

	// ExternalId - The external ID of the group. Set by the provisioning client. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readWrite\".
	ExternalId *string `json:"externalId,omitempty"`

	// Members - The list of members in the group.
	Members *[]Scimv2memberreference `json:"members,omitempty"`

	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimv2group) SetField(field string, fieldValue interface{}) {
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

func (o Scimv2group) MarshalJSON() ([]byte, error) {
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
	type Alias Scimv2group
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schemas *[]string `json:"schemas,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Members *[]Scimv2memberreference `json:"members,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Schemas: o.Schemas,
		
		DisplayName: o.DisplayName,
		
		ExternalId: o.ExternalId,
		
		Members: o.Members,
		
		Meta: o.Meta,
		Alias:    (Alias)(o),
	})
}

func (o *Scimv2group) UnmarshalJSON(b []byte) error {
	var Scimv2groupMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2groupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := Scimv2groupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Schemas, ok := Scimv2groupMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if DisplayName, ok := Scimv2groupMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if ExternalId, ok := Scimv2groupMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if Members, ok := Scimv2groupMap["members"].([]interface{}); ok {
		MembersString, _ := json.Marshal(Members)
		json.Unmarshal(MembersString, &o.Members)
	}
	
	if Meta, ok := Scimv2groupMap["meta"].(map[string]interface{}); ok {
		MetaString, _ := json.Marshal(Meta)
		json.Unmarshal(MetaString, &o.Meta)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2group) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
