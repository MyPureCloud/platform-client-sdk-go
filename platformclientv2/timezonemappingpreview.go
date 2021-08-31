package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timezonemappingpreview
type Timezonemappingpreview struct { 
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

func (o *Timezonemappingpreview) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
