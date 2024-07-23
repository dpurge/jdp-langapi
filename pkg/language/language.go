package language

import "strconv"

type Language struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func languages() []Language {
	return []Language{
		{
			ID:   1,
			Code: "eng",
			Name: "English",
		},
		{
			ID:   2,
			Code: "spa",
			Name: "Spanish",
		},
		{
			ID:   3,
			Code: "deu",
			Name: "German",
		},
		{
			ID:   4,
			Code: "fra",
			Name: "French",
		},
		{
			ID:   5,
			Code: "pol",
			Name: "Polish",
		},
	}
}

func loadLanguages() map[string]Language {
	languages := languages()
	res := make(map[string]Language, len(languages))

	for _, x := range languages {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}
