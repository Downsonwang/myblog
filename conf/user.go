/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-19 20:25:24
 * @LastEditTime: 2023-12-22 22:45:35
 */
 // DWs
package conf

type HomeInfo struct {
	SiteName string `json:"siteName"`

	Author string `json:"author"`

	Icp string `json:"icp"`

	TimeLayout string `json:"timeLayout"`

	Port int `json:"port"`

	WebHookSecret string `json:"webHookSecret"`

	CategoryDisplayQuantity int `json:"categoryDisplayQuantity"`

	TagDisplayQuantity int `json:"tagDisplayQuantity"`

	UtterancesRepo string `json:"utterancesRepo"`

	PageSize int `json:"pageSize"`

	DescriptionLen int `json:"descriptionLen"`

	DocumentGitUrl string `json:"documentGitUrl"`

	HtmlKeywords string `json:"htmlKeywords"`

	HtmlDescription string `json:"htmlDescription"`

	ThemeColor string `json:"themeColor"`

	ThemeOption []string `json:"themeOption"`

	DashboardEntrance string `json:"dashboardEntrance"`
}
