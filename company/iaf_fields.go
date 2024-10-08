package company

type InternshipDescription struct {
	Profile          string           `json:"profile" bson:"profile"`
	Description      string           `json:"jd" bson:"jd"`
	MinimumHires     int              `json:"minimumHires" bson:"minimumHires"`
	ExpectedHires    int              `json:"expectedHires" bson:"expectedHires"`
	DurationInWeeks  int              `json:"internshipDuration" bson:"internshipDuration"`
	Locations        string           `json:"locations" bson:"locations"`
	IsCGPACriteria   bool             `json:"isCGPACriteria" bson:"isCGPACriteria"`
	EligibleCourses  []string         `json:"courses" bson:"courses"`
	BTech            IAFCourseDetails `json:"btech" bson:"btech"`
	IDD              IAFCourseDetails `json:"idd" bson:"idd"`
	MTech            IAFCourseDetails `json:"mtech" bson:"mtech"`
	PHd              IAFCourseDetails `json:"phd" bson:"phd"`
	BArch            IAFCourseDetails `json:"barch" bson:"barch"`
	JDAttachmentLink []string         `json:"jdAttachments" bson:"jdAttachments"`
}

type InternshipDescriptionCSV struct {
	Profile          string              `json:"profile" bson:"profile" csv:"Profile"`
	Description      string              `json:"jd" bson:"jd" csv:"Description"`
	MinimumHires     string              `json:"minimumHires" bson:"minimumHires" csv:"Minimum Hires"`
	ExpectedHires    string              `json:"expectedHires" bson:"expectedHires" csv:"Expected Hires"`
	DurationInWeeks  string              `json:"internshipDuration" bson:"internshipDuration" csv:"Duration In Weeks"`
	Locations        string              `json:"locations" bson:"locations" csv:"Locations"`
	IsCGPACriteria   string              `json:"isCGPACriteria" bson:"isCGPACriteria" csv:"Is there a CGPA Criteria"`
	EligibleCourses  string              `json:"courses" bson:"courses" csv:"Eligible Courses"`
	BTech            IAFCourseDetailsCSV `json:"btech" bson:"btech" csv:"BTech ,inline"`
	IDD              IAFCourseDetailsCSV `json:"idd" bson:"idd" csv:"IDD ,inline"`
	MTech            IAFCourseDetailsCSV `json:"mtech" bson:"mtech" csv:"MTech ,inline"`
	PHd              IAFCourseDetailsCSV `json:"phd" bson:"phd" csv:"PhD ,inline"`
	BArch            IAFCourseDetailsCSV `json:"barch" bson:"barch" csv:"BArch ,inline"`
	JDAttachmentLink string              `json:"jdAttachments" bson:"jdAttachments" csv:"JD links"`
}

type IAFCourseDetails struct {
	Eligibility EligibilityCriteria `json:"eligibility" bson:"eligibility"`
	Stipend     StipendDetails      `json:"compensationDetails" bson:"compensationDetails"`
}

type IAFCourseDetailsCSV struct {
	Eligibility EligibilityCriteriaCSV `json:"eligibility" bson:"eligibility" csv:"Eligibility ,inline"`
	Stipend     StipendDetailsCSV      `json:"compensationDetails" bson:"compensationDetails" csv:",inline"`
}

type StipendDetails struct {
	MonthlyStipend  string `json:"stipendPerMonth" bson:"stipendPerMonth" csv:"Monthly Stipend"`
	Accomodation    string `json:"accommodation" bson:"accommodation" csv:"Accomodation"`
	RelocationBonus string `json:"relocationBonus" bson:"relocationBonus" csv:"Relocation Bonus"`
	Incentives      string `json:"incentives" bson:"incentives" csv:"Incentives"`
	ProvisionForPPO bool   `json:"provisionForPPO" bson:"provisionForPPO" csv:"PPO Provision on performance"`
	TentativeCTC    string `json:"tentativeCTC" bson:"tentativeCTC" csv:"Tentative CTC on PPO"`
}

type StipendDetailsCSV struct {
	MonthlyStipend  string `json:"stipendPerMonth" bson:"stipendPerMonth" csv:"Monthly Stipend"`
	Accomodation    string `json:"accommodation" bson:"accommodation" csv:"Accomodation"`
	RelocationBonus string `json:"relocationBonus" bson:"relocationBonus" csv:"Relocation Bonus"`
	Incentives      string `json:"incentives" bson:"incentives" csv:"Incentives"`
	ProvisionForPPO string `json:"provisionForPPO" bson:"provisionForPPO" csv:"PPO Provision on performance"`
	TentativeCTC    string `json:"tentativeCTC" bson:"tentativeCTC" csv:"Tentative CTC on PPO"`
}
