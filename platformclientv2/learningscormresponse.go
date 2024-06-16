package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningscormresponse - Learning SCORM package
type Learningscormresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Status - The status of the SCORM package
	Status *string `json:"status,omitempty"`

	// ErrorCode - The error code of the SCORM package (on failure)
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorMessage - The error message associated with the error code
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// PercentageUnpacked - The percentage of the SCORM package that has been unpacked
	PercentageUnpacked *float32 `json:"percentageUnpacked,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningscormresponse) SetField(field string, fieldValue interface{}) {
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

func (o Learningscormresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Learningscormresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		PercentageUnpacked *float32 `json:"percentageUnpacked,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		
		PercentageUnpacked: o.PercentageUnpacked,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Learningscormresponse) UnmarshalJSON(b []byte) error {
	var LearningscormresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningscormresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningscormresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := LearningscormresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ErrorCode, ok := LearningscormresponseMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := LearningscormresponseMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if PercentageUnpacked, ok := LearningscormresponseMap["percentageUnpacked"].(float64); ok {
		PercentageUnpackedFloat32 := float32(PercentageUnpacked)
		o.PercentageUnpacked = &PercentageUnpackedFloat32
	}
	
	if SelfUri, ok := LearningscormresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningscormresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
