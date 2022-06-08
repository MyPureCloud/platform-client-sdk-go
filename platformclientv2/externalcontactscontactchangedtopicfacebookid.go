package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopicfacebookid
type Externalcontactscontactchangedtopicfacebookid struct { 
	// Ids
	Ids *[]Externalcontactscontactchangedtopicfacebookscopedid `json:"ids,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Externalcontactscontactchangedtopicfacebookid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopicfacebookid
	
	return json.Marshal(&struct { 
		Ids *[]Externalcontactscontactchangedtopicfacebookscopedid `json:"ids,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopicfacebookid) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopicfacebookidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopicfacebookidMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := ExternalcontactscontactchangedtopicfacebookidMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if DisplayName, ok := ExternalcontactscontactchangedtopicfacebookidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopicfacebookid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
