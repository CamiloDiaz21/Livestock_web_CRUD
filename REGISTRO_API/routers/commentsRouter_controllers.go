package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:IdentificacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioComunController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock_web_CRUD/REGISTRO_API/controllers:UsuarioVendedorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
