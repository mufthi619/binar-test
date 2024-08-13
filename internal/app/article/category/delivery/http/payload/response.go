package payload

type GetCategoryResponse struct {
	ArticleCategory []ArticleGroup `json:"article_category"`
}

type ArticleGroup struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
