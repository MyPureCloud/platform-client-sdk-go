package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Timezonemappingpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
