package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Additionallanguagesintent
type Additionallanguagesintent struct { 
	// Id - ID of the intent for respective additional language
	Id *string `json:"id,omitempty"`


	// Utterances - Utterances list for additional language
	Utterances *[]Nluutterance `json:"utterances,omitempty"`

}

func (o *Additionallanguagesintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Additionallanguagesintent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Utterances *[]Nluutterance `json:"utterances,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Utterances: o.Utterances,
		Alias:    (*Alias)(o),
	})
}

func (o *Additionallanguagesintent) UnmarshalJSON(b []byte) error {
	var AdditionallanguagesintentMap map[string]interface{}
	err := json.Unmarshal(b, &AdditionallanguagesintentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdditionallanguagesintentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Utterances, ok := AdditionallanguagesintentMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Additionallanguagesintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
