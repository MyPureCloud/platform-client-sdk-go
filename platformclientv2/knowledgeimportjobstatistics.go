package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobstatistics
type Knowledgeimportjobstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CountDocumentImportActivityCreate - Number of documents will be created by the import.
	CountDocumentImportActivityCreate *int `json:"countDocumentImportActivityCreate,omitempty"`

	// CountDocumentImportActivityUpdate - Number of documents will be updated by the import.
	CountDocumentImportActivityUpdate *int `json:"countDocumentImportActivityUpdate,omitempty"`

	// CountDocumentStateDraft - Number of documents will be imported as draft.
	CountDocumentStateDraft *int `json:"countDocumentStateDraft,omitempty"`

	// CountDocumentStatePublished - Number of documents will be imported as published.
	CountDocumentStatePublished *int `json:"countDocumentStatePublished,omitempty"`

	// CountDocumentValidationSuccess - Number of documents that validated successfully for import.
	CountDocumentValidationSuccess *int `json:"countDocumentValidationSuccess,omitempty"`

	// CountDocumentValidationFailure - Number of documents that failed validation for import.
	CountDocumentValidationFailure *int `json:"countDocumentValidationFailure,omitempty"`

	// CountDocumentImportSuccess - Number of imported documents.
	CountDocumentImportSuccess *int `json:"countDocumentImportSuccess,omitempty"`

	// CountDocumentImportFailure - Number of documents failed to import.
	CountDocumentImportFailure *int `json:"countDocumentImportFailure,omitempty"`

	// CountCategoryValidationSuccess - Number of categories that validated successfully for import.
	CountCategoryValidationSuccess *int `json:"countCategoryValidationSuccess,omitempty"`

	// CountCategoryValidationFailure - Number of categories that failed validation for import.
	CountCategoryValidationFailure *int `json:"countCategoryValidationFailure,omitempty"`

	// CountCategoryImportSuccess - Number of imported categories.
	CountCategoryImportSuccess *int `json:"countCategoryImportSuccess,omitempty"`

	// CountCategoryImportFailure - Number of categories failed to import.
	CountCategoryImportFailure *int `json:"countCategoryImportFailure,omitempty"`

	// CountLabelValidationSuccess - Number of labels that validated successfully for import.
	CountLabelValidationSuccess *int `json:"countLabelValidationSuccess,omitempty"`

	// CountLabelValidationFailure - Number of labels that failed validation for import.
	CountLabelValidationFailure *int `json:"countLabelValidationFailure,omitempty"`

	// CountLabelImportSuccess - Number of imported labels.
	CountLabelImportSuccess *int `json:"countLabelImportSuccess,omitempty"`

	// CountLabelImportFailure - Number of labels failed to import.
	CountLabelImportFailure *int `json:"countLabelImportFailure,omitempty"`

	// MigrationDetected - Shows whether the import treated as migration or not.
	MigrationDetected *bool `json:"migrationDetected,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeimportjobstatistics) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeimportjobstatistics) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgeimportjobstatistics
	
	return json.Marshal(&struct { 
		CountDocumentImportActivityCreate *int `json:"countDocumentImportActivityCreate,omitempty"`
		
		CountDocumentImportActivityUpdate *int `json:"countDocumentImportActivityUpdate,omitempty"`
		
		CountDocumentStateDraft *int `json:"countDocumentStateDraft,omitempty"`
		
		CountDocumentStatePublished *int `json:"countDocumentStatePublished,omitempty"`
		
		CountDocumentValidationSuccess *int `json:"countDocumentValidationSuccess,omitempty"`
		
		CountDocumentValidationFailure *int `json:"countDocumentValidationFailure,omitempty"`
		
		CountDocumentImportSuccess *int `json:"countDocumentImportSuccess,omitempty"`
		
		CountDocumentImportFailure *int `json:"countDocumentImportFailure,omitempty"`
		
		CountCategoryValidationSuccess *int `json:"countCategoryValidationSuccess,omitempty"`
		
		CountCategoryValidationFailure *int `json:"countCategoryValidationFailure,omitempty"`
		
		CountCategoryImportSuccess *int `json:"countCategoryImportSuccess,omitempty"`
		
		CountCategoryImportFailure *int `json:"countCategoryImportFailure,omitempty"`
		
		CountLabelValidationSuccess *int `json:"countLabelValidationSuccess,omitempty"`
		
		CountLabelValidationFailure *int `json:"countLabelValidationFailure,omitempty"`
		
		CountLabelImportSuccess *int `json:"countLabelImportSuccess,omitempty"`
		
		CountLabelImportFailure *int `json:"countLabelImportFailure,omitempty"`
		
		MigrationDetected *bool `json:"migrationDetected,omitempty"`
		Alias
	}{ 
		CountDocumentImportActivityCreate: o.CountDocumentImportActivityCreate,
		
		CountDocumentImportActivityUpdate: o.CountDocumentImportActivityUpdate,
		
		CountDocumentStateDraft: o.CountDocumentStateDraft,
		
		CountDocumentStatePublished: o.CountDocumentStatePublished,
		
		CountDocumentValidationSuccess: o.CountDocumentValidationSuccess,
		
		CountDocumentValidationFailure: o.CountDocumentValidationFailure,
		
		CountDocumentImportSuccess: o.CountDocumentImportSuccess,
		
		CountDocumentImportFailure: o.CountDocumentImportFailure,
		
		CountCategoryValidationSuccess: o.CountCategoryValidationSuccess,
		
		CountCategoryValidationFailure: o.CountCategoryValidationFailure,
		
		CountCategoryImportSuccess: o.CountCategoryImportSuccess,
		
		CountCategoryImportFailure: o.CountCategoryImportFailure,
		
		CountLabelValidationSuccess: o.CountLabelValidationSuccess,
		
		CountLabelValidationFailure: o.CountLabelValidationFailure,
		
		CountLabelImportSuccess: o.CountLabelImportSuccess,
		
		CountLabelImportFailure: o.CountLabelImportFailure,
		
		MigrationDetected: o.MigrationDetected,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeimportjobstatistics) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobstatisticsMap)
	if err != nil {
		return err
	}
	
	if CountDocumentImportActivityCreate, ok := KnowledgeimportjobstatisticsMap["countDocumentImportActivityCreate"].(float64); ok {
		CountDocumentImportActivityCreateInt := int(CountDocumentImportActivityCreate)
		o.CountDocumentImportActivityCreate = &CountDocumentImportActivityCreateInt
	}
	
	if CountDocumentImportActivityUpdate, ok := KnowledgeimportjobstatisticsMap["countDocumentImportActivityUpdate"].(float64); ok {
		CountDocumentImportActivityUpdateInt := int(CountDocumentImportActivityUpdate)
		o.CountDocumentImportActivityUpdate = &CountDocumentImportActivityUpdateInt
	}
	
	if CountDocumentStateDraft, ok := KnowledgeimportjobstatisticsMap["countDocumentStateDraft"].(float64); ok {
		CountDocumentStateDraftInt := int(CountDocumentStateDraft)
		o.CountDocumentStateDraft = &CountDocumentStateDraftInt
	}
	
	if CountDocumentStatePublished, ok := KnowledgeimportjobstatisticsMap["countDocumentStatePublished"].(float64); ok {
		CountDocumentStatePublishedInt := int(CountDocumentStatePublished)
		o.CountDocumentStatePublished = &CountDocumentStatePublishedInt
	}
	
	if CountDocumentValidationSuccess, ok := KnowledgeimportjobstatisticsMap["countDocumentValidationSuccess"].(float64); ok {
		CountDocumentValidationSuccessInt := int(CountDocumentValidationSuccess)
		o.CountDocumentValidationSuccess = &CountDocumentValidationSuccessInt
	}
	
	if CountDocumentValidationFailure, ok := KnowledgeimportjobstatisticsMap["countDocumentValidationFailure"].(float64); ok {
		CountDocumentValidationFailureInt := int(CountDocumentValidationFailure)
		o.CountDocumentValidationFailure = &CountDocumentValidationFailureInt
	}
	
	if CountDocumentImportSuccess, ok := KnowledgeimportjobstatisticsMap["countDocumentImportSuccess"].(float64); ok {
		CountDocumentImportSuccessInt := int(CountDocumentImportSuccess)
		o.CountDocumentImportSuccess = &CountDocumentImportSuccessInt
	}
	
	if CountDocumentImportFailure, ok := KnowledgeimportjobstatisticsMap["countDocumentImportFailure"].(float64); ok {
		CountDocumentImportFailureInt := int(CountDocumentImportFailure)
		o.CountDocumentImportFailure = &CountDocumentImportFailureInt
	}
	
	if CountCategoryValidationSuccess, ok := KnowledgeimportjobstatisticsMap["countCategoryValidationSuccess"].(float64); ok {
		CountCategoryValidationSuccessInt := int(CountCategoryValidationSuccess)
		o.CountCategoryValidationSuccess = &CountCategoryValidationSuccessInt
	}
	
	if CountCategoryValidationFailure, ok := KnowledgeimportjobstatisticsMap["countCategoryValidationFailure"].(float64); ok {
		CountCategoryValidationFailureInt := int(CountCategoryValidationFailure)
		o.CountCategoryValidationFailure = &CountCategoryValidationFailureInt
	}
	
	if CountCategoryImportSuccess, ok := KnowledgeimportjobstatisticsMap["countCategoryImportSuccess"].(float64); ok {
		CountCategoryImportSuccessInt := int(CountCategoryImportSuccess)
		o.CountCategoryImportSuccess = &CountCategoryImportSuccessInt
	}
	
	if CountCategoryImportFailure, ok := KnowledgeimportjobstatisticsMap["countCategoryImportFailure"].(float64); ok {
		CountCategoryImportFailureInt := int(CountCategoryImportFailure)
		o.CountCategoryImportFailure = &CountCategoryImportFailureInt
	}
	
	if CountLabelValidationSuccess, ok := KnowledgeimportjobstatisticsMap["countLabelValidationSuccess"].(float64); ok {
		CountLabelValidationSuccessInt := int(CountLabelValidationSuccess)
		o.CountLabelValidationSuccess = &CountLabelValidationSuccessInt
	}
	
	if CountLabelValidationFailure, ok := KnowledgeimportjobstatisticsMap["countLabelValidationFailure"].(float64); ok {
		CountLabelValidationFailureInt := int(CountLabelValidationFailure)
		o.CountLabelValidationFailure = &CountLabelValidationFailureInt
	}
	
	if CountLabelImportSuccess, ok := KnowledgeimportjobstatisticsMap["countLabelImportSuccess"].(float64); ok {
		CountLabelImportSuccessInt := int(CountLabelImportSuccess)
		o.CountLabelImportSuccess = &CountLabelImportSuccessInt
	}
	
	if CountLabelImportFailure, ok := KnowledgeimportjobstatisticsMap["countLabelImportFailure"].(float64); ok {
		CountLabelImportFailureInt := int(CountLabelImportFailure)
		o.CountLabelImportFailure = &CountLabelImportFailureInt
	}
	
	if MigrationDetected, ok := KnowledgeimportjobstatisticsMap["migrationDetected"].(bool); ok {
		o.MigrationDetected = &MigrationDetected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
