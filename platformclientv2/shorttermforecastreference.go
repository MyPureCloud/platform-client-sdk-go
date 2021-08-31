package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shorttermforecastreference - A pointer to a short term forecast
type Shorttermforecastreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - The description of the short term forecast
	Description *string `json:"description,omitempty"`

}

func (o *Shorttermforecastreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shorttermforecastreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		WeekDate: o.WeekDate,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Shorttermforecastreference) UnmarshalJSON(b []byte) error {
	var ShorttermforecastreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &ShorttermforecastreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ShorttermforecastreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SelfUri, ok := ShorttermforecastreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if WeekDate, ok := ShorttermforecastreferenceMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
	
	if Description, ok := ShorttermforecastreferenceMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
