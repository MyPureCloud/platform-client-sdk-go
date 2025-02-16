package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Authorizationpolicy
type Authorizationpolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// TargetResource - The targeted resource to which the policy should apply, in the form of domain:entity:action
	TargetResource *string `json:"targetResource,omitempty"`

	// Subject - The subject to whom the policy will apply, including type and id
	Subject *Subject `json:"subject,omitempty"`

	// Effect - The effect this policy should have when its conditions are met
	Effect *string `json:"effect,omitempty"`

	// Condition - The condition tree the policy will evaluate
	Condition *interface{} `json:"condition,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// DateModified - Date this policy was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// PresetAttributes - Map of names and values of preset attributes to use in policy evaluation
	PresetAttributes *map[string]Typedattribute `json:"presetAttributes,omitempty"`

	// Active - Flag for active enforcement. If this value is false or null, the policy will be saved but will not be checked or enforced on users.
	Active *bool `json:"active,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Authorizationpolicy) SetField(field string, fieldValue interface{}) {
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

func (o Authorizationpolicy) MarshalJSON() ([]byte, error) {
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
	type Alias Authorizationpolicy
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TargetResource *string `json:"targetResource,omitempty"`
		
		Subject *Subject `json:"subject,omitempty"`
		
		Effect *string `json:"effect,omitempty"`
		
		Condition *interface{} `json:"condition,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		PresetAttributes *map[string]Typedattribute `json:"presetAttributes,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		TargetResource: o.TargetResource,
		
		Subject: o.Subject,
		
		Effect: o.Effect,
		
		Condition: o.Condition,
		
		Description: o.Description,
		
		DateModified: DateModified,
		
		PresetAttributes: o.PresetAttributes,
		
		Active: o.Active,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Authorizationpolicy) UnmarshalJSON(b []byte) error {
	var AuthorizationpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &AuthorizationpolicyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuthorizationpolicyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AuthorizationpolicyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TargetResource, ok := AuthorizationpolicyMap["targetResource"].(string); ok {
		o.TargetResource = &TargetResource
	}
    
	if Subject, ok := AuthorizationpolicyMap["subject"].(map[string]interface{}); ok {
		SubjectString, _ := json.Marshal(Subject)
		json.Unmarshal(SubjectString, &o.Subject)
	}
	
	if Effect, ok := AuthorizationpolicyMap["effect"].(string); ok {
		o.Effect = &Effect
	}
    
	if Condition, ok := AuthorizationpolicyMap["condition"].(map[string]interface{}); ok {
		ConditionString, _ := json.Marshal(Condition)
		json.Unmarshal(ConditionString, &o.Condition)
	}
	
	if Description, ok := AuthorizationpolicyMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateModifiedString, ok := AuthorizationpolicyMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if PresetAttributes, ok := AuthorizationpolicyMap["presetAttributes"].(map[string]interface{}); ok {
		PresetAttributesString, _ := json.Marshal(PresetAttributes)
		json.Unmarshal(PresetAttributesString, &o.PresetAttributes)
	}
	
	if Active, ok := AuthorizationpolicyMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if SelfUri, ok := AuthorizationpolicyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Authorizationpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
