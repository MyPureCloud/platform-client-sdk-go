package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupcreate
type Groupcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The group name.
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// DateModified - Last modified date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// MemberCount - Number of members.
	MemberCount *int `json:"memberCount,omitempty"`

	// State - Active, inactive, or deleted state.
	State *string `json:"state,omitempty"`

	// Version - Current version for this resource.
	Version *int `json:"version,omitempty"`

	// VarType - Type of group.
	VarType *string `json:"type,omitempty"`

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

	// IncludeOwners - Allow owners to be included as members of the group
	IncludeOwners *bool `json:"includeOwners,omitempty"`

	// CallsEnabled - Allow calls to be placed to this group.
	CallsEnabled *bool `json:"callsEnabled,omitempty"`

	// OwnerIds - Owners of the group
	OwnerIds *[]string `json:"ownerIds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Groupcreate) SetField(field string, fieldValue interface{}) {
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

func (o Groupcreate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateModified", }
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
	type Alias Groupcreate
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Addresses *[]Groupcontact `json:"addresses,omitempty"`
		
		RulesVisible *bool `json:"rulesVisible,omitempty"`
		
		Visibility *string `json:"visibility,omitempty"`
		
		RolesEnabled *bool `json:"rolesEnabled,omitempty"`
		
		IncludeOwners *bool `json:"includeOwners,omitempty"`
		
		CallsEnabled *bool `json:"callsEnabled,omitempty"`
		
		OwnerIds *[]string `json:"ownerIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateModified: DateModified,
		
		MemberCount: o.MemberCount,
		
		State: o.State,
		
		Version: o.Version,
		
		VarType: o.VarType,
		
		Images: o.Images,
		
		Addresses: o.Addresses,
		
		RulesVisible: o.RulesVisible,
		
		Visibility: o.Visibility,
		
		RolesEnabled: o.RolesEnabled,
		
		IncludeOwners: o.IncludeOwners,
		
		CallsEnabled: o.CallsEnabled,
		
		OwnerIds: o.OwnerIds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Groupcreate) UnmarshalJSON(b []byte) error {
	var GroupcreateMap map[string]interface{}
	err := json.Unmarshal(b, &GroupcreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GroupcreateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := GroupcreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := GroupcreateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateModifiedString, ok := GroupcreateMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if MemberCount, ok := GroupcreateMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if State, ok := GroupcreateMap["state"].(string); ok {
		o.State = &State
	}
    
	if Version, ok := GroupcreateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if VarType, ok := GroupcreateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Images, ok := GroupcreateMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Addresses, ok := GroupcreateMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if RulesVisible, ok := GroupcreateMap["rulesVisible"].(bool); ok {
		o.RulesVisible = &RulesVisible
	}
    
	if Visibility, ok := GroupcreateMap["visibility"].(string); ok {
		o.Visibility = &Visibility
	}
    
	if RolesEnabled, ok := GroupcreateMap["rolesEnabled"].(bool); ok {
		o.RolesEnabled = &RolesEnabled
	}
    
	if IncludeOwners, ok := GroupcreateMap["includeOwners"].(bool); ok {
		o.IncludeOwners = &IncludeOwners
	}
    
	if CallsEnabled, ok := GroupcreateMap["callsEnabled"].(bool); ok {
		o.CallsEnabled = &CallsEnabled
	}
    
	if OwnerIds, ok := GroupcreateMap["ownerIds"].([]interface{}); ok {
		OwnerIdsString, _ := json.Marshal(OwnerIds)
		json.Unmarshal(OwnerIdsString, &o.OwnerIds)
	}
	
	if SelfUri, ok := GroupcreateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Groupcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
