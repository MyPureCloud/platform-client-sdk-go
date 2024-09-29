package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesyncjobstatistics
type Knowledgesyncjobstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CountDocumentImportActivityCreate - Number of documents will be created by the sync.
	CountDocumentImportActivityCreate *int `json:"countDocumentImportActivityCreate,omitempty"`

	// CountDocumentImportActivityUpdate - Number of documents will be updated by the sync.
	CountDocumentImportActivityUpdate *int `json:"countDocumentImportActivityUpdate,omitempty"`

	// CountDocumentStateDraft - Number of documents will be imported as draft.
	CountDocumentStateDraft *int `json:"countDocumentStateDraft,omitempty"`

	// CountDocumentStatePublished - Number of documents will be imported as published.
	CountDocumentStatePublished *int `json:"countDocumentStatePublished,omitempty"`

	// CountDocumentImportSuccess - Number of imported documents.
	CountDocumentImportSuccess *int `json:"countDocumentImportSuccess,omitempty"`

	// CountDocumentImportFailure - Number of documents failed to import.
	CountDocumentImportFailure *int `json:"countDocumentImportFailure,omitempty"`

	// CountCategoryImportSuccess - Number of imported categories.
	CountCategoryImportSuccess *int `json:"countCategoryImportSuccess,omitempty"`

	// CountCategoryImportFailure - Number of categories failed to import.
	CountCategoryImportFailure *int `json:"countCategoryImportFailure,omitempty"`

	// CountLabelImportSuccess - Number of imported labels.
	CountLabelImportSuccess *int `json:"countLabelImportSuccess,omitempty"`

	// CountLabelImportFailure - Number of labels failed to import.
	CountLabelImportFailure *int `json:"countLabelImportFailure,omitempty"`

	// CountDocumentDeleteSuccess - Number of documents will be deleted by the sync.
	CountDocumentDeleteSuccess *int `json:"countDocumentDeleteSuccess,omitempty"`

	// CountDocumentDeleteFailure - Number of documents failed to delete.
	CountDocumentDeleteFailure *int `json:"countDocumentDeleteFailure,omitempty"`

	// CountCategoryDeleteSuccess - Number of successfully deleted categories.
	CountCategoryDeleteSuccess *int `json:"countCategoryDeleteSuccess,omitempty"`

	// CountCategoryDeleteFailure - Number of categories failed to delete.
	CountCategoryDeleteFailure *int `json:"countCategoryDeleteFailure,omitempty"`

	// CountLabelDeleteSuccess - Number of successfully deleted labels.
	CountLabelDeleteSuccess *int `json:"countLabelDeleteSuccess,omitempty"`

	// CountLabelDeleteFailure - Number of labels failed to delete.
	CountLabelDeleteFailure *int `json:"countLabelDeleteFailure,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesyncjobstatistics) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesyncjobstatistics) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesyncjobstatistics
	
	return json.Marshal(&struct { 
		CountDocumentImportActivityCreate *int `json:"countDocumentImportActivityCreate,omitempty"`
		
		CountDocumentImportActivityUpdate *int `json:"countDocumentImportActivityUpdate,omitempty"`
		
		CountDocumentStateDraft *int `json:"countDocumentStateDraft,omitempty"`
		
		CountDocumentStatePublished *int `json:"countDocumentStatePublished,omitempty"`
		
		CountDocumentImportSuccess *int `json:"countDocumentImportSuccess,omitempty"`
		
		CountDocumentImportFailure *int `json:"countDocumentImportFailure,omitempty"`
		
		CountCategoryImportSuccess *int `json:"countCategoryImportSuccess,omitempty"`
		
		CountCategoryImportFailure *int `json:"countCategoryImportFailure,omitempty"`
		
		CountLabelImportSuccess *int `json:"countLabelImportSuccess,omitempty"`
		
		CountLabelImportFailure *int `json:"countLabelImportFailure,omitempty"`
		
		CountDocumentDeleteSuccess *int `json:"countDocumentDeleteSuccess,omitempty"`
		
		CountDocumentDeleteFailure *int `json:"countDocumentDeleteFailure,omitempty"`
		
		CountCategoryDeleteSuccess *int `json:"countCategoryDeleteSuccess,omitempty"`
		
		CountCategoryDeleteFailure *int `json:"countCategoryDeleteFailure,omitempty"`
		
		CountLabelDeleteSuccess *int `json:"countLabelDeleteSuccess,omitempty"`
		
		CountLabelDeleteFailure *int `json:"countLabelDeleteFailure,omitempty"`
		Alias
	}{ 
		CountDocumentImportActivityCreate: o.CountDocumentImportActivityCreate,
		
		CountDocumentImportActivityUpdate: o.CountDocumentImportActivityUpdate,
		
		CountDocumentStateDraft: o.CountDocumentStateDraft,
		
		CountDocumentStatePublished: o.CountDocumentStatePublished,
		
		CountDocumentImportSuccess: o.CountDocumentImportSuccess,
		
		CountDocumentImportFailure: o.CountDocumentImportFailure,
		
		CountCategoryImportSuccess: o.CountCategoryImportSuccess,
		
		CountCategoryImportFailure: o.CountCategoryImportFailure,
		
		CountLabelImportSuccess: o.CountLabelImportSuccess,
		
		CountLabelImportFailure: o.CountLabelImportFailure,
		
		CountDocumentDeleteSuccess: o.CountDocumentDeleteSuccess,
		
		CountDocumentDeleteFailure: o.CountDocumentDeleteFailure,
		
		CountCategoryDeleteSuccess: o.CountCategoryDeleteSuccess,
		
		CountCategoryDeleteFailure: o.CountCategoryDeleteFailure,
		
		CountLabelDeleteSuccess: o.CountLabelDeleteSuccess,
		
		CountLabelDeleteFailure: o.CountLabelDeleteFailure,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesyncjobstatistics) UnmarshalJSON(b []byte) error {
	var KnowledgesyncjobstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesyncjobstatisticsMap)
	if err != nil {
		return err
	}
	
	if CountDocumentImportActivityCreate, ok := KnowledgesyncjobstatisticsMap["countDocumentImportActivityCreate"].(float64); ok {
		CountDocumentImportActivityCreateInt := int(CountDocumentImportActivityCreate)
		o.CountDocumentImportActivityCreate = &CountDocumentImportActivityCreateInt
	}
	
	if CountDocumentImportActivityUpdate, ok := KnowledgesyncjobstatisticsMap["countDocumentImportActivityUpdate"].(float64); ok {
		CountDocumentImportActivityUpdateInt := int(CountDocumentImportActivityUpdate)
		o.CountDocumentImportActivityUpdate = &CountDocumentImportActivityUpdateInt
	}
	
	if CountDocumentStateDraft, ok := KnowledgesyncjobstatisticsMap["countDocumentStateDraft"].(float64); ok {
		CountDocumentStateDraftInt := int(CountDocumentStateDraft)
		o.CountDocumentStateDraft = &CountDocumentStateDraftInt
	}
	
	if CountDocumentStatePublished, ok := KnowledgesyncjobstatisticsMap["countDocumentStatePublished"].(float64); ok {
		CountDocumentStatePublishedInt := int(CountDocumentStatePublished)
		o.CountDocumentStatePublished = &CountDocumentStatePublishedInt
	}
	
	if CountDocumentImportSuccess, ok := KnowledgesyncjobstatisticsMap["countDocumentImportSuccess"].(float64); ok {
		CountDocumentImportSuccessInt := int(CountDocumentImportSuccess)
		o.CountDocumentImportSuccess = &CountDocumentImportSuccessInt
	}
	
	if CountDocumentImportFailure, ok := KnowledgesyncjobstatisticsMap["countDocumentImportFailure"].(float64); ok {
		CountDocumentImportFailureInt := int(CountDocumentImportFailure)
		o.CountDocumentImportFailure = &CountDocumentImportFailureInt
	}
	
	if CountCategoryImportSuccess, ok := KnowledgesyncjobstatisticsMap["countCategoryImportSuccess"].(float64); ok {
		CountCategoryImportSuccessInt := int(CountCategoryImportSuccess)
		o.CountCategoryImportSuccess = &CountCategoryImportSuccessInt
	}
	
	if CountCategoryImportFailure, ok := KnowledgesyncjobstatisticsMap["countCategoryImportFailure"].(float64); ok {
		CountCategoryImportFailureInt := int(CountCategoryImportFailure)
		o.CountCategoryImportFailure = &CountCategoryImportFailureInt
	}
	
	if CountLabelImportSuccess, ok := KnowledgesyncjobstatisticsMap["countLabelImportSuccess"].(float64); ok {
		CountLabelImportSuccessInt := int(CountLabelImportSuccess)
		o.CountLabelImportSuccess = &CountLabelImportSuccessInt
	}
	
	if CountLabelImportFailure, ok := KnowledgesyncjobstatisticsMap["countLabelImportFailure"].(float64); ok {
		CountLabelImportFailureInt := int(CountLabelImportFailure)
		o.CountLabelImportFailure = &CountLabelImportFailureInt
	}
	
	if CountDocumentDeleteSuccess, ok := KnowledgesyncjobstatisticsMap["countDocumentDeleteSuccess"].(float64); ok {
		CountDocumentDeleteSuccessInt := int(CountDocumentDeleteSuccess)
		o.CountDocumentDeleteSuccess = &CountDocumentDeleteSuccessInt
	}
	
	if CountDocumentDeleteFailure, ok := KnowledgesyncjobstatisticsMap["countDocumentDeleteFailure"].(float64); ok {
		CountDocumentDeleteFailureInt := int(CountDocumentDeleteFailure)
		o.CountDocumentDeleteFailure = &CountDocumentDeleteFailureInt
	}
	
	if CountCategoryDeleteSuccess, ok := KnowledgesyncjobstatisticsMap["countCategoryDeleteSuccess"].(float64); ok {
		CountCategoryDeleteSuccessInt := int(CountCategoryDeleteSuccess)
		o.CountCategoryDeleteSuccess = &CountCategoryDeleteSuccessInt
	}
	
	if CountCategoryDeleteFailure, ok := KnowledgesyncjobstatisticsMap["countCategoryDeleteFailure"].(float64); ok {
		CountCategoryDeleteFailureInt := int(CountCategoryDeleteFailure)
		o.CountCategoryDeleteFailure = &CountCategoryDeleteFailureInt
	}
	
	if CountLabelDeleteSuccess, ok := KnowledgesyncjobstatisticsMap["countLabelDeleteSuccess"].(float64); ok {
		CountLabelDeleteSuccessInt := int(CountLabelDeleteSuccess)
		o.CountLabelDeleteSuccess = &CountLabelDeleteSuccessInt
	}
	
	if CountLabelDeleteFailure, ok := KnowledgesyncjobstatisticsMap["countLabelDeleteFailure"].(float64); ok {
		CountLabelDeleteFailureInt := int(CountLabelDeleteFailure)
		o.CountLabelDeleteFailure = &CountLabelDeleteFailureInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesyncjobstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
