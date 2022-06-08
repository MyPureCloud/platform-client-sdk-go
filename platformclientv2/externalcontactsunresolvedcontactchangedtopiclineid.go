package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopiclineid
type Externalcontactsunresolvedcontactchangedtopiclineid struct { 
	// Ids
	Ids *[]Externalcontactsunresolvedcontactchangedtopiclineuserid `json:"ids,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopiclineid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopiclineid
	
	return json.Marshal(&struct { 
		Ids *[]Externalcontactsunresolvedcontactchangedtopiclineuserid `json:"ids,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopiclineid) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopiclineidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopiclineidMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := ExternalcontactsunresolvedcontactchangedtopiclineidMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if DisplayName, ok := ExternalcontactsunresolvedcontactchangedtopiclineidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopiclineid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
