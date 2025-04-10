// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/sena_2824182/Livestock_web_CRUD/REGISTRO_API/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/tipo_usuario",
			beego.NSInclude(
				&controllers.TipoUsuarioController{},
			),
		),

		beego.NSNamespace("/registro_usuario",
			beego.NSInclude(
				&controllers.RegistroUsuarioController{},
			),
		),

		beego.NSNamespace("/credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
