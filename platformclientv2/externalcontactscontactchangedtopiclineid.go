package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopiclineid
type Externalcontactscontactchangedtopiclineid struct { 
	// Ids
	Ids *[]Externalcontactscontactchangedtopiclineuserid `json:"ids,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Externalcontactscontactchangedtopiclineid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopiclineid
	
	return json.Marshal(&struct { 
		Ids *[]Externalcontactscontactchangedtopiclineuserid `json:"ids,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopiclineid) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopiclineidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopiclineidMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := ExternalcontactscontactchangedtopiclineidMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if DisplayName, ok := ExternalcontactscontactchangedtopiclineidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopiclineid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
