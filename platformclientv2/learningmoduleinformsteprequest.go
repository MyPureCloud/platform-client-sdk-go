package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmoduleinformsteprequest - Learning module inform steps request
type Learningmoduleinformsteprequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The learning module inform step type
	VarType *string `json:"type,omitempty"`

	// Name - The name of the inform step or content
	Name *string `json:"name,omitempty"`

	// Value - The value for inform step
	Value *string `json:"value,omitempty"`

	// SharingUri - The sharing uri for Content type inform step
	SharingUri *string `json:"sharingUri,omitempty"`

	// ContentType - The document type for Content type Inform step
	ContentType *string `json:"contentType,omitempty"`

	// Order - The order of inform step in a learning module
	Order *int `json:"order,omitempty"`

	// DisplayName - The display name for the inform step
	DisplayName *string `json:"displayName,omitempty"`

	// Description - The description for the inform step
	Description *string `json:"description,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmoduleinformsteprequest) SetField(field string, fieldValue interface{}) {
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

func (o Learningmoduleinformsteprequest) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmoduleinformsteprequest
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		SharingUri *string `json:"sharingUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Description *string `json:"description,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Name: o.Name,
		
		Value: o.Value,
		
		SharingUri: o.SharingUri,
		
		ContentType: o.ContentType,
		
		Order: o.Order,
		
		DisplayName: o.DisplayName,
		
		Description: o.Description,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmoduleinformsteprequest) UnmarshalJSON(b []byte) error {
	var LearningmoduleinformsteprequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmoduleinformsteprequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := LearningmoduleinformsteprequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Name, ok := LearningmoduleinformsteprequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Value, ok := LearningmoduleinformsteprequestMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if SharingUri, ok := LearningmoduleinformsteprequestMap["sharingUri"].(string); ok {
		o.SharingUri = &SharingUri
	}
    
	if ContentType, ok := LearningmoduleinformsteprequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Order, ok := LearningmoduleinformsteprequestMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if DisplayName, ok := LearningmoduleinformsteprequestMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Description, ok := LearningmoduleinformsteprequestMap["description"].(string); ok {
		o.Description = &Description
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmoduleinformsteprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
