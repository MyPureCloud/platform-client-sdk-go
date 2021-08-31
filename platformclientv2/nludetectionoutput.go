package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectionoutput
type Nludetectionoutput struct { 
	// Intents - The detected intents.
	Intents *[]Detectedintent `json:"intents,omitempty"`


	// DialogActs - The detected dialog acts.
	DialogActs *[]Detecteddialogact `json:"dialogActs,omitempty"`

}

func (o *Nludetectionoutput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectionoutput
	
	return json.Marshal(&struct { 
		Intents *[]Detectedintent `json:"intents,omitempty"`
		
		DialogActs *[]Detecteddialogact `json:"dialogActs,omitempty"`
		*Alias
	}{ 
		Intents: o.Intents,
		
		DialogActs: o.DialogActs,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludetectionoutput) UnmarshalJSON(b []byte) error {
	var NludetectionoutputMap map[string]interface{}
	err := json.Unmarshal(b, &NludetectionoutputMap)
	if err != nil {
		return err
	}
	
	if Intents, ok := NludetectionoutputMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if DialogActs, ok := NludetectionoutputMap["dialogActs"].([]interface{}); ok {
		DialogActsString, _ := json.Marshal(DialogActs)
		json.Unmarshal(DialogActsString, &o.DialogActs)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nludetectionoutput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
