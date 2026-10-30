package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V3sourcefilecountentry
type V3sourcefilecountentry struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Source - Knowledge source reference for this entry.
	Source *V3sourceref `json:"source,omitempty"`

	// SelectedTagsFileCount - Unique file count matching the tag filter for this source.
	SelectedTagsFileCount *int `json:"selectedTagsFileCount,omitempty"`

	// TotalFileCount - Total unique file count for this source without tag filters.
	TotalFileCount *int `json:"totalFileCount,omitempty"`

	// FileNames - File names when includeFileNames is enabled.
	FileNames *[]string `json:"fileNames,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V3sourcefilecountentry) SetField(field string, fieldValue interface{}) {
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

func (o V3sourcefilecountentry) MarshalJSON() ([]byte, error) {
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
	type Alias V3sourcefilecountentry
	
	return json.Marshal(&struct { 
		Source *V3sourceref `json:"source,omitempty"`
		
		SelectedTagsFileCount *int `json:"selectedTagsFileCount,omitempty"`
		
		TotalFileCount *int `json:"totalFileCount,omitempty"`
		
		FileNames *[]string `json:"fileNames,omitempty"`
		Alias
	}{ 
		Source: o.Source,
		
		SelectedTagsFileCount: o.SelectedTagsFileCount,
		
		TotalFileCount: o.TotalFileCount,
		
		FileNames: o.FileNames,
		Alias:    (Alias)(o),
	})
}

func (o *V3sourcefilecountentry) UnmarshalJSON(b []byte) error {
	var V3sourcefilecountentryMap map[string]interface{}
	err := json.Unmarshal(b, &V3sourcefilecountentryMap)
	if err != nil {
		return err
	}
	
	if Source, ok := V3sourcefilecountentryMap["source"].(map[string]interface{}); ok {
		SourceString, _ := json.Marshal(Source)
		json.Unmarshal(SourceString, &o.Source)
	}
	
	if SelectedTagsFileCount, ok := V3sourcefilecountentryMap["selectedTagsFileCount"].(float64); ok {
		SelectedTagsFileCountInt := int(SelectedTagsFileCount)
		o.SelectedTagsFileCount = &SelectedTagsFileCountInt
	}
	
	if TotalFileCount, ok := V3sourcefilecountentryMap["totalFileCount"].(float64); ok {
		TotalFileCountInt := int(TotalFileCount)
		o.TotalFileCount = &TotalFileCountInt
	}
	
	if FileNames, ok := V3sourcefilecountentryMap["fileNames"].([]interface{}); ok {
		FileNamesString, _ := json.Marshal(FileNames)
		json.Unmarshal(FileNamesString, &o.FileNames)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V3sourcefilecountentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
