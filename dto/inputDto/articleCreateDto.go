package inputDto

type ArticleInputDto struct {
	TagId int `json:"tag_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	State int `json:"state"`
}