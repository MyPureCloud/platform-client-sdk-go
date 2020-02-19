package platformclientv2
import (
	"encoding/json"
)

// Timezonemappingpreview
type Timezonemappingpreview struct { 
	// ContactList - The associated ContactList
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// ContactsPerTimeZone - The number of contacts per time zone that mapped to only that time zone
	ContactsPerTimeZone *map[string]int64 `json:"contactsPerTimeZone,omitempty"`


	// ContactsMappedUsingZipCode - The number of contacts per time zone that mapped to only that time zone and were mapped using the zip code column
	ContactsMappedUsingZipCode *map[string]int64 `json:"contactsMappedUsingZipCode,omitempty"`


	// ContactsMappedToASingleZone - The total number of contacts that mapped to a single time zone
	ContactsMappedToASingleZone *int64 `json:"contactsMappedToASingleZone,omitempty"`


	// ContactsMappedToASingleZoneUsingZipCode - The total number of contacts that mapped to a single time zone and were mapped using the zip code column
	ContactsMappedToASingleZoneUsingZipCode *int64 `json:"contactsMappedToASingleZoneUsingZipCode,omitempty"`


	// ContactsMappedToMultipleZones - The total number of contacts that mapped to multiple time zones
	ContactsMappedToMultipleZones *int64 `json:"contactsMappedToMultipleZones,omitempty"`


	// ContactsMappedToMultipleZonesUsingZipCode - The total number of contacts that mapped to multiple time zones and were mapped using the zip code column
	ContactsMappedToMultipleZonesUsingZipCode *int64 `json:"contactsMappedToMultipleZonesUsingZipCode,omitempty"`


	// ContactsInDefaultWindow - The total number of contacts that will be dialed during the default window
	ContactsInDefaultWindow *int64 `json:"contactsInDefaultWindow,omitempty"`


	// ContactListSize - The total number of contacts in the contact list
	ContactListSize *int64 `json:"contactListSize,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timezonemappingpreview) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
