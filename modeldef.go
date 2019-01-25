package tableFormat

type Field struct {
	DataType string `json:"DataType"`
	Desc     string `json:"Desc"`
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Size     string `json:"Size"`
	Flag     string `json:"flag"`
	Index    string `json:"index"`
	Keygroup string `json:"keygroup"`
	Shift    string `json:"shift"`
}

type Tablex struct {
	Alias    string  `json:"Alias"`
	ID       string  `json:"ID"`
	Name     string  `json:"Name"`
	PriKey   string  `json:"PriKey"`
	Fieldnum string  `json:"fieldnum"`
	Fields   []Field `json:"fields"`
	Flag     string  `json:"flag"`
	Index    string  `json:"index"`
	Path     string  `json:"path"`
}

type DB struct {
	Alias string   `json:"Alias"`
	ID    string   `json:"ID"`
	Name  string   `json:"Name"`
	Table []Tablex `json:"table"`
	Type  string   `json:"type"`
}

type Settingx struct {
	Database []DB `json:"database"`
}

type result struct {
	Setting Settingx `json:"Setting"`
}

type replyCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ModelDef struct {
	ReplyCode replyCode `json:"replyCode"`
	Result    result    `json:"result"`
}
