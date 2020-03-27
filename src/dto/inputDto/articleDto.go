package inputDto

type ArticleCreateInputDto struct {
	TagId int `json:"tag_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	State int `json:"state"`
}

type ArticleEditInputDto struct {
	Id int `json:"article_id"`
	TagId int `json:"tag_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	Modifier string `json:"modifier"`
	State int `json:"state"`
}

type ArticleSearchInputDto struct {

}