package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformpolicy
type Crossplatformpolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// Order
	Order *int `json:"order,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

	// MediaPolicies - Conditions and actions per media type
	MediaPolicies *Crossplatformmediapolicies `json:"mediaPolicies,omitempty"`

	// Conditions - Conditions
	Conditions *Policyconditions `json:"conditions,omitempty"`

	// Actions - Actions
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`

	// PolicyErrors
	PolicyErrors *Policyerrors `json:"policyErrors,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Crossplatformpolicy) SetField(field string, fieldValue interface{}) {
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

func (o Crossplatformpolicy) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifiedDate","CreatedDate", }
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
	type Alias Crossplatformpolicy
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MediaPolicies *Crossplatformmediapolicies `json:"mediaPolicies,omitempty"`
		
		Conditions *Policyconditions `json:"conditions,omitempty"`
		
		Actions *Crossplatformpolicyactions `json:"actions,omitempty"`
		
		PolicyErrors *Policyerrors `json:"policyErrors,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ModifiedDate: ModifiedDate,
		
		CreatedDate: CreatedDate,
		
		Order: o.Order,
		
		Description: o.Description,
		
		Enabled: o.Enabled,
		
		MediaPolicies: o.MediaPolicies,
		
		Conditions: o.Conditions,
		
		Actions: o.Actions,
		
		PolicyErrors: o.PolicyErrors,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Crossplatformpolicy) UnmarshalJSON(b []byte) error {
	var CrossplatformpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &CrossplatformpolicyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CrossplatformpolicyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CrossplatformpolicyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if modifiedDateString, ok := CrossplatformpolicyMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if createdDateString, ok := CrossplatformpolicyMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if Order, ok := CrossplatformpolicyMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if Description, ok := CrossplatformpolicyMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Enabled, ok := CrossplatformpolicyMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MediaPolicies, ok := CrossplatformpolicyMap["mediaPolicies"].(map[string]interface{}); ok {
		MediaPoliciesString, _ := json.Marshal(MediaPolicies)
		json.Unmarshal(MediaPoliciesString, &o.MediaPolicies)
	}
	
	if Conditions, ok := CrossplatformpolicyMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Actions, ok := CrossplatformpolicyMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if PolicyErrors, ok := CrossplatformpolicyMap["policyErrors"].(map[string]interface{}); ok {
		PolicyErrorsString, _ := json.Marshal(PolicyErrors)
		json.Unmarshal(PolicyErrorsString, &o.PolicyErrors)
	}
	
	if SelfUri, ok := CrossplatformpolicyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Crossplatformpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
