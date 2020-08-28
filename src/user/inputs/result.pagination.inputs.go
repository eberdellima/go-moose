package inputs

// Paginator structure representing "from" which index
// data should be counted and what "size" (number) the data
// should have. Is used when querying data
type Paginator struct {
	From uint `form:"from" binding:"required"`
	Size uint `form:"size" binding:"required"`
}
