package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentflowentityref
type Webdeploymentflowentityref struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The Flow ID
	Id *string `json:"id,omitempty"`

	// Name - The Flow name
	Name *string `json:"name,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// FlowDescription - The flow description for the webdeployment
	FlowDescription *string `json:"flowDescription,omitempty"`

	// PublishVersion - The published config version for the associated deployment
	PublishVersion *Flowversion `json:"publishVersion,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webdeploymentflowentityref) SetField(field string, fieldValue interface{}) {
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

func (o Webdeploymentflowentityref) MarshalJSON() ([]byte, error) {
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
	type Alias Webdeploymentflowentityref
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		FlowDescription *string `json:"flowDescription,omitempty"`
		
		PublishVersion *Flowversion `json:"publishVersion,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		FlowDescription: o.FlowDescription,
		
		PublishVersion: o.PublishVersion,
		Alias:    (Alias)(o),
	})
}

func (o *Webdeploymentflowentityref) UnmarshalJSON(b []byte) error {
	var WebdeploymentflowentityrefMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentflowentityrefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentflowentityrefMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebdeploymentflowentityrefMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := WebdeploymentflowentityrefMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if FlowDescription, ok := WebdeploymentflowentityrefMap["flowDescription"].(string); ok {
		o.FlowDescription = &FlowDescription
	}
    
	if PublishVersion, ok := WebdeploymentflowentityrefMap["publishVersion"].(map[string]interface{}); ok {
		PublishVersionString, _ := json.Marshal(PublishVersion)
		json.Unmarshal(PublishVersionString, &o.PublishVersion)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentflowentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
