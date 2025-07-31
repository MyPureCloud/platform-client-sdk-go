package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createreprocessjobrequest
type Createreprocessjobrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name given for the job.
	Name *string `json:"name,omitempty"`

	// Description - The description of the job.
	Description *string `json:"description,omitempty"`

	// DateStart - The start date for the search criteria. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The end date for the search criteria. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// Programs - A list of program IDs to filter conversations by.
	Programs *[]string `json:"programs,omitempty"`

	// MediaTypes - The types of conversations to reprocess.
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// Dialects - The dialects to filter by the conversations.
	Dialects *[]string `json:"dialects,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createreprocessjobrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createreprocessjobrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd", }
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
	type Alias Createreprocessjobrequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Programs *[]string `json:"programs,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		Dialects *[]string `json:"dialects,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		Programs: o.Programs,
		
		MediaTypes: o.MediaTypes,
		
		Dialects: o.Dialects,
		Alias:    (Alias)(o),
	})
}

func (o *Createreprocessjobrequest) UnmarshalJSON(b []byte) error {
	var CreatereprocessjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatereprocessjobrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatereprocessjobrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CreatereprocessjobrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateStartString, ok := CreatereprocessjobrequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := CreatereprocessjobrequestMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if Programs, ok := CreatereprocessjobrequestMap["programs"].([]interface{}); ok {
		ProgramsString, _ := json.Marshal(Programs)
		json.Unmarshal(ProgramsString, &o.Programs)
	}
	
	if MediaTypes, ok := CreatereprocessjobrequestMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if Dialects, ok := CreatereprocessjobrequestMap["dialects"].([]interface{}); ok {
		DialectsString, _ := json.Marshal(Dialects)
		json.Unmarshal(DialectsString, &o.Dialects)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createreprocessjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
