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

type BookProperties struct {
	Properties  Properties
	ScorePeople ScorePeople `json:"ScorePeople"`
	Score       Score       `json:"Score"`
	Author      Author      `json:"Author"`
	OrigAuthor  OrigAuthor  `json:"OrigAuthor"`
	Translator  Translator  `json:"Translator"`
	Summary     Summary     `json:"Summary"`
	Kinds       Kinds       `json:"Kinds"`
	WordCount   WordCount   `json:"WordCount"`
	FixedPrice  FixedPrice  `json:"FixedPrice"`
	SalesPrice  SalesPrice  `json:"SalesPrice"`
}

type HighScoreTVShowRankingProperties struct {
	Properties  Properties
	ScorePeople ScorePeople `json:"ScorePeople"`
	Score       Score       `json:"Score"`
	Subtitle    Subtitle    `json:"Subtitle"`
	Type        Type        `json:"Type"`
	Year        Year        `json:"Year"`
	Tags        Tags        `json:"Tags"`
	HotComment  HotComment  `json:"HotComment"`
}

type MovieComingSoonProperties struct {
	Properties  Properties
	Region      Region      `json:"Region"`
	ReleaseDate ReleaseDate `json:"ReleaseDate"`
	Type        Type        `json:"Type"`
	WantToSee   WantToSee   `json:"WantToSee"`
}

type MovieNowShowingProperties struct {
	Properties  Properties
	ScorePeople ScorePeople `json:"ScorePeople"`
	Score       Score       `json:"Score"`
	Release     Release     `json:"Release"`
	Duration    Duration    `json:"Duration"`
	Region      Region      `json:"Region"`
	Director    Director    `json:"Director"`
	Actors      Actors      `json:"Actors"`
}

type NewMovieRankingProperties struct {
	Properties  Properties
	ScorePeople ScorePeople `json:"ScorePeople"`
	Score       Score       `json:"Score"`
	Intro       Intro       `json:"Intro"`
}

type Top250BookRankingProperties struct {
	Properties                  Properties
	ScorePeople                 ScorePeople                 `json:"ScorePeople"`
	Score                       Score                       `json:"Score"`
	AuthorPressPublicationPrice AuthorPressPublicationPrice `json:"AuthorPressPublicationPrice"`
	Quote                       Quote                       `json:"Quote"`
}

type Top250MovieRankingProperties struct {
	Properties           Properties
	ScorePeople          ScorePeople          `json:"ScorePeople"`
	Score                Score                `json:"Score"`
	DirectorAndActors    DirectorAndActors    `json:"DirectorAndActors"`
	YearAndRegionAndType YearAndRegionAndType `json:"YearAndRegionAndType"`
	Quote                Quote                `json:"Quote"`
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

type Actors struct {
	RichText []RichText `json:"rich_text"`
}

type Author struct {
	RichText []RichText `json:"rich_text"`
}

type AuthorPressPublicationPrice struct {
	RichText []RichText `json:"rich_text"`
}

type ScorePeople struct {
	RichText []RichText `json:"rich_text"`
}

type Score struct {
	RichText []RichText `json:"rich_text"`
}

type OrigAuthor struct {
	RichText []RichText `json:"rich_text"`
}

type Translator struct {
	RichText []RichText `json:"rich_text"`
}

type Summary struct {
	RichText []RichText `json:"rich_text"`
}

type Kinds struct {
	RichText []RichText `json:"rich_text"`
}

type WordCount struct {
	RichText []RichText `json:"rich_text"`
}

type FixedPrice struct {
	RichText []RichText `json:"rich_text"`
}

type SalesPrice struct {
	RichText []RichText `json:"rich_text"`
}

type DataCreateTime struct {
	RichText []RichText `json:"rich_text"`
}

type Director struct {
	RichText []RichText `json:"rich_text"`
}

type DirectorAndActors struct {
	RichText []RichText `json:"rich_text"`
}

type Duration struct {
	RichText []RichText `json:"rich_text"`
}

type HotComment struct {
	RichText []RichText `json:"rich_text"`
}

type Intro struct {
	RichText []RichText `json:"rich_text"`
}

type Name struct {
	RichText []RichText `json:"rich_text"`
}

type Quote struct {
	RichText []RichText `json:"rich_text"`
}

type Region struct {
	RichText []RichText `json:"rich_text"`
}

type Release struct {
	RichText []RichText `json:"rich_text"`
}

type ReleaseDate struct {
	RichText []RichText `json:"rich_text"`
}

type Subtitle struct {
	RichText []RichText `json:"rich_text"`
}

type Tags struct {
	RichText []RichText `json:"rich_text"`
}

type Type struct {
	RichText []RichText `json:"rich_text"`
}

type WantToSee struct {
	RichText []RichText `json:"rich_text"`
}

type Year struct {
	RichText []RichText `json:"rich_text"`
}

type YearAndRegionAndType struct {
	RichText []RichText `json:"rich_text"`
}
