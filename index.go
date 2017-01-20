package decoder

type Pack struct {
	ID    ID `json:"id"`
	Blobs []struct {
		ID     ID     `json:"id"`
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Length int    `json:"length"`
	} `json:"blobs"`
}

type Index struct {
	Supersedes []string `json:"supersedes"`
	Packs      []Pack   `json:"packs"`
}
