package forum

type ForumPostContentType string
type ForumPostType string
type Notify string

const (
	FORUMPOSTCOLLECTION = "forumpost"

	// ForumPostContentType
	TEXT         ForumPostContentType = "TEXT"
	STUDENT_LIST ForumPostContentType = "STUDENT_LIST"

	// ForumPostType
	CUSTOM               ForumPostType = "CUSTOM"
	TEST_SCHEDULING      ForumPostType = "TEST_SCHEDULING"
	TEST_VENUE           ForumPostType = "TEST_VENUE"
	TEST_SHORTLIST       ForumPostType = "TEST_SHORTLIST"
	INTERVIEW_SCHEDULING ForumPostType = "INTERVIEW_SCHEDULING"
	INTERVIEW_VENUE      ForumPostType = "INTERVIEW_VENUE"
	INTERVIEW_SHORTLIST  ForumPostType = "INTERVIEW_SHORTLIST"
	OFFERS               ForumPostType = "OFFERS"
	ANNOUNCEMENT         ForumPostType = "ANNOUNCEMENT"
	OFFCAMPUS            ForumPostType = "OFFCAMPUS"
	INTERVIEW_EXPERIENCE ForumPostType = "INTERVIEW_EXPERIENCE"
	TEST_EXPERIENCE      ForumPostType = "TEST_EXPERIENCE"
	OTHER                ForumPostType = "OTHER"

	// Notify
	ALL Notify = "ALL"
	TPR Notify = "TPR"
)
