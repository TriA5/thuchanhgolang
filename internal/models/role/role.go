package role

type Role string

const (
	MANAGER            Role = "MANAGER"
	REGION_MANAGER     Role = "REGION_MANAGER"
	BRANCH_MANAGER     Role = "BRANCH_MANAGER"
	HEAD_OF_DEPARTMENT Role = "HEAD_OF_DEPARTMENT"
	EMPLOYEE           Role = "EMPLOYEE"
)
