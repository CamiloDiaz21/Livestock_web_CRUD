package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:ComentarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/RESENAS_API/controllers:TipoPublicacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
