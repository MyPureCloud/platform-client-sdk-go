package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Share
type Share struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// SharedEntityType
	SharedEntityType *string `json:"sharedEntityType,omitempty"`

	// SharedEntity
	SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`

	// MemberType
	MemberType *string `json:"memberType,omitempty"`

	// Member
	Member *Domainentityref `json:"member,omitempty"`

	// SharedBy
	SharedBy *Domainentityref `json:"sharedBy,omitempty"`

	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`

	// User
	User *User `json:"user,omitempty"`

	// Group
	Group *Group `json:"group,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Share) SetField(field string, fieldValue interface{}) {
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

func (o Share) MarshalJSON() ([]byte, error) {
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
	type Alias Share
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SharedEntityType *string `json:"sharedEntityType,omitempty"`
		
		SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Domainentityref `json:"member,omitempty"`
		
		SharedBy *Domainentityref `json:"sharedBy,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SharedEntityType: o.SharedEntityType,
		
		SharedEntity: o.SharedEntity,
		
		MemberType: o.MemberType,
		
		Member: o.Member,
		
		SharedBy: o.SharedBy,
		
		Workspace: o.Workspace,
		
		User: o.User,
		
		Group: o.Group,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Share) UnmarshalJSON(b []byte) error {
	var ShareMap map[string]interface{}
	err := json.Unmarshal(b, &ShareMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ShareMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ShareMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SharedEntityType, ok := ShareMap["sharedEntityType"].(string); ok {
		o.SharedEntityType = &SharedEntityType
	}
    
	if SharedEntity, ok := ShareMap["sharedEntity"].(map[string]interface{}); ok {
		SharedEntityString, _ := json.Marshal(SharedEntity)
		json.Unmarshal(SharedEntityString, &o.SharedEntity)
	}
	
	if MemberType, ok := ShareMap["memberType"].(string); ok {
		o.MemberType = &MemberType
	}
    
	if Member, ok := ShareMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	
	if SharedBy, ok := ShareMap["sharedBy"].(map[string]interface{}); ok {
		SharedByString, _ := json.Marshal(SharedBy)
		json.Unmarshal(SharedByString, &o.SharedBy)
	}
	
	if Workspace, ok := ShareMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if User, ok := ShareMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Group, ok := ShareMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if SelfUri, ok := ShareMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Share) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
