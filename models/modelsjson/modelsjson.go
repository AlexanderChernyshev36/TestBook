package modelsjson

type AuthorsBook struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Website string `json:"website"`
	Email   string `json:"email"`
}

type Books struct {
	Id          int    `json:"id"`
	IdAuthor    int    `json:"id_author"`
	NameBook    string `json:"name_book"`
	NumberPages int    `json:"number_pages"`
	Zhanr       int    `json:"zhanr"`
}

type Zhanrs struct {
	Id        int    `json:"id"`
	NameZhanr string `json:"name_zhanr"`
}
