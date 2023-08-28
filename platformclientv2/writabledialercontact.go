package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Writabledialercontact
type Writabledialercontact struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ContactListId - The identifier of the contact list containing this contact.
	ContactListId *string `json:"contactListId,omitempty"`

	// Data - An ordered map of the contact's columns and corresponding values.
	Data *map[string]string `json:"data,omitempty"`

	// LatestSmsEvaluations - A map of SMS records for the contact phone columns.
	LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`

	// LatestEmailEvaluations - A map of email records for the contact email columns.
	LatestEmailEvaluations *map[string]Messageevaluation `json:"latestEmailEvaluations,omitempty"`

	// Callable - Indicates whether or not the contact can be called.
	Callable *bool `json:"callable,omitempty"`

	// PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
	PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`

	// ContactableStatus - A map of media types (Voice, SMS and Email) to ContactableStatus, which indicates if the contact can be contacted using the specified media type.
	ContactableStatus *map[string]Contactablestatus `json:"contactableStatus,omitempty"`

	// DateCreated - Timestamp for when the contact was added. Contacts added prior to 2023 September 1 may be missing this value. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Writabledialercontact) SetField(field string, fieldValue interface{}) {
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

func (o Writabledialercontact) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Writabledialercontact
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		Data *map[string]string `json:"data,omitempty"`
		
		LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`
		
		LatestEmailEvaluations *map[string]Messageevaluation `json:"latestEmailEvaluations,omitempty"`
		
		Callable *bool `json:"callable,omitempty"`
		
		PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`
		
		ContactableStatus *map[string]Contactablestatus `json:"contactableStatus,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ContactListId: o.ContactListId,
		
		Data: o.Data,
		
		LatestSmsEvaluations: o.LatestSmsEvaluations,
		
		LatestEmailEvaluations: o.LatestEmailEvaluations,
		
		Callable: o.Callable,
		
		PhoneNumberStatus: o.PhoneNumberStatus,
		
		ContactableStatus: o.ContactableStatus,
		
		DateCreated: DateCreated,
		Alias:    (Alias)(o),
	})
}

func (o *Writabledialercontact) UnmarshalJSON(b []byte) error {
	var WritabledialercontactMap map[string]interface{}
	err := json.Unmarshal(b, &WritabledialercontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WritabledialercontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ContactListId, ok := WritabledialercontactMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
    
	if Data, ok := WritabledialercontactMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if LatestSmsEvaluations, ok := WritabledialercontactMap["latestSmsEvaluations"].(map[string]interface{}); ok {
		LatestSmsEvaluationsString, _ := json.Marshal(LatestSmsEvaluations)
		json.Unmarshal(LatestSmsEvaluationsString, &o.LatestSmsEvaluations)
	}
	
	if LatestEmailEvaluations, ok := WritabledialercontactMap["latestEmailEvaluations"].(map[string]interface{}); ok {
		LatestEmailEvaluationsString, _ := json.Marshal(LatestEmailEvaluations)
		json.Unmarshal(LatestEmailEvaluationsString, &o.LatestEmailEvaluations)
	}
	
	if Callable, ok := WritabledialercontactMap["callable"].(bool); ok {
		o.Callable = &Callable
	}
    
	if PhoneNumberStatus, ok := WritabledialercontactMap["phoneNumberStatus"].(map[string]interface{}); ok {
		PhoneNumberStatusString, _ := json.Marshal(PhoneNumberStatus)
		json.Unmarshal(PhoneNumberStatusString, &o.PhoneNumberStatus)
	}
	
	if ContactableStatus, ok := WritabledialercontactMap["contactableStatus"].(map[string]interface{}); ok {
		ContactableStatusString, _ := json.Marshal(ContactableStatus)
		json.Unmarshal(ContactableStatusString, &o.ContactableStatus)
	}
	
	if dateCreatedString, ok := WritabledialercontactMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Writabledialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
