package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HaciendaLoteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:HistorialVentasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoGanadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/livestock/DATOS_VENTAS/controllers:TipoVentaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
