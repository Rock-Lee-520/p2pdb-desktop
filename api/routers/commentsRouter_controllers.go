package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["api/controllers:MainController"] = append(beego.GlobalControllerRouter["api/controllers:MainController"],
        beego.ControllerComments{
            Method: "SendEmail",
            Router: "/send-email",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
