package muddler

// Trigger represents the Muddler spec for a Mudlet Trigger
// See: https://github.com/demonnic/muddler/wiki/Triggers
type Trigger struct {
	Name           string    `json:"name"`
	IsActive       string    `json:"isActive"`
	IsFolder       string    `json:"isFolder"`
	Command        string    `json:"command,omitempty"`
	Multiline      string    `json:"multiline"`
	MultilineDelta string    `json:"multilineDelta"`
	Matchall       string    `json:"matchall"`
	Filter         string    `json:"filter"`
	FireLength     string    `json:"fireLength"`
	SoundFile      string    `json:"soundFile,omitempty"`
	Highlight      string    `json:"highlight"`
	HighlightFG    string    `json:"highlightFG,omitempty"`
	HighlightBG    string    `json:"highlightBG,omitempty"`
	Script         string    `json:"script"`
	Patterns       []Pattern `json:"patterns"`
	Children       []Trigger `json:"children,omitempty"`
}

type Pattern struct {
	Pattern string `json:"pattern"`
	Type    string `json:"type"`
}
