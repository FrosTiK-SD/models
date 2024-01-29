package exam

import (
	"github.com/FrosTiK-SD/models/misc"
)

type StudentApplication struct {
	StudentID  misc.DBRef
	ResumeID   misc.DBRef
	AppliedFor misc.DBRef
}
