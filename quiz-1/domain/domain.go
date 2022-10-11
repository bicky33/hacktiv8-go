package domain

type Relation struct {
	InfluencedBy []string `json:"influenced-by,omitempty"`
	Influences   []string `json:"influences,omitempty"`
}

type JsonObject struct {
	Language       string   `json:"language,omitempty"`
	Appeared       int      `json:"appeared,omitempty"`
	Created        []string `json:"created,omitempty"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation,omitempty"`
}
