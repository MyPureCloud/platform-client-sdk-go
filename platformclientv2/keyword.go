package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Keyword
type Keyword struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Phrase - The word or phrase which is being looked for with speech recognition.
	Phrase *string `json:"phrase,omitempty"`


	// Confidence - A sensitivity threshold that can be increased to lower false positives or decreased to reduce false negatives.
	Confidence *int `json:"confidence,omitempty"`


	// AgentScoreModifier - A modifier to the evaluation score when the phrase is spotted in the agent channel
	AgentScoreModifier *int `json:"agentScoreModifier,omitempty"`


	// CustomerScoreModifier - A modifier to the evaluation score when the phrase is spotted in the customer channel
	CustomerScoreModifier *int `json:"customerScoreModifier,omitempty"`


	// AlternateSpellings - Other spellings of the phrase that can be added to reduce missed spots (false negatives).
	AlternateSpellings *[]string `json:"alternateSpellings,omitempty"`


	// Pronunciations - The phonetic spellings for the phrase and alternate spellings.
	Pronunciations *[]string `json:"pronunciations,omitempty"`


	// AntiWords - Words that are similar to the phrase but not desired. Added to reduce incorrect spots (false positives).
	AntiWords *[]string `json:"antiWords,omitempty"`


	// AntiPronunciations - The phonetic spellings for the antiWords.
	AntiPronunciations *[]string `json:"antiPronunciations,omitempty"`


	// SpotabilityIndex - A prediction of how easy it is to unambiguously spot the keyword within its language based on spelling.
	SpotabilityIndex *float64 `json:"spotabilityIndex,omitempty"`


	// MarginOfError
	MarginOfError *float64 `json:"marginOfError,omitempty"`


	// Pronunciation
	Pronunciation *string `json:"pronunciation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Keyword) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
