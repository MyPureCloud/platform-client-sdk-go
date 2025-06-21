package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontableexecutionresponse
type Decisiontableexecutionresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Table - The decision table version entity that was executed.
	Table *Decisiontableversionentity `json:"table,omitempty"`

	// TotalMatchRowCount - Total number of rows that matched execution input and would return results
	TotalMatchRowCount *int `json:"totalMatchRowCount,omitempty"`

	// TopMatchRows - Top 5 rows matching execution input, excluding the one produced the result.
	TopMatchRows *[]Decisiontablerowentityref `json:"topMatchRows,omitempty"`

	// RowExecutionOutputs - The output data for each executed row for which output is collected.
	RowExecutionOutputs *[]Decisiontablerowexecutionoutput `json:"rowExecutionOutputs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontableexecutionresponse) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontableexecutionresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Decisiontableexecutionresponse
	
	return json.Marshal(&struct { 
		Table *Decisiontableversionentity `json:"table,omitempty"`
		
		TotalMatchRowCount *int `json:"totalMatchRowCount,omitempty"`
		
		TopMatchRows *[]Decisiontablerowentityref `json:"topMatchRows,omitempty"`
		
		RowExecutionOutputs *[]Decisiontablerowexecutionoutput `json:"rowExecutionOutputs,omitempty"`
		Alias
	}{ 
		Table: o.Table,
		
		TotalMatchRowCount: o.TotalMatchRowCount,
		
		TopMatchRows: o.TopMatchRows,
		
		RowExecutionOutputs: o.RowExecutionOutputs,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontableexecutionresponse) UnmarshalJSON(b []byte) error {
	var DecisiontableexecutionresponseMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontableexecutionresponseMap)
	if err != nil {
		return err
	}
	
	if Table, ok := DecisiontableexecutionresponseMap["table"].(map[string]interface{}); ok {
		TableString, _ := json.Marshal(Table)
		json.Unmarshal(TableString, &o.Table)
	}
	
	if TotalMatchRowCount, ok := DecisiontableexecutionresponseMap["totalMatchRowCount"].(float64); ok {
		TotalMatchRowCountInt := int(TotalMatchRowCount)
		o.TotalMatchRowCount = &TotalMatchRowCountInt
	}
	
	if TopMatchRows, ok := DecisiontableexecutionresponseMap["topMatchRows"].([]interface{}); ok {
		TopMatchRowsString, _ := json.Marshal(TopMatchRows)
		json.Unmarshal(TopMatchRowsString, &o.TopMatchRows)
	}
	
	if RowExecutionOutputs, ok := DecisiontableexecutionresponseMap["rowExecutionOutputs"].([]interface{}); ok {
		RowExecutionOutputsString, _ := json.Marshal(RowExecutionOutputs)
		json.Unmarshal(RowExecutionOutputsString, &o.RowExecutionOutputs)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontableexecutionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
