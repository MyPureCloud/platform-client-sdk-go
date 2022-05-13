package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Commoncampaigndivisionview
type Commoncampaigndivisionview struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Campaign.
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// MediaType - The media type used for this campaign.
	MediaType *string `json:"mediaType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Commoncampaigndivisionview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Commoncampaigndivisionview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		MediaType: o.MediaType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Commoncampaigndivisionview) UnmarshalJSON(b []byte) error {
	var CommoncampaigndivisionviewMap map[string]interface{}
	err := json.Unmarshal(b, &CommoncampaigndivisionviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CommoncampaigndivisionviewMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CommoncampaigndivisionviewMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := CommoncampaigndivisionviewMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if MediaType, ok := CommoncampaigndivisionviewMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if SelfUri, ok := CommoncampaigndivisionviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Commoncampaigndivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
