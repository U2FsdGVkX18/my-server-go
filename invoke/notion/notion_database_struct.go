package notion

// JsonRootBean 根节点
type JsonRootBean struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

// Parent 主节点
type Parent struct {
	DatabaseID string `json:"database_id"`
	Type       string `json:"type"`
}

// Properties 主节点
type Properties struct {
	ID             ID       `json:"ID"`
	Name           Title    `json:"Name"`
	Img            Img      `json:"Img"`
	ImgUrl         ImgUrl   `json:"ImgUrl"`
	Details        Details  `json:"Details"`
	DataCreateTime RichText `json:"DataCreateTime"`
}

type ID struct {
	Number int `json:"number"`
}

type Title struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type Img struct {
	Files []File `json:"files"`
	Type  string `json:"type"`
}

type File struct {
	External External `json:"external"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
}

type External struct {
	URL string `json:"url"`
}

type ImgUrl struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Details struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type RichText struct {
	Text Text   `json:"text"`
	Type string `json:"type"`
}
