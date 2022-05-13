package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Minerintent
type Minerintent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Intent name.
	Name *string `json:"name,omitempty"`


	// Miner - The miner to which the intent belongs.
	Miner *Miner `json:"miner,omitempty"`


	// Utterances - The utterances that are extracted for an Intent.
	Utterances *[]Utterance `json:"utterances,omitempty"`


	// AnalyticVolumePercent - Percentage of conversations belonging to the intent.
	AnalyticVolumePercent *float64 `json:"analyticVolumePercent,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Minerintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minerintent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Miner *Miner `json:"miner,omitempty"`
		
		Utterances *[]Utterance `json:"utterances,omitempty"`
		
		AnalyticVolumePercent *float64 `json:"analyticVolumePercent,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Miner: o.Miner,
		
		Utterances: o.Utterances,
		
		AnalyticVolumePercent: o.AnalyticVolumePercent,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Minerintent) UnmarshalJSON(b []byte) error {
	var MinerintentMap map[string]interface{}
	err := json.Unmarshal(b, &MinerintentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MinerintentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MinerintentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Miner, ok := MinerintentMap["miner"].(map[string]interface{}); ok {
		MinerString, _ := json.Marshal(Miner)
		json.Unmarshal(MinerString, &o.Miner)
	}
	
	if Utterances, ok := MinerintentMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	
	if AnalyticVolumePercent, ok := MinerintentMap["analyticVolumePercent"].(float64); ok {
		o.AnalyticVolumePercent = &AnalyticVolumePercent
	}
    
	if SelfUri, ok := MinerintentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Minerintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
