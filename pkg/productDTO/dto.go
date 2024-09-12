package productDTO

type Product struct {
	Name  string `json:"name" binding:"required"`
	Brand string `json:"brand" binding:"required"`
}
