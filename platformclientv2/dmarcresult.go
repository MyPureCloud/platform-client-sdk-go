package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dmarcresult - Represents the DMARC verification result for an email domain
type Dmarcresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - The DMARC status of this domain
	Status *string `json:"status,omitempty"`

	// DetectedPolicy - The DMARC policy that was detected in the associated DNS record, if one is present
	DetectedPolicy *string `json:"detectedPolicy,omitempty"`

	// DateChecked - The date of the most recent check of the domain's DMARC DNS record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateChecked *time.Time `json:"dateChecked,omitempty"`

	// Records - The minimum DMARC DNS record that Genesys Cloud recommends
	Records *[]Record `json:"records,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dmarcresult) SetField(field string, fieldValue interface{}) {
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

func (o Dmarcresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateChecked", }
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
	type Alias Dmarcresult
	
	DateChecked := new(string)
	if o.DateChecked != nil {
		
		*DateChecked = timeutil.Strftime(o.DateChecked, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateChecked = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		DetectedPolicy *string `json:"detectedPolicy,omitempty"`
		
		DateChecked *string `json:"dateChecked,omitempty"`
		
		Records *[]Record `json:"records,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		DetectedPolicy: o.DetectedPolicy,
		
		DateChecked: DateChecked,
		
		Records: o.Records,
		Alias:    (Alias)(o),
	})
}

func (o *Dmarcresult) UnmarshalJSON(b []byte) error {
	var DmarcresultMap map[string]interface{}
	err := json.Unmarshal(b, &DmarcresultMap)
	if err != nil {
		return err
	}
	
	if Status, ok := DmarcresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if DetectedPolicy, ok := DmarcresultMap["detectedPolicy"].(string); ok {
		o.DetectedPolicy = &DetectedPolicy
	}
    
	if dateCheckedString, ok := DmarcresultMap["dateChecked"].(string); ok {
		DateChecked, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCheckedString)
		o.DateChecked = &DateChecked
	}
	
	if Records, ok := DmarcresultMap["records"].([]interface{}); ok {
		RecordsString, _ := json.Marshal(Records)
		json.Unmarshal(RecordsString, &o.Records)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dmarcresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
