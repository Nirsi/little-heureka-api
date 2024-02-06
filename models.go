package main

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Product struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"categoryId"`
	Title      string `json:"title"`
}

type Offer struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"productId"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	ImgURL      string  `json:"imgUrl"`
	Price       float32 `json:"price"`
}

type PageInfo struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

type Response[T Category | Product | Offer] struct {
	PageInfo PageInfo `json:"pageInfo"`
	Data     []T      `json:"data"`
}
