package constant

const DB = "portal"

const IAFCollection = "IAF"
const IAFDraftCollection = "IAFDraft"

const JAFCollection = "JAF"
const JAFDraftCollection = "JAFDraft"

const RecruiterCollection = "recruiters"
const CompanyCollection = "companies"
const AuthGroupCollection = "groups"
const CurrentAcademicSession = S_2023_2024

var GenericEmailProviders = []string{"gmail.com", "outlook.com", "protonmail.com", "tutanota.com", "yahoo.com"}

type AcademicSession string

const (
	S_2020_2021 AcademicSession = "2020-21"
	S_2021_2022 AcademicSession = "2021-22"
	S_2022_2023 AcademicSession = "2022-23"
	S_2023_2024 AcademicSession = "2023-24"
	S_2024_2025 AcademicSession = "2024-25"
	S_2025_2026 AcademicSession = "2025-26"
	S_2026_2027 AcademicSession = "2026-27"
	S_2027_2028 AcademicSession = "2027-28"
	S_2028_2029 AcademicSession = "2028-29"
	S_2029_2030 AcademicSession = "2029-30"
	S_2030_2031 AcademicSession = "2030-31"
)

const (
	BTECH = "btech"
	IDD   = "idd"
	MTECH = "mtech"
	PHD   = "phd"
)
