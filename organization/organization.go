package organization

type Organization struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OrganizationMember struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Member struct {
	OrgId  int64  `json:"org_id"`
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
}

func (m *Member) isRoleValid() bool {
	return m.Role == "Admin" || m.Role == "Developer" || m.Role == "Auditor"
}
