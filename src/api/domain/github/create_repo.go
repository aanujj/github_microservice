package github

/*func CreateRepo(){
	request := map[string]interface{}{
		"name":"Hello Github From Gin",
		"private":false,
	}
}*/

type CreateRepoRequest struct {
	Name string 	`json:"name"`
	Description string `json:"description"`
	Homepage string `json:"homepage"`
	Private bool `json:"private"`
	HasIssues bool `json:"has_issues"`
	HasProjects bool `json:"has_projects"`
	HasWiki bool `json:"has_wiki"`
}