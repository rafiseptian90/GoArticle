package requests

type Tabler interface {
	TableName() string
}

type TagRequest struct {
	Name string `json:"name" binding:"required"`
}

// TableName overrides the table name used by TagResponse to `tags`
func (TagRequest) TableName() string {
	return "tags"
}
