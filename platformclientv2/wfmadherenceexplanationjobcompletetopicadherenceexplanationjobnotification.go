package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification
type Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// AdherenceExplanation
	AdherenceExplanation *Wfmadherenceexplanationjobcompletetopicadherenceexplanationchangednotification `json:"adherenceExplanation,omitempty"`

	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// VarError
	VarError *Wfmadherenceexplanationjobcompletetopicerrorbody `json:"error,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		AdherenceExplanation *Wfmadherenceexplanationjobcompletetopicadherenceexplanationchangednotification `json:"adherenceExplanation,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		VarError *Wfmadherenceexplanationjobcompletetopicerrorbody `json:"error,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		AdherenceExplanation: o.AdherenceExplanation,
		
		DownloadUrl: o.DownloadUrl,
		
		VarError: o.VarError,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification) UnmarshalJSON(b []byte) error {
	var WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Status, ok := WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if AdherenceExplanation, ok := WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap["adherenceExplanation"].(map[string]interface{}); ok {
		AdherenceExplanationString, _ := json.Marshal(AdherenceExplanation)
		json.Unmarshal(AdherenceExplanationString, &o.AdherenceExplanation)
	}
	
	if DownloadUrl, ok := WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if VarError, ok := WfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotificationMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmadherenceexplanationjobcompletetopicadherenceexplanationjobnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
