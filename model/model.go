package model

type Episode struct {
	Start   string `json:"start"`
	Title   string `json:"title"`
	Episode struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	} `json:"episode"`
}

type epg struct {
	Episodes []Episode `json:"timelines"`
}

type Pluto struct {
	EPG []epg `json:"EPG"`
}
