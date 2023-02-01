package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workspacemember
type Workspacemember struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`

	// MemberType - The workspace member type.
	MemberType *string `json:"memberType,omitempty"`

	// Member
	Member *Domainentityref `json:"member,omitempty"`

	// User
	User *User `json:"user,omitempty"`

	// Group
	Group *Group `json:"group,omitempty"`

	// SecurityProfile
	SecurityProfile *Securityprofile `json:"securityProfile,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workspacemember) SetField(field string, fieldValue interface{}) {
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

func (o Workspacemember) MarshalJSON() ([]byte, error) {
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
	type Alias Workspacemember
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Domainentityref `json:"member,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		SecurityProfile *Securityprofile `json:"securityProfile,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Workspace: o.Workspace,
		
		MemberType: o.MemberType,
		
		Member: o.Member,
		
		User: o.User,
		
		Group: o.Group,
		
		SecurityProfile: o.SecurityProfile,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workspacemember) UnmarshalJSON(b []byte) error {
	var WorkspacememberMap map[string]interface{}
	err := json.Unmarshal(b, &WorkspacememberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkspacememberMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkspacememberMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Workspace, ok := WorkspacememberMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if MemberType, ok := WorkspacememberMap["memberType"].(string); ok {
		o.MemberType = &MemberType
	}
    
	if Member, ok := WorkspacememberMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	
	if User, ok := WorkspacememberMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Group, ok := WorkspacememberMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if SecurityProfile, ok := WorkspacememberMap["securityProfile"].(map[string]interface{}); ok {
		SecurityProfileString, _ := json.Marshal(SecurityProfile)
		json.Unmarshal(SecurityProfileString, &o.SecurityProfile)
	}
	
	if SelfUri, ok := WorkspacememberMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workspacemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
