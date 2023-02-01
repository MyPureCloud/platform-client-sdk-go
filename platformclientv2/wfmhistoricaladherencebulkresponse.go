package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkresponse
type Wfmhistoricaladherencebulkresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Job - A reference to the job that was started by the request
	Job *Wfmhistoricaladherencebulkjobreference `json:"job,omitempty"`

	// DownloadUrls - The uri list to GET the results of the Historical Adherence query. This field is populated only if query state is Complete
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`

	// DownloadResult - Results will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Wfmhistoricaladherencebulkresult `json:"downloadResult,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmhistoricaladherencebulkresponse) SetField(field string, fieldValue interface{}) {
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

func (o Wfmhistoricaladherencebulkresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmhistoricaladherencebulkresponse
	
	return json.Marshal(&struct { 
		Job *Wfmhistoricaladherencebulkjobreference `json:"job,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		DownloadResult *Wfmhistoricaladherencebulkresult `json:"downloadResult,omitempty"`
		Alias
	}{ 
		Job: o.Job,
		
		DownloadUrls: o.DownloadUrls,
		
		DownloadResult: o.DownloadResult,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkresponse) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkresponseMap)
	if err != nil {
		return err
	}
	
	if Job, ok := WfmhistoricaladherencebulkresponseMap["job"].(map[string]interface{}); ok {
		JobString, _ := json.Marshal(Job)
		json.Unmarshal(JobString, &o.Job)
	}
	
	if DownloadUrls, ok := WfmhistoricaladherencebulkresponseMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if DownloadResult, ok := WfmhistoricaladherencebulkresponseMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
