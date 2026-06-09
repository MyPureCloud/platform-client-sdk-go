package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontableimportrowmetrics - Progress metrics for a decision table import job
type Decisiontableimportrowmetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TotalRows - Total number of rows in the import file (set after parsing completes)
	TotalRows *int `json:"totalRows,omitempty"`

	// RowsParsed - Number of rows successfully parsed so far
	RowsParsed *int `json:"rowsParsed,omitempty"`

	// RowParseFailed - Number of rows that failed to parse
	RowParseFailed *int `json:"rowParseFailed,omitempty"`

	// RowsCreated - Number of rows successfully created so far
	RowsCreated *int `json:"rowsCreated,omitempty"`

	// RowsUpdated - Number of rows successfully updated so far
	RowsUpdated *int `json:"rowsUpdated,omitempty"`

	// RowsDeleted - Number of rows deleted (Replace mode only)
	RowsDeleted *int `json:"rowsDeleted,omitempty"`

	// RowCreateFailed - Number of rows that failed during batch create
	RowCreateFailed *int `json:"rowCreateFailed,omitempty"`

	// RowUpdateFailed - Number of rows that failed during batch update
	RowUpdateFailed *int `json:"rowUpdateFailed,omitempty"`

	// RowDeleteFailed - Number of rows that failed during delete
	RowDeleteFailed *int `json:"rowDeleteFailed,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontableimportrowmetrics) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontableimportrowmetrics) MarshalJSON() ([]byte, error) {
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
	type Alias Decisiontableimportrowmetrics
	
	return json.Marshal(&struct { 
		TotalRows *int `json:"totalRows,omitempty"`
		
		RowsParsed *int `json:"rowsParsed,omitempty"`
		
		RowParseFailed *int `json:"rowParseFailed,omitempty"`
		
		RowsCreated *int `json:"rowsCreated,omitempty"`
		
		RowsUpdated *int `json:"rowsUpdated,omitempty"`
		
		RowsDeleted *int `json:"rowsDeleted,omitempty"`
		
		RowCreateFailed *int `json:"rowCreateFailed,omitempty"`
		
		RowUpdateFailed *int `json:"rowUpdateFailed,omitempty"`
		
		RowDeleteFailed *int `json:"rowDeleteFailed,omitempty"`
		Alias
	}{ 
		TotalRows: o.TotalRows,
		
		RowsParsed: o.RowsParsed,
		
		RowParseFailed: o.RowParseFailed,
		
		RowsCreated: o.RowsCreated,
		
		RowsUpdated: o.RowsUpdated,
		
		RowsDeleted: o.RowsDeleted,
		
		RowCreateFailed: o.RowCreateFailed,
		
		RowUpdateFailed: o.RowUpdateFailed,
		
		RowDeleteFailed: o.RowDeleteFailed,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontableimportrowmetrics) UnmarshalJSON(b []byte) error {
	var DecisiontableimportrowmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontableimportrowmetricsMap)
	if err != nil {
		return err
	}
	
	if TotalRows, ok := DecisiontableimportrowmetricsMap["totalRows"].(float64); ok {
		TotalRowsInt := int(TotalRows)
		o.TotalRows = &TotalRowsInt
	}
	
	if RowsParsed, ok := DecisiontableimportrowmetricsMap["rowsParsed"].(float64); ok {
		RowsParsedInt := int(RowsParsed)
		o.RowsParsed = &RowsParsedInt
	}
	
	if RowParseFailed, ok := DecisiontableimportrowmetricsMap["rowParseFailed"].(float64); ok {
		RowParseFailedInt := int(RowParseFailed)
		o.RowParseFailed = &RowParseFailedInt
	}
	
	if RowsCreated, ok := DecisiontableimportrowmetricsMap["rowsCreated"].(float64); ok {
		RowsCreatedInt := int(RowsCreated)
		o.RowsCreated = &RowsCreatedInt
	}
	
	if RowsUpdated, ok := DecisiontableimportrowmetricsMap["rowsUpdated"].(float64); ok {
		RowsUpdatedInt := int(RowsUpdated)
		o.RowsUpdated = &RowsUpdatedInt
	}
	
	if RowsDeleted, ok := DecisiontableimportrowmetricsMap["rowsDeleted"].(float64); ok {
		RowsDeletedInt := int(RowsDeleted)
		o.RowsDeleted = &RowsDeletedInt
	}
	
	if RowCreateFailed, ok := DecisiontableimportrowmetricsMap["rowCreateFailed"].(float64); ok {
		RowCreateFailedInt := int(RowCreateFailed)
		o.RowCreateFailed = &RowCreateFailedInt
	}
	
	if RowUpdateFailed, ok := DecisiontableimportrowmetricsMap["rowUpdateFailed"].(float64); ok {
		RowUpdateFailedInt := int(RowUpdateFailed)
		o.RowUpdateFailed = &RowUpdateFailedInt
	}
	
	if RowDeleteFailed, ok := DecisiontableimportrowmetricsMap["rowDeleteFailed"].(float64); ok {
		RowDeleteFailedInt := int(RowDeleteFailed)
		o.RowDeleteFailed = &RowDeleteFailedInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontableimportrowmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
