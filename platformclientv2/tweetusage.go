package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Tweetusage
type Tweetusage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IngestionLimit - Ingestion limit
	IngestionLimit *int `json:"ingestionLimit,omitempty"`

	// TweetCount - The number of tweets consumed
	TweetCount *int `json:"tweetCount,omitempty"`

	// DateStart - The start of the usage period for the currently consumed tweets. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Tweetusage) SetField(field string, fieldValue interface{}) {
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

func (o Tweetusage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart", }
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
	type Alias Tweetusage
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		IngestionLimit *int `json:"ingestionLimit,omitempty"`
		
		TweetCount *int `json:"tweetCount,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		Alias
	}{ 
		IngestionLimit: o.IngestionLimit,
		
		TweetCount: o.TweetCount,
		
		DateStart: DateStart,
		Alias:    (Alias)(o),
	})
}

func (o *Tweetusage) UnmarshalJSON(b []byte) error {
	var TweetusageMap map[string]interface{}
	err := json.Unmarshal(b, &TweetusageMap)
	if err != nil {
		return err
	}
	
	if IngestionLimit, ok := TweetusageMap["ingestionLimit"].(float64); ok {
		IngestionLimitInt := int(IngestionLimit)
		o.IngestionLimit = &IngestionLimitInt
	}
	
	if TweetCount, ok := TweetusageMap["tweetCount"].(float64); ok {
		TweetCountInt := int(TweetCount)
		o.TweetCount = &TweetCountInt
	}
	
	if dateStartString, ok := TweetusageMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Tweetusage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
