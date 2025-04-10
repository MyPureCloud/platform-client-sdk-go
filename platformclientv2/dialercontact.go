package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontact
type Dialercontact struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ContactListId - The identifier of the contact list containing this contact.
	ContactListId *string `json:"contactListId,omitempty"`

	// Data - An ordered map of the contact's columns and corresponding values.
	Data *map[string]string `json:"data,omitempty"`

	// CallRecords - A map of call records for the contact phone columns.
	CallRecords *map[string]Callrecord `json:"callRecords,omitempty"`

	// LatestSmsEvaluations - A map of SMS records for the contact phone columns.
	LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`

	// LatestEmailEvaluations - A map of email records for the contact email columns.
	LatestEmailEvaluations *map[string]Messageevaluation `json:"latestEmailEvaluations,omitempty"`

	// LatestWhatsAppEvaluations - A map of whatsapp records for the contact whatsapp columns.
	LatestWhatsAppEvaluations *map[string]Messageevaluation `json:"latestWhatsAppEvaluations,omitempty"`

	// Callable - Indicates whether or not the contact can be called.
	Callable *bool `json:"callable,omitempty"`

	// PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
	PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`

	// ContactableStatus - A map of media types (Voice, SMS and Email) to ContactableStatus, which indicates if the contact can be contacted using the specified media type.
	ContactableStatus *map[string]Contactablestatus `json:"contactableStatus,omitempty"`

	// ContactColumnTimeZones - Map containing data about the timezone the contact is mapped to. This will only be populated if the contact list has automatic timezone mapping turned on. The key is the column name. The value is the timezone it mapped to and the type of column: Phone or Zip
	ContactColumnTimeZones *map[string]Contactcolumntimezone `json:"contactColumnTimeZones,omitempty"`

	// ConfigurationOverrides - the priority property within ConfigurationOverides indicates whether or not the contact to be placed in front of the queue or at the end of the queue
	ConfigurationOverrides *Configurationoverrides `json:"configurationOverrides,omitempty"`

	// DateCreated - Timestamp for when the contact was added. Contacts added prior to 2023 September 1 may be missing this value. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercontact) SetField(field string, fieldValue interface{}) {
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

func (o Dialercontact) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercontact
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		Data *map[string]string `json:"data,omitempty"`
		
		CallRecords *map[string]Callrecord `json:"callRecords,omitempty"`
		
		LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`
		
		LatestEmailEvaluations *map[string]Messageevaluation `json:"latestEmailEvaluations,omitempty"`
		
		LatestWhatsAppEvaluations *map[string]Messageevaluation `json:"latestWhatsAppEvaluations,omitempty"`
		
		Callable *bool `json:"callable,omitempty"`
		
		PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`
		
		ContactableStatus *map[string]Contactablestatus `json:"contactableStatus,omitempty"`
		
		ContactColumnTimeZones *map[string]Contactcolumntimezone `json:"contactColumnTimeZones,omitempty"`
		
		ConfigurationOverrides *Configurationoverrides `json:"configurationOverrides,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ContactListId: o.ContactListId,
		
		Data: o.Data,
		
		CallRecords: o.CallRecords,
		
		LatestSmsEvaluations: o.LatestSmsEvaluations,
		
		LatestEmailEvaluations: o.LatestEmailEvaluations,
		
		LatestWhatsAppEvaluations: o.LatestWhatsAppEvaluations,
		
		Callable: o.Callable,
		
		PhoneNumberStatus: o.PhoneNumberStatus,
		
		ContactableStatus: o.ContactableStatus,
		
		ContactColumnTimeZones: o.ContactColumnTimeZones,
		
		ConfigurationOverrides: o.ConfigurationOverrides,
		
		DateCreated: DateCreated,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Dialercontact) UnmarshalJSON(b []byte) error {
	var DialercontactMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialercontactMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ContactListId, ok := DialercontactMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
    
	if Data, ok := DialercontactMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if CallRecords, ok := DialercontactMap["callRecords"].(map[string]interface{}); ok {
		CallRecordsString, _ := json.Marshal(CallRecords)
		json.Unmarshal(CallRecordsString, &o.CallRecords)
	}
	
	if LatestSmsEvaluations, ok := DialercontactMap["latestSmsEvaluations"].(map[string]interface{}); ok {
		LatestSmsEvaluationsString, _ := json.Marshal(LatestSmsEvaluations)
		json.Unmarshal(LatestSmsEvaluationsString, &o.LatestSmsEvaluations)
	}
	
	if LatestEmailEvaluations, ok := DialercontactMap["latestEmailEvaluations"].(map[string]interface{}); ok {
		LatestEmailEvaluationsString, _ := json.Marshal(LatestEmailEvaluations)
		json.Unmarshal(LatestEmailEvaluationsString, &o.LatestEmailEvaluations)
	}
	
	if LatestWhatsAppEvaluations, ok := DialercontactMap["latestWhatsAppEvaluations"].(map[string]interface{}); ok {
		LatestWhatsAppEvaluationsString, _ := json.Marshal(LatestWhatsAppEvaluations)
		json.Unmarshal(LatestWhatsAppEvaluationsString, &o.LatestWhatsAppEvaluations)
	}
	
	if Callable, ok := DialercontactMap["callable"].(bool); ok {
		o.Callable = &Callable
	}
    
	if PhoneNumberStatus, ok := DialercontactMap["phoneNumberStatus"].(map[string]interface{}); ok {
		PhoneNumberStatusString, _ := json.Marshal(PhoneNumberStatus)
		json.Unmarshal(PhoneNumberStatusString, &o.PhoneNumberStatus)
	}
	
	if ContactableStatus, ok := DialercontactMap["contactableStatus"].(map[string]interface{}); ok {
		ContactableStatusString, _ := json.Marshal(ContactableStatus)
		json.Unmarshal(ContactableStatusString, &o.ContactableStatus)
	}
	
	if ContactColumnTimeZones, ok := DialercontactMap["contactColumnTimeZones"].(map[string]interface{}); ok {
		ContactColumnTimeZonesString, _ := json.Marshal(ContactColumnTimeZones)
		json.Unmarshal(ContactColumnTimeZonesString, &o.ContactColumnTimeZones)
	}
	
	if ConfigurationOverrides, ok := DialercontactMap["configurationOverrides"].(map[string]interface{}); ok {
		ConfigurationOverridesString, _ := json.Marshal(ConfigurationOverrides)
		json.Unmarshal(ConfigurationOverridesString, &o.ConfigurationOverrides)
	}
	
	if dateCreatedString, ok := DialercontactMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if SelfUri, ok := DialercontactMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
