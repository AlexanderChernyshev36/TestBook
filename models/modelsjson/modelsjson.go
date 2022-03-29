package modelsjson

type QuerySaveBook struct {
	NameList1  string `json:"name_list_1"`
	NameList2  string `json:"name_list_2"`
	NameList3  string `json:"name_list_3"`
	PrintList1 bool   `json:"print_list_1"`
	PrintList2 bool   `json:"print_list_2"`
	PrintList3 bool   `json:"print_list_3"`
}

type ErrorStruct struct {
	Code      int    `json:"code"`
	TextError string `json:"text_error"`
	Error     bool   `json:"error"`
}

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
	Zhanr       string `json:"zhanr"`
}

type Zhanrs struct {
	Id        int    `json:"id"`
	NameZhanr string `json:"name_zhanr"`
}

type Result_Zhanrs struct {
	Data  []Zhanrs    `json:"data"`
	Error ErrorStruct `json:"error"`
}

type Result_AuthorsBook struct {
	Data  AuthorsBook `json:"data"`
	Error ErrorStruct `json:"error"`
}

type Result_Books struct {
	Data  Books       `json:"data"`
	Error ErrorStruct `json:"error"`
}
