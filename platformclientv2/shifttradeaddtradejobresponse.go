package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradeaddtradejobresponse
type Shifttradeaddtradejobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Status - The status of the job
	Status *string `json:"status,omitempty"`

	// VarType - The type of the job
	VarType *string `json:"type,omitempty"`

	// DownloadUrl - The URL where completed results might be available for download in case the result body for that job type is too large
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// VarError - Any error information, only set if the status == 'Error'
	VarError *Errorbody `json:"error,omitempty"`

	// AddTradeResult - Results for AddTrade job type
	AddTradeResult *Shifttraderesponseitem `json:"addTradeResult,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shifttradeaddtradejobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Shifttradeaddtradejobresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Shifttradeaddtradejobresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		AddTradeResult *Shifttraderesponseitem `json:"addTradeResult,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		VarType: o.VarType,
		
		DownloadUrl: o.DownloadUrl,
		
		VarError: o.VarError,
		
		AddTradeResult: o.AddTradeResult,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Shifttradeaddtradejobresponse) UnmarshalJSON(b []byte) error {
	var ShifttradeaddtradejobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradeaddtradejobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ShifttradeaddtradejobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := ShifttradeaddtradejobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarType, ok := ShifttradeaddtradejobresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if DownloadUrl, ok := ShifttradeaddtradejobresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if VarError, ok := ShifttradeaddtradejobresponseMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if AddTradeResult, ok := ShifttradeaddtradejobresponseMap["addTradeResult"].(map[string]interface{}); ok {
		AddTradeResultString, _ := json.Marshal(AddTradeResult)
		json.Unmarshal(AddTradeResultString, &o.AddTradeResult)
	}
	
	if SelfUri, ok := ShifttradeaddtradejobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradeaddtradejobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
