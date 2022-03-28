package modelssql

type AuthorsBook struct {
	Id      int    `gorm:"size:30;AUTO_INCREMENT;PRIMARY_KEY;unique_index"`
	Name    string `gorm:"type:varchar(128)"`
	Age     int    `gorm:"type:int"`
	Website string `gorm:"type:varchar(255)"`
	Email   string `gorm:"type:varchar(128)"`
}

type Books struct {
	Id          int    `gorm:"size:30;AUTO_INCREMENT;PRIMARY_KEY;unique_index"`
	IdAuthor    int    `gorm:"type:int"`
	NameBook    string `gorm:"type:varchar(128)"`
	NumberPages int    `gorm:"type:int"`
	Zhanr       string `gorm:"type:varchar(128)"`
}

type Zhanrs struct {
	Id        int    `gorm:"size:30;AUTO_INCREMENT;PRIMARY_KEY;unique_index"`
	NameZhanr string `gorm:"type:varchar(128)"`
}
