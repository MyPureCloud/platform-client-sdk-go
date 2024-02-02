package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupupdate
type Groupupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The group name.
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// State - State of the group.
	State *string `json:"state,omitempty"`

	// Version - Current version for this resource.
	Version *int `json:"version,omitempty"`

	// Images
	Images *[]Userimage `json:"images,omitempty"`

	// Addresses
	Addresses *[]Groupcontact `json:"addresses,omitempty"`

	// RulesVisible - Are membership rules visible to the person requesting to view the group
	RulesVisible *bool `json:"rulesVisible,omitempty"`

	// Visibility - Who can view this group
	Visibility *string `json:"visibility,omitempty"`

	// RolesEnabled - Allow roles to be assigned to this group
	RolesEnabled *bool `json:"rolesEnabled,omitempty"`

	// OwnerIds - Owners of the group
	OwnerIds *[]string `json:"ownerIds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Groupupdate) SetField(field string, fieldValue interface{}) {
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

func (o Groupupdate) MarshalJSON() ([]byte, error) {
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
	type Alias Groupupdate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Addresses *[]Groupcontact `json:"addresses,omitempty"`
		
		RulesVisible *bool `json:"rulesVisible,omitempty"`
		
		Visibility *string `json:"visibility,omitempty"`
		
		RolesEnabled *bool `json:"rolesEnabled,omitempty"`
		
		OwnerIds *[]string `json:"ownerIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		State: o.State,
		
		Version: o.Version,
		
		Images: o.Images,
		
		Addresses: o.Addresses,
		
		RulesVisible: o.RulesVisible,
		
		Visibility: o.Visibility,
		
		RolesEnabled: o.RolesEnabled,
		
		OwnerIds: o.OwnerIds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Groupupdate) UnmarshalJSON(b []byte) error {
	var GroupupdateMap map[string]interface{}
	err := json.Unmarshal(b, &GroupupdateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GroupupdateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := GroupupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := GroupupdateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if State, ok := GroupupdateMap["state"].(string); ok {
		o.State = &State
	}
    
	if Version, ok := GroupupdateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Images, ok := GroupupdateMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Addresses, ok := GroupupdateMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if RulesVisible, ok := GroupupdateMap["rulesVisible"].(bool); ok {
		o.RulesVisible = &RulesVisible
	}
    
	if Visibility, ok := GroupupdateMap["visibility"].(string); ok {
		o.Visibility = &Visibility
	}
    
	if RolesEnabled, ok := GroupupdateMap["rolesEnabled"].(bool); ok {
		o.RolesEnabled = &RolesEnabled
	}
    
	if OwnerIds, ok := GroupupdateMap["ownerIds"].([]interface{}); ok {
		OwnerIdsString, _ := json.Marshal(OwnerIds)
		json.Unmarshal(OwnerIdsString, &o.OwnerIds)
	}
	
	if SelfUri, ok := GroupupdateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Groupupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
