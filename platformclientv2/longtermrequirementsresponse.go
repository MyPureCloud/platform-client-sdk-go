package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermrequirementsresponse
type Longtermrequirementsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - Status of the long term forecast
	Status *string `json:"status,omitempty"`

	// ErrorCode - Error code when status is Failed
	ErrorCode *string `json:"errorCode,omitempty"`

	// LongTermRequirements - For schema documentation only, always null, schema for staffing forecast result at downloadUrl
	LongTermRequirements *Longtermrequirements `json:"longTermRequirements,omitempty"`

	// DownloadUrl - Download URL for the staffing forecast result
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Longtermrequirementsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Longtermrequirementsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Longtermrequirementsresponse
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		LongTermRequirements *Longtermrequirements `json:"longTermRequirements,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		ErrorCode: o.ErrorCode,
		
		LongTermRequirements: o.LongTermRequirements,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (Alias)(o),
	})
}

func (o *Longtermrequirementsresponse) UnmarshalJSON(b []byte) error {
	var LongtermrequirementsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermrequirementsresponseMap)
	if err != nil {
		return err
	}
	
	if Status, ok := LongtermrequirementsresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ErrorCode, ok := LongtermrequirementsresponseMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if LongTermRequirements, ok := LongtermrequirementsresponseMap["longTermRequirements"].(map[string]interface{}); ok {
		LongTermRequirementsString, _ := json.Marshal(LongTermRequirements)
		json.Unmarshal(LongTermRequirementsString, &o.LongTermRequirements)
	}
	
	if DownloadUrl, ok := LongtermrequirementsresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermrequirementsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
