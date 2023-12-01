package company

type JobDescription struct {
	Profile          string           `json:"profile" bson:"profile"`
	Description      string           `json:"jd" bson:"jd"`
	ExpectedHires    int              `json:"expectedHires" bson:"expectedHires"`
	Locations        string           `json:"locations" bson:"locations"`
	EligibleCourses  []string         `json:"courses" bson:"courses"`
	BTech            JAFCourseDetails `json:"btech" bson:"btech"`
	IDD              JAFCourseDetails `json:"idd" bson:"idd"`
	MTech            JAFCourseDetails `json:"mtech" bson:"mtech"`
	PHd              JAFCourseDetails `json:"phd" bson:"phd"`
	JDAttachmentLink []string         `json:"jdAttachments" bson:"jdAttachments"`
}

type JobDescriptionCSV struct {
	Profile          string              `json:"profile" bson:"profile" csv:"Profile"`
	Description      string              `json:"jd" bson:"jd" csv:"Description"`
	ExpectedHires    string              `json:"expectedHires" bson:"expectedHires" csv:"Expected Hires"`
	Locations        string              `json:"locations" bson:"locations" csv:"Locations"`
	EligibleCourses  string              `json:"courses" bson:"courses" csv:"Eligible Courses"`
	BTech            JAFCourseDetailsCSV `json:"btech" bson:"btech" csv:"BTech ,inline"`
	IDD              JAFCourseDetailsCSV `json:"idd" bson:"idd" csv:"IDD ,inline"`
	MTech            JAFCourseDetailsCSV `json:"mtech" bson:"mtech" csv:"MTech ,inline"`
	PHd              JAFCourseDetailsCSV `json:"phd" bson:"phd" csv:"PhD ,inline"`
	JDAttachmentLink string              `json:"jdAttachments" bson:"JD Links"`
}

type JAFCourseDetails struct {
	Eligibility EligibilityCriteria `json:"eligibility" bson:"eligibility"`
	CTC         CTCDetails          `json:"compensationDetails" bson:"compensationDetails"`
}

type JAFCourseDetailsCSV struct {
	Eligibility EligibilityCriteriaCSV `json:"eligibility" bson:"eligibility" csv:"Eligibility ,inline"`
	CTC         CTCDetails             `json:"compensationDetails" bson:"compensationDetails" csv:",inline"`
}

type CTCDetails struct {
	TotalCTC                 string `json:"totalCTC" bson:"totalCTC" csv:"Total CTC"`
	FirstYearCTC             string `json:"firstYearCTC" bson:"firstYearCTC" csv:"First Year CTC"`
	FixedSalary              string `json:"fixedSalary" bson:"fixedSalary" csv:"Fixed Salary"`
	ProvidentFundAndGratuity string `json:"providentFundAndGratuity" bson:"providentFundAndGratuity" csv:"Provident Fund and Gratuity"`
	JoiningBonus             string `json:"joiningBonus" bson:"joiningBonus" csv:"Joining Bonus"`
	PatentBonus              string `json:"patentBonus" bson:"patentBonus" csv:"Patent Bonus"`
	RelocationBonus          string `json:"relocationBonus" bson:"relocationBonus" csv:"Relocation Bonus"`
	EducationReimbursement   string `json:"educationReimbursement" bson:"educationReimbursement" csv:"Education Reimbursement"`
	Bonds                    string `json:"bonds" bson:"bonds" csv:"Bonds"`
	StocksWithVestingPlan    string `json:"stocksWithVestingPlan" bson:"stocksWithVestingPlan" csv:"Stocks with Vesting Plan"`
	OtherBenefits            string `json:"otherBenefits" bson:"otherBenefits" csv:"Other Benefits"`
}
