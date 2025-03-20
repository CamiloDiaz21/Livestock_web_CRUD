package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/sena_2824182/livestock/DATOS_VENTAS/models"

	"github.com/astaxie/beego"
)

// TipoGanadoController operations for TipoGanado
type TipoGanadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoGanadoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoGanado
// @Param	body		body 	models.TipoGanado	true		"body for TipoGanado content"
// @Success 201 {int} models.TipoGanado
// @Failure 403 body is empty
// @router / [post]
func (c *TipoGanadoController) Post() {
	var v models.TipoGanado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTipoGanado(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": 201, "Message": "Creado correctamente", "Datos creados con id": v}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get TipoGanado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoGanado
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoGanadoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoGanadoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Consulta correctamente", "Consulta de id": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TipoGanado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoGanado
// @Failure 403
// @router / [get]
func (c *TipoGanadoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllTipoGanado(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Consulta correctamente", "Consulta de id": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TipoGanado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoGanado	true		"body for TipoGanado content"
// @Success 200 {object} models.TipoGanado
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoGanadoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.TipoGanado{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTipoGanadoById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Actualizacion correcta", "Actualizacion de id": v}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the TipoGanado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoGanadoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoGanado(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "eliminado correctamente", "Dato eliminado con id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
