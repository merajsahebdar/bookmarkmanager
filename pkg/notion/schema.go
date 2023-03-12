package notion

// A Parent represents a link to a specific notion entity.
type Parent struct {
	Type       string `json:"type"`
	PageID     string `json:"page_id"`
	DatabaseID string `json:"database_id"`
	BlockID    string `json:"block_id"`
	Workspace  bool   `json:"workspace"`
}

// An Emoji represents the information about emoji character.
type Emoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

// A RichText represents the data that notion uses to display formatted text, mentions and inline equations.
type RichText struct {
	Type        string      `json:"type"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	HREF        string      `json:"href"`
	Equation    Equation    `json:"equation"`
	Mention     Mention     `json:"mention"`
	Text        Text        `json:"text"`
}

// An Annotations represents the styling for the rich text.
type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Underline     bool   `json:"underline"`
	Strikethrough bool   `json:"strikethrough"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

// An Equation represents an equation in a rich text.
type Equation struct {
	Expression string `json:"expression"`
}

// A Mention represents a mention in a rich text.
type Mention struct {
	Type     string          `json:"type"`
	Page     PageMention     `json:"page"`
	Database DatabaseMention `json:"database"`
}

// A Text represents a text in a rich text.
type Text struct {
	Content string `json:"content"`
	Link    Link   `json:"link"`
}

// A Link represents a link to a resource.
type Link struct {
	URL string `json:"url"`
}

// A PageMention represents a page in mention.
type PageMention struct {
	ID string `json:"id"`
}

// A DatabaseMention represents a database in mention.
type DatabaseMention struct {
	ID string `json:"id"`
}

// A SearchResultItem represents the item of a search response.
type SearchResultItem struct {
	Title []RichText `json:"title"`
}

// A SearchResponse represents the response of a search operation.
type SearchResponse struct {
	Results []SearchResultItem `json:"results"`
}
