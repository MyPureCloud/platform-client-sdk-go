package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonesreboot
type Phonesreboot struct { 
	// PhoneIds - The list of phone Ids to reboot.
	PhoneIds *[]string `json:"phoneIds,omitempty"`


	// SiteId - ID of the site for which to reboot all phones at that site. no.active.edge and phone.cannot.resolve errors are ignored.
	SiteId *string `json:"siteId,omitempty"`

}

func (o *Phonesreboot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonesreboot
	
	return json.Marshal(&struct { 
		PhoneIds *[]string `json:"phoneIds,omitempty"`
		
		SiteId *string `json:"siteId,omitempty"`
		*Alias
	}{ 
		PhoneIds: o.PhoneIds,
		
		SiteId: o.SiteId,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonesreboot) UnmarshalJSON(b []byte) error {
	var PhonesrebootMap map[string]interface{}
	err := json.Unmarshal(b, &PhonesrebootMap)
	if err != nil {
		return err
	}
	
	if PhoneIds, ok := PhonesrebootMap["phoneIds"].([]interface{}); ok {
		PhoneIdsString, _ := json.Marshal(PhoneIds)
		json.Unmarshal(PhoneIdsString, &o.PhoneIds)
	}
	
	if SiteId, ok := PhonesrebootMap["siteId"].(string); ok {
		o.SiteId = &SiteId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phonesreboot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
