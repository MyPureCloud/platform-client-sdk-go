package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopicfacebookid
type Externalcontactsunresolvedcontactchangedtopicfacebookid struct { 
	// Ids
	Ids *[]Externalcontactsunresolvedcontactchangedtopicfacebookscopedid `json:"ids,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopicfacebookid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopicfacebookid
	
	return json.Marshal(&struct { 
		Ids *[]Externalcontactsunresolvedcontactchangedtopicfacebookscopedid `json:"ids,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopicfacebookid) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopicfacebookidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopicfacebookidMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := ExternalcontactsunresolvedcontactchangedtopicfacebookidMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if DisplayName, ok := ExternalcontactsunresolvedcontactchangedtopicfacebookidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopicfacebookid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
