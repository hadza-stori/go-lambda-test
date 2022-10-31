package entities

type Pokemon struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
}
