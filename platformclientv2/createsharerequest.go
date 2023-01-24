package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createsharerequest
type Createsharerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SharedEntityType - The share entity type
	SharedEntityType *string `json:"sharedEntityType,omitempty"`

	// SharedEntity - The entity that will be shared
	SharedEntity *Sharedentity `json:"sharedEntity,omitempty"`

	// MemberType
	MemberType *string `json:"memberType,omitempty"`

	// Member - The member that will have access to this share. Only required if a list of members is not provided.
	Member *Sharedentity `json:"member,omitempty"`

	// Members
	Members *[]Createsharerequestmember `json:"members,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createsharerequest) SetField(field string, fieldValue interface{}) {
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

func (o Createsharerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createsharerequest
	
	return json.Marshal(&struct { 
		SharedEntityType *string `json:"sharedEntityType,omitempty"`
		
		SharedEntity *Sharedentity `json:"sharedEntity,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Sharedentity `json:"member,omitempty"`
		
		Members *[]Createsharerequestmember `json:"members,omitempty"`
		Alias
	}{ 
		SharedEntityType: o.SharedEntityType,
		
		SharedEntity: o.SharedEntity,
		
		MemberType: o.MemberType,
		
		Member: o.Member,
		
		Members: o.Members,
		Alias:    (Alias)(o),
	})
}

func (o *Createsharerequest) UnmarshalJSON(b []byte) error {
	var CreatesharerequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatesharerequestMap)
	if err != nil {
		return err
	}
	
	if SharedEntityType, ok := CreatesharerequestMap["sharedEntityType"].(string); ok {
		o.SharedEntityType = &SharedEntityType
	}
    
	if SharedEntity, ok := CreatesharerequestMap["sharedEntity"].(map[string]interface{}); ok {
		SharedEntityString, _ := json.Marshal(SharedEntity)
		json.Unmarshal(SharedEntityString, &o.SharedEntity)
	}
	
	if MemberType, ok := CreatesharerequestMap["memberType"].(string); ok {
		o.MemberType = &MemberType
	}
    
	if Member, ok := CreatesharerequestMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	
	if Members, ok := CreatesharerequestMap["members"].([]interface{}); ok {
		MembersString, _ := json.Marshal(Members)
		json.Unmarshal(MembersString, &o.Members)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createsharerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
