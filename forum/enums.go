package forum

type ForumPostContentType string
type ForumPostType string
type Notify string

const (
	TEXT         ForumPostContentType = "TEXT"
	STUDENT_LIST ForumPostContentType = "STUDENT_LIST"

	CUSTOM         ForumPostType = "CUSTOM"
	TEST_SHORTLIST ForumPostType = "TEST_SHORTLIST"

	ALL Notify = "ALL"
	TPR Notify = "TPR"
)
