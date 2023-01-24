package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timezonemappingpreview
type Timezonemappingpreview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactList - The associated ContactList
	ContactList *Domainentityref `json:"contactList,omitempty"`

	// ContactsPerTimeZone - The number of contacts per time zone that mapped to only that time zone
	ContactsPerTimeZone *map[string]int `json:"contactsPerTimeZone,omitempty"`

	// ContactsMappedUsingZipCode - The number of contacts per time zone that mapped to only that time zone and were mapped using the zip code column
	ContactsMappedUsingZipCode *map[string]int `json:"contactsMappedUsingZipCode,omitempty"`

	// ContactsMappedToASingleZone - The total number of contacts that mapped to a single time zone
	ContactsMappedToASingleZone *int `json:"contactsMappedToASingleZone,omitempty"`

	// ContactsMappedToASingleZoneUsingZipCode - The total number of contacts that mapped to a single time zone and were mapped using the zip code column
	ContactsMappedToASingleZoneUsingZipCode *int `json:"contactsMappedToASingleZoneUsingZipCode,omitempty"`

	// ContactsMappedToMultipleZones - The total number of contacts that mapped to multiple time zones
	ContactsMappedToMultipleZones *int `json:"contactsMappedToMultipleZones,omitempty"`

	// ContactsMappedToMultipleZonesUsingZipCode - The total number of contacts that mapped to multiple time zones and were mapped using the zip code column
	ContactsMappedToMultipleZonesUsingZipCode *int `json:"contactsMappedToMultipleZonesUsingZipCode,omitempty"`

	// ContactsInDefaultWindow - The total number of contacts that will be dialed during the default window
	ContactsInDefaultWindow *int `json:"contactsInDefaultWindow,omitempty"`

	// ContactListSize - The total number of contacts in the contact list
	ContactListSize *int `json:"contactListSize,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timezonemappingpreview) SetField(field string, fieldValue interface{}) {
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

func (o Timezonemappingpreview) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Timezonemappingpreview
	
	return json.Marshal(&struct { 
		ContactList *Domainentityref `json:"contactList,omitempty"`
		
		ContactsPerTimeZone *map[string]int `json:"contactsPerTimeZone,omitempty"`
		
		ContactsMappedUsingZipCode *map[string]int `json:"contactsMappedUsingZipCode,omitempty"`
		
		ContactsMappedToASingleZone *int `json:"contactsMappedToASingleZone,omitempty"`
		
		ContactsMappedToASingleZoneUsingZipCode *int `json:"contactsMappedToASingleZoneUsingZipCode,omitempty"`
		
		ContactsMappedToMultipleZones *int `json:"contactsMappedToMultipleZones,omitempty"`
		
		ContactsMappedToMultipleZonesUsingZipCode *int `json:"contactsMappedToMultipleZonesUsingZipCode,omitempty"`
		
		ContactsInDefaultWindow *int `json:"contactsInDefaultWindow,omitempty"`
		
		ContactListSize *int `json:"contactListSize,omitempty"`
		Alias
	}{ 
		ContactList: o.ContactList,
		
		ContactsPerTimeZone: o.ContactsPerTimeZone,
		
		ContactsMappedUsingZipCode: o.ContactsMappedUsingZipCode,
		
		ContactsMappedToASingleZone: o.ContactsMappedToASingleZone,
		
		ContactsMappedToASingleZoneUsingZipCode: o.ContactsMappedToASingleZoneUsingZipCode,
		
		ContactsMappedToMultipleZones: o.ContactsMappedToMultipleZones,
		
		ContactsMappedToMultipleZonesUsingZipCode: o.ContactsMappedToMultipleZonesUsingZipCode,
		
		ContactsInDefaultWindow: o.ContactsInDefaultWindow,
		
		ContactListSize: o.ContactListSize,
		Alias:    (Alias)(o),
	})
}

func (o *Timezonemappingpreview) UnmarshalJSON(b []byte) error {
	var TimezonemappingpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &TimezonemappingpreviewMap)
	if err != nil {
		return err
	}
	
	if ContactList, ok := TimezonemappingpreviewMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if ContactsPerTimeZone, ok := TimezonemappingpreviewMap["contactsPerTimeZone"].(map[string]interface{}); ok {
		ContactsPerTimeZoneString, _ := json.Marshal(ContactsPerTimeZone)
		json.Unmarshal(ContactsPerTimeZoneString, &o.ContactsPerTimeZone)
	}
	
	if ContactsMappedUsingZipCode, ok := TimezonemappingpreviewMap["contactsMappedUsingZipCode"].(map[string]interface{}); ok {
		ContactsMappedUsingZipCodeString, _ := json.Marshal(ContactsMappedUsingZipCode)
		json.Unmarshal(ContactsMappedUsingZipCodeString, &o.ContactsMappedUsingZipCode)
	}
	
	if ContactsMappedToASingleZone, ok := TimezonemappingpreviewMap["contactsMappedToASingleZone"].(float64); ok {
		ContactsMappedToASingleZoneInt := int(ContactsMappedToASingleZone)
		o.ContactsMappedToASingleZone = &ContactsMappedToASingleZoneInt
	}
	
	if ContactsMappedToASingleZoneUsingZipCode, ok := TimezonemappingpreviewMap["contactsMappedToASingleZoneUsingZipCode"].(float64); ok {
		ContactsMappedToASingleZoneUsingZipCodeInt := int(ContactsMappedToASingleZoneUsingZipCode)
		o.ContactsMappedToASingleZoneUsingZipCode = &ContactsMappedToASingleZoneUsingZipCodeInt
	}
	
	if ContactsMappedToMultipleZones, ok := TimezonemappingpreviewMap["contactsMappedToMultipleZones"].(float64); ok {
		ContactsMappedToMultipleZonesInt := int(ContactsMappedToMultipleZones)
		o.ContactsMappedToMultipleZones = &ContactsMappedToMultipleZonesInt
	}
	
	if ContactsMappedToMultipleZonesUsingZipCode, ok := TimezonemappingpreviewMap["contactsMappedToMultipleZonesUsingZipCode"].(float64); ok {
		ContactsMappedToMultipleZonesUsingZipCodeInt := int(ContactsMappedToMultipleZonesUsingZipCode)
		o.ContactsMappedToMultipleZonesUsingZipCode = &ContactsMappedToMultipleZonesUsingZipCodeInt
	}
	
	if ContactsInDefaultWindow, ok := TimezonemappingpreviewMap["contactsInDefaultWindow"].(float64); ok {
		ContactsInDefaultWindowInt := int(ContactsInDefaultWindow)
		o.ContactsInDefaultWindow = &ContactsInDefaultWindowInt
	}
	
	if ContactListSize, ok := TimezonemappingpreviewMap["contactListSize"].(float64); ok {
		ContactListSizeInt := int(ContactListSize)
		o.ContactListSize = &ContactListSizeInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timezonemappingpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
