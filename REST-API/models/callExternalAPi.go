package models

// Post represents the structure of a single post
type CallExternalApi struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	Reactions struct {
		Likes    int `json:"likes"`
		Dislikes int `json:"dislikes"`
	} `json:"reactions"`
	Views  int `json:"views"`
	UserID int `json:"userId"`
}

// PostsResponse represents the structure of the JSON response containing multiple posts
type CallExternalApiResponse struct {
	Posts []CallExternalApi `json:"posts"`
	Total int               `json:"total"`
	Skip  int               `json:"skip"`
	Limit int               `json:"limit"`
}
