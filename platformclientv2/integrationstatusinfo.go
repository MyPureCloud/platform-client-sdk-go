package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationstatusinfo - Status information for an Integration.
type Integrationstatusinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Code - Machine-readable status as reported by the integration.
	Code *string `json:"code,omitempty"`

	// Effective - Localized, human-readable, effective status of the integration.
	Effective *string `json:"effective,omitempty"`

	// Detail - Localizable status details for the integration.
	Detail *Messageinfo `json:"detail,omitempty"`

	// LastUpdated - Date and time (in UTC) when the integration status (i.e. the code field) was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Integrationstatusinfo) SetField(field string, fieldValue interface{}) {
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

func (o Integrationstatusinfo) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "LastUpdated", }
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
	type Alias Integrationstatusinfo
	
	LastUpdated := new(string)
	if o.LastUpdated != nil {
		
		*LastUpdated = timeutil.Strftime(o.LastUpdated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastUpdated = nil
	}
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Effective *string `json:"effective,omitempty"`
		
		Detail *Messageinfo `json:"detail,omitempty"`
		
		LastUpdated *string `json:"lastUpdated,omitempty"`
		Alias
	}{ 
		Code: o.Code,
		
		Effective: o.Effective,
		
		Detail: o.Detail,
		
		LastUpdated: LastUpdated,
		Alias:    (Alias)(o),
	})
}

func (o *Integrationstatusinfo) UnmarshalJSON(b []byte) error {
	var IntegrationstatusinfoMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationstatusinfoMap)
	if err != nil {
		return err
	}
	
	if Code, ok := IntegrationstatusinfoMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Effective, ok := IntegrationstatusinfoMap["effective"].(string); ok {
		o.Effective = &Effective
	}
    
	if Detail, ok := IntegrationstatusinfoMap["detail"].(map[string]interface{}); ok {
		DetailString, _ := json.Marshal(Detail)
		json.Unmarshal(DetailString, &o.Detail)
	}
	
	if lastUpdatedString, ok := IntegrationstatusinfoMap["lastUpdated"].(string); ok {
		LastUpdated, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastUpdatedString)
		o.LastUpdated = &LastUpdated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationstatusinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
