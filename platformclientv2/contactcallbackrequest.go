package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcallbackrequest
type Contactcallbackrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CampaignId - Campaign identifier
	CampaignId *string `json:"campaignId,omitempty"`

	// ContactListId - Contact list identifier
	ContactListId *string `json:"contactListId,omitempty"`

	// ContactId - Contact identifier
	ContactId *string `json:"contactId,omitempty"`

	// PhoneColumn - Name of the phone column containing the number to be called
	PhoneColumn *string `json:"phoneColumn,omitempty"`

	// Schedule - The scheduled time for the callback as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ\", example = \"2016-01-02T16:59:59\"
	Schedule *string `json:"schedule,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactcallbackrequest) SetField(field string, fieldValue interface{}) {
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

func (o Contactcallbackrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Contactcallbackrequest
	
	return json.Marshal(&struct { 
		CampaignId *string `json:"campaignId,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		ContactId *string `json:"contactId,omitempty"`
		
		PhoneColumn *string `json:"phoneColumn,omitempty"`
		
		Schedule *string `json:"schedule,omitempty"`
		Alias
	}{ 
		CampaignId: o.CampaignId,
		
		ContactListId: o.ContactListId,
		
		ContactId: o.ContactId,
		
		PhoneColumn: o.PhoneColumn,
		
		Schedule: o.Schedule,
		Alias:    (Alias)(o),
	})
}

func (o *Contactcallbackrequest) UnmarshalJSON(b []byte) error {
	var ContactcallbackrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcallbackrequestMap)
	if err != nil {
		return err
	}
	
	if CampaignId, ok := ContactcallbackrequestMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
    
	if ContactListId, ok := ContactcallbackrequestMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
    
	if ContactId, ok := ContactcallbackrequestMap["contactId"].(string); ok {
		o.ContactId = &ContactId
	}
    
	if PhoneColumn, ok := ContactcallbackrequestMap["phoneColumn"].(string); ok {
		o.PhoneColumn = &PhoneColumn
	}
    
	if Schedule, ok := ContactcallbackrequestMap["schedule"].(string); ok {
		o.Schedule = &Schedule
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcallbackrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
