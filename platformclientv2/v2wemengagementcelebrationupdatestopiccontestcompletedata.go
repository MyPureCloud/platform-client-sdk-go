package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2wemengagementcelebrationupdatestopiccontestcompletedata
type V2wemengagementcelebrationupdatestopiccontestcompletedata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateEnd
	DateEnd *string `json:"dateEnd,omitempty"`

	// Anonymization
	Anonymization *string `json:"anonymization,omitempty"`

	// Metrics
	Metrics *[]V2wemengagementcelebrationupdatestopiccontestmetrics `json:"metrics,omitempty"`

	// Prizes
	Prizes *[]V2wemengagementcelebrationupdatestopiccontestprizes `json:"prizes,omitempty"`

	// Winners
	Winners *[]V2wemengagementcelebrationupdatestopiccontestwinners `json:"winners,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2wemengagementcelebrationupdatestopiccontestcompletedata) SetField(field string, fieldValue interface{}) {
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

func (o V2wemengagementcelebrationupdatestopiccontestcompletedata) MarshalJSON() ([]byte, error) {
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
	type Alias V2wemengagementcelebrationupdatestopiccontestcompletedata
	
	return json.Marshal(&struct { 
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Anonymization *string `json:"anonymization,omitempty"`
		
		Metrics *[]V2wemengagementcelebrationupdatestopiccontestmetrics `json:"metrics,omitempty"`
		
		Prizes *[]V2wemengagementcelebrationupdatestopiccontestprizes `json:"prizes,omitempty"`
		
		Winners *[]V2wemengagementcelebrationupdatestopiccontestwinners `json:"winners,omitempty"`
		Alias
	}{ 
		DateEnd: o.DateEnd,
		
		Anonymization: o.Anonymization,
		
		Metrics: o.Metrics,
		
		Prizes: o.Prizes,
		
		Winners: o.Winners,
		Alias:    (Alias)(o),
	})
}

func (o *V2wemengagementcelebrationupdatestopiccontestcompletedata) UnmarshalJSON(b []byte) error {
	var V2wemengagementcelebrationupdatestopiccontestcompletedataMap map[string]interface{}
	err := json.Unmarshal(b, &V2wemengagementcelebrationupdatestopiccontestcompletedataMap)
	if err != nil {
		return err
	}
	
	if DateEnd, ok := V2wemengagementcelebrationupdatestopiccontestcompletedataMap["dateEnd"].(string); ok {
		o.DateEnd = &DateEnd
	}
    
	if Anonymization, ok := V2wemengagementcelebrationupdatestopiccontestcompletedataMap["anonymization"].(string); ok {
		o.Anonymization = &Anonymization
	}
    
	if Metrics, ok := V2wemengagementcelebrationupdatestopiccontestcompletedataMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if Prizes, ok := V2wemengagementcelebrationupdatestopiccontestcompletedataMap["prizes"].([]interface{}); ok {
		PrizesString, _ := json.Marshal(Prizes)
		json.Unmarshal(PrizesString, &o.Prizes)
	}
	
	if Winners, ok := V2wemengagementcelebrationupdatestopiccontestcompletedataMap["winners"].([]interface{}); ok {
		WinnersString, _ := json.Marshal(Winners)
		json.Unmarshal(WinnersString, &o.Winners)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2wemengagementcelebrationupdatestopiccontestcompletedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
