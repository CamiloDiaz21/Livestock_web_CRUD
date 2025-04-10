package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:RegistroUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers:TipoUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
