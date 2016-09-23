package web

import "github.com/tedsuo/rata"

const (
	Index                 = "Index"
	RobotsTxt             = "RobotsTxt"
	Pipeline              = "Pipeline"
	GetBuild              = "GetBuild"
	GetBuilds             = "GetBuilds"
	GetJoblessBuild       = "GetJoblessBuild"
	Public                = "Public"
	GetResource           = "GetResource"
	GetJob                = "GetJob"
	LogIn                 = "LogIn"
	TeamLogIn             = "TeamLogIn"
	ProcessBasicAuthLogIn = "ProcessBasicAuthLogIn"

	// so links don't break
	MainPipeline    = "MainPipeline"
	MainGetJob      = "MainGetJob"
	MainGetResource = "MainGetResource"
	MainGetBuild    = "MainGetBuild"
)

var Routes = rata.Routes{
	// public
	{Path: "/", Method: "GET", Name: Index},
	{Path: "/robots.txt", Method: "GET", Name: RobotsTxt},

	// --- delete from here down

	{Path: "/teams/:team_name/pipelines/:pipeline", Method: "GET", Name: Pipeline},
	{Path: "/teams/:team_name/pipelines/:pipeline_name/jobs/:job", Method: "GET", Name: GetJob},
	{Path: "/teams/:team_name/pipelines/:pipeline_name/resources/:resource", Method: "GET", Name: GetResource},

	{Path: "/public/:filename", Method: "GET", Name: Public},
	{Path: "/public/fonts/:filename", Method: "GET", Name: Public},
	{Path: "/public/favicons/:filename", Method: "GET", Name: Public},
	{Path: "/public/images/:filename", Method: "GET", Name: Public},

	// public jobs only
	{Path: "/teams/:team_name/pipelines/:pipeline_name/jobs/:job/builds/:build", Method: "GET", Name: GetBuild},

	// private
	{Path: "/builds", Method: "GET", Name: GetBuilds},
	{Path: "/builds/:build_id", Method: "GET", Name: GetJoblessBuild},

	//probably leave these in, but remove the above eventually
	{Path: "/pipelines/:pipeline", Method: "GET", Name: MainPipeline},
	{Path: "/pipelines/:pipeline_name/jobs/:job", Method: "GET", Name: MainGetJob},
	{Path: "/pipelines/:pipeline_name/resources/:resource", Method: "GET", Name: MainGetResource},
	{Path: "/pipelines/:pipeline_name/jobs/:job/builds/:build", Method: "GET", Name: MainGetBuild},
}
