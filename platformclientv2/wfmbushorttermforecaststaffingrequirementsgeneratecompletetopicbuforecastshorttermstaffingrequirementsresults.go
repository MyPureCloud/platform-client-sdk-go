package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults
type Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WeekNumber
	WeekNumber *int `json:"weekNumber,omitempty"`

	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// DownloadUrlExpirationDate
	DownloadUrlExpirationDate *time.Time `json:"downloadUrlExpirationDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DownloadUrlExpirationDate", }
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
	type Alias Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults
	
	DownloadUrlExpirationDate := new(string)
	if o.DownloadUrlExpirationDate != nil {
		
		*DownloadUrlExpirationDate = timeutil.Strftime(o.DownloadUrlExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DownloadUrlExpirationDate = nil
	}
	
	return json.Marshal(&struct { 
		WeekNumber *int `json:"weekNumber,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadUrlExpirationDate *string `json:"downloadUrlExpirationDate,omitempty"`
		Alias
	}{ 
		WeekNumber: o.WeekNumber,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadUrlExpirationDate: DownloadUrlExpirationDate,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresultsMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresultsMap)
	if err != nil {
		return err
	}
	
	if WeekNumber, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresultsMap["weekNumber"].(float64); ok {
		WeekNumberInt := int(WeekNumber)
		o.WeekNumber = &WeekNumberInt
	}
	
	if DownloadUrl, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresultsMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if downloadUrlExpirationDateString, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresultsMap["downloadUrlExpirationDate"].(string); ok {
		DownloadUrlExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", downloadUrlExpirationDateString)
		o.DownloadUrlExpirationDate = &DownloadUrlExpirationDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
