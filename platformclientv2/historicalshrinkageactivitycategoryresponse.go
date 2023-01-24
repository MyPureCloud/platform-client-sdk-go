package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageactivitycategoryresponse
type Historicalshrinkageactivitycategoryresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityCategory - Activity category for which shrinkage data is provided
	ActivityCategory *string `json:"activityCategory,omitempty"`

	// ShrinkageForActivityCategory - Aggregated shrinkage data for the activity category
	ShrinkageForActivityCategory *Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCategory,omitempty"`

	// ShrinkageForActivityCodes - Shrinkage for the activity codes under this activity category
	ShrinkageForActivityCodes *[]Historicalshrinkageactivitycoderesponse `json:"shrinkageForActivityCodes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historicalshrinkageactivitycategoryresponse) SetField(field string, fieldValue interface{}) {
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

func (o Historicalshrinkageactivitycategoryresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Historicalshrinkageactivitycategoryresponse
	
	return json.Marshal(&struct { 
		ActivityCategory *string `json:"activityCategory,omitempty"`
		
		ShrinkageForActivityCategory *Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCategory,omitempty"`
		
		ShrinkageForActivityCodes *[]Historicalshrinkageactivitycoderesponse `json:"shrinkageForActivityCodes,omitempty"`
		Alias
	}{ 
		ActivityCategory: o.ActivityCategory,
		
		ShrinkageForActivityCategory: o.ShrinkageForActivityCategory,
		
		ShrinkageForActivityCodes: o.ShrinkageForActivityCodes,
		Alias:    (Alias)(o),
	})
}

func (o *Historicalshrinkageactivitycategoryresponse) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageactivitycategoryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageactivitycategoryresponseMap)
	if err != nil {
		return err
	}
	
	if ActivityCategory, ok := HistoricalshrinkageactivitycategoryresponseMap["activityCategory"].(string); ok {
		o.ActivityCategory = &ActivityCategory
	}
    
	if ShrinkageForActivityCategory, ok := HistoricalshrinkageactivitycategoryresponseMap["shrinkageForActivityCategory"].(map[string]interface{}); ok {
		ShrinkageForActivityCategoryString, _ := json.Marshal(ShrinkageForActivityCategory)
		json.Unmarshal(ShrinkageForActivityCategoryString, &o.ShrinkageForActivityCategory)
	}
	
	if ShrinkageForActivityCodes, ok := HistoricalshrinkageactivitycategoryresponseMap["shrinkageForActivityCodes"].([]interface{}); ok {
		ShrinkageForActivityCodesString, _ := json.Marshal(ShrinkageForActivityCodes)
		json.Unmarshal(ShrinkageForActivityCodesString, &o.ShrinkageForActivityCodes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageactivitycategoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
