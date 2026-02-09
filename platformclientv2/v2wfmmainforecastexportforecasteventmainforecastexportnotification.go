package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2wfmmainforecastexportforecasteventmainforecastexportnotification
type V2wfmmainforecastexportforecasteventmainforecastexportnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExportType
	ExportType *string `json:"exportType,omitempty"`

	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// VarError
	VarError *V2wfmmainforecastexportforecasteventmainforecasterrorbody `json:"error,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2wfmmainforecastexportforecasteventmainforecastexportnotification) SetField(field string, fieldValue interface{}) {
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

func (o V2wfmmainforecastexportforecasteventmainforecastexportnotification) MarshalJSON() ([]byte, error) {
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
	type Alias V2wfmmainforecastexportforecasteventmainforecastexportnotification
	
	return json.Marshal(&struct { 
		ExportType *string `json:"exportType,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		VarError *V2wfmmainforecastexportforecasteventmainforecasterrorbody `json:"error,omitempty"`
		Alias
	}{ 
		ExportType: o.ExportType,
		
		BusinessUnitId: o.BusinessUnitId,
		
		Id: o.Id,
		
		Status: o.Status,
		
		DownloadUrl: o.DownloadUrl,
		
		VarError: o.VarError,
		Alias:    (Alias)(o),
	})
}

func (o *V2wfmmainforecastexportforecasteventmainforecastexportnotification) UnmarshalJSON(b []byte) error {
	var V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap)
	if err != nil {
		return err
	}
	
	if ExportType, ok := V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap["exportType"].(string); ok {
		o.ExportType = &ExportType
	}
    
	if BusinessUnitId, ok := V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if Id, ok := V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if DownloadUrl, ok := V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if VarError, ok := V2wfmmainforecastexportforecasteventmainforecastexportnotificationMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2wfmmainforecastexportforecasteventmainforecastexportnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
