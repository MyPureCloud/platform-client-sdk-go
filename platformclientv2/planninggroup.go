package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Planninggroup
type Planninggroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ServiceGoalTemplate - The ID of the service goal template associated with this planning group
	ServiceGoalTemplate *Servicegoaltemplatereference `json:"serviceGoalTemplate,omitempty"`

	// RoutePaths - Set of route paths associated with the planning group
	RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`

	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Planninggroup) SetField(field string, fieldValue interface{}) {
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

func (o Planninggroup) MarshalJSON() ([]byte, error) {
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
	type Alias Planninggroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ServiceGoalTemplate *Servicegoaltemplatereference `json:"serviceGoalTemplate,omitempty"`
		
		RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ServiceGoalTemplate: o.ServiceGoalTemplate,
		
		RoutePaths: o.RoutePaths,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Planninggroup) UnmarshalJSON(b []byte) error {
	var PlanninggroupMap map[string]interface{}
	err := json.Unmarshal(b, &PlanninggroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PlanninggroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := PlanninggroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ServiceGoalTemplate, ok := PlanninggroupMap["serviceGoalTemplate"].(map[string]interface{}); ok {
		ServiceGoalTemplateString, _ := json.Marshal(ServiceGoalTemplate)
		json.Unmarshal(ServiceGoalTemplateString, &o.ServiceGoalTemplate)
	}
	
	if RoutePaths, ok := PlanninggroupMap["routePaths"].([]interface{}); ok {
		RoutePathsString, _ := json.Marshal(RoutePaths)
		json.Unmarshal(RoutePathsString, &o.RoutePaths)
	}
	
	if Metadata, ok := PlanninggroupMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := PlanninggroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Planninggroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
