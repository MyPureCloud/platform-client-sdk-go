package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createplanninggrouprequest
type Createplanninggrouprequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the planning group
	Name *string `json:"name,omitempty"`

	// RoutePaths - Set of route paths to associate with the planning group
	RoutePaths *[]Routepathrequest `json:"routePaths,omitempty"`

	// ServiceGoalTemplateId - The ID of the service goal template to associate with this planning group
	ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createplanninggrouprequest) SetField(field string, fieldValue interface{}) {
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

func (o Createplanninggrouprequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createplanninggrouprequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		RoutePaths *[]Routepathrequest `json:"routePaths,omitempty"`
		
		ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		RoutePaths: o.RoutePaths,
		
		ServiceGoalTemplateId: o.ServiceGoalTemplateId,
		Alias:    (Alias)(o),
	})
}

func (o *Createplanninggrouprequest) UnmarshalJSON(b []byte) error {
	var CreateplanninggrouprequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateplanninggrouprequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateplanninggrouprequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if RoutePaths, ok := CreateplanninggrouprequestMap["routePaths"].([]interface{}); ok {
		RoutePathsString, _ := json.Marshal(RoutePaths)
		json.Unmarshal(RoutePathsString, &o.RoutePaths)
	}
	
	if ServiceGoalTemplateId, ok := CreateplanninggrouprequestMap["serviceGoalTemplateId"].(string); ok {
		o.ServiceGoalTemplateId = &ServiceGoalTemplateId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
