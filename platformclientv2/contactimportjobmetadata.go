package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactimportjobmetadata
type Contactimportjobmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FileName
	FileName *string `json:"fileName,omitempty"`

	// DryRunFailedCount
	DryRunFailedCount *int `json:"dryRunFailedCount,omitempty"`

	// DryRunSuccessCount
	DryRunSuccessCount *int `json:"dryRunSuccessCount,omitempty"`

	// DryRunReportDownloadUrl
	DryRunReportDownloadUrl *string `json:"dryRunReportDownloadUrl,omitempty"`

	// ImportFailedCount
	ImportFailedCount *int `json:"importFailedCount,omitempty"`

	// ImportSuccessCount
	ImportSuccessCount *int `json:"importSuccessCount,omitempty"`

	// ImportReportDownloadUrl
	ImportReportDownloadUrl *string `json:"importReportDownloadUrl,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactimportjobmetadata) SetField(field string, fieldValue interface{}) {
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

func (o Contactimportjobmetadata) MarshalJSON() ([]byte, error) {
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
	type Alias Contactimportjobmetadata
	
	return json.Marshal(&struct { 
		FileName *string `json:"fileName,omitempty"`
		
		DryRunFailedCount *int `json:"dryRunFailedCount,omitempty"`
		
		DryRunSuccessCount *int `json:"dryRunSuccessCount,omitempty"`
		
		DryRunReportDownloadUrl *string `json:"dryRunReportDownloadUrl,omitempty"`
		
		ImportFailedCount *int `json:"importFailedCount,omitempty"`
		
		ImportSuccessCount *int `json:"importSuccessCount,omitempty"`
		
		ImportReportDownloadUrl *string `json:"importReportDownloadUrl,omitempty"`
		Alias
	}{ 
		FileName: o.FileName,
		
		DryRunFailedCount: o.DryRunFailedCount,
		
		DryRunSuccessCount: o.DryRunSuccessCount,
		
		DryRunReportDownloadUrl: o.DryRunReportDownloadUrl,
		
		ImportFailedCount: o.ImportFailedCount,
		
		ImportSuccessCount: o.ImportSuccessCount,
		
		ImportReportDownloadUrl: o.ImportReportDownloadUrl,
		Alias:    (Alias)(o),
	})
}

func (o *Contactimportjobmetadata) UnmarshalJSON(b []byte) error {
	var ContactimportjobmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ContactimportjobmetadataMap)
	if err != nil {
		return err
	}
	
	if FileName, ok := ContactimportjobmetadataMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if DryRunFailedCount, ok := ContactimportjobmetadataMap["dryRunFailedCount"].(float64); ok {
		DryRunFailedCountInt := int(DryRunFailedCount)
		o.DryRunFailedCount = &DryRunFailedCountInt
	}
	
	if DryRunSuccessCount, ok := ContactimportjobmetadataMap["dryRunSuccessCount"].(float64); ok {
		DryRunSuccessCountInt := int(DryRunSuccessCount)
		o.DryRunSuccessCount = &DryRunSuccessCountInt
	}
	
	if DryRunReportDownloadUrl, ok := ContactimportjobmetadataMap["dryRunReportDownloadUrl"].(string); ok {
		o.DryRunReportDownloadUrl = &DryRunReportDownloadUrl
	}
    
	if ImportFailedCount, ok := ContactimportjobmetadataMap["importFailedCount"].(float64); ok {
		ImportFailedCountInt := int(ImportFailedCount)
		o.ImportFailedCount = &ImportFailedCountInt
	}
	
	if ImportSuccessCount, ok := ContactimportjobmetadataMap["importSuccessCount"].(float64); ok {
		ImportSuccessCountInt := int(ImportSuccessCount)
		o.ImportSuccessCount = &ImportSuccessCountInt
	}
	
	if ImportReportDownloadUrl, ok := ContactimportjobmetadataMap["importReportDownloadUrl"].(string); ok {
		o.ImportReportDownloadUrl = &ImportReportDownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactimportjobmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
