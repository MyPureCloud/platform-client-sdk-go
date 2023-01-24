package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityauditqueryexecutionresultsresponse
type Qualityauditqueryexecutionresultsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Id of the audit query execution request.
	Id *string `json:"id,omitempty"`

	// PageSize - Number of results in a page.
	PageSize *int `json:"pageSize,omitempty"`

	// Cursor - Optional cursor to indicate where to resume the results.
	Cursor *string `json:"cursor,omitempty"`

	// Entities - List of audit messages.
	Entities *[]Qualityauditlogmessage `json:"entities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Qualityauditqueryexecutionresultsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Qualityauditqueryexecutionresultsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Qualityauditqueryexecutionresultsresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Cursor *string `json:"cursor,omitempty"`
		
		Entities *[]Qualityauditlogmessage `json:"entities,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		PageSize: o.PageSize,
		
		Cursor: o.Cursor,
		
		Entities: o.Entities,
		Alias:    (Alias)(o),
	})
}

func (o *Qualityauditqueryexecutionresultsresponse) UnmarshalJSON(b []byte) error {
	var QualityauditqueryexecutionresultsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &QualityauditqueryexecutionresultsresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QualityauditqueryexecutionresultsresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PageSize, ok := QualityauditqueryexecutionresultsresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Cursor, ok := QualityauditqueryexecutionresultsresponseMap["cursor"].(string); ok {
		o.Cursor = &Cursor
	}
    
	if Entities, ok := QualityauditqueryexecutionresultsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Qualityauditqueryexecutionresultsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
