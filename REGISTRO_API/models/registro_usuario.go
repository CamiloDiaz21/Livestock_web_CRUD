package models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RegistroUsuario struct {
	Id                int            `orm:"column(id);pk;auto"`
	Nombre            string         `orm:"column(nombre)"`
	Apellido          string         `orm:"column(apellido)"`
	FNacimiento       string         `orm:"column(f_nacimiento);null"`
	IdTipoDocumento   *TipoDocumento `orm:"column(id_tipo_documento);rel(fk)"`
	NDocumento        string         `orm:"column(n_documento)"`
	CorreoElectronico string         `orm:"column(correo_electronico)"`
	Celular           string         `orm:"column(celular)"`
	IdTipoUsuario     *TipoUsuario   `orm:"column(id_tipo_usuario);rel(fk)"`
	Contrasena        *Credenciales  `orm:"column(contrasena);rel(fk)"`
	Activo            bool           `orm:"column(activo)"`
	FCreacion         time.Time      `orm:"column(f_creacion);type(timestamp with time zone);auto_now_add"`
	FModificacion     time.Time      `orm:"column(f_modificacion);type(timestamp with time zone);auto_now"`
	FotoPerfil        string         `orm:"column(foto_perfil);type(text),null"`
}

func (t *RegistroUsuario) TableName() string {
	return "registro_usuario"
}

func init() {
	orm.RegisterModel(new(RegistroUsuario))
}

// AddRegistroUsuario insert a new RegistroUsuario into database and returns
// last inserted Id on success.
func AddRegistroUsuario(m *RegistroUsuario) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	id, err = o.Insert(m)
	return
}

// GetRegistroUsuarioById retrieves RegistroUsuario by Id. Returns error if
// Id doesn't exist
func GetRegistroUsuarioById(id int) (v *RegistroUsuario, err error) {
	o := orm.NewOrm()
	v = &RegistroUsuario{Id: id}
	if err = o.QueryTable(new(RegistroUsuario)).RelatedSel().Filter("Id", id).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroUsuario retrieves all RegistroUsuario matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroUsuario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroUsuario)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RegistroUsuario
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRegistroUsuario updates RegistroUsuario by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroUsuarioById(m *RegistroUsuario) (err error) {
	o := orm.NewOrm()
	v := RegistroUsuario{Id: m.Id}

	if err = o.Read(&v); err == nil {
		// Si viene una imagen nueva en base64
		if m.FotoPerfil != "" && strings.HasPrefix(m.FotoPerfil, "data:image") {
			// Separar encabezado de los datos
			dataIndex := strings.Index(m.FotoPerfil, ",")
			if dataIndex >= 0 {
				imageData := m.FotoPerfil[dataIndex+1:]

				// Decodificar base64
				decoded, decodeErr := base64.StdEncoding.DecodeString(imageData)
				if decodeErr != nil {
					return fmt.Errorf("Error al decodificar imagen: %v", decodeErr)
				}

				// Guardar imagen en disco (puedes usar el ID o correo como nombre)
				imagePath := fmt.Sprintf("static/uploads/user_%d.png", m.Id)
				writeErr := os.WriteFile(imagePath, decoded, 0644)
				if writeErr != nil {
					return fmt.Errorf("Error al guardar imagen: %v", writeErr)
				}

				// Guardar la ruta en la base de datos (opcional)
				m.FotoPerfil = imagePath
			}
		}

		// Actualizar los datos
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("✔️ Número de registros actualizados en la base de datos:", num)
		}
	}

	return
}

// DeleteRegistroUsuario deletes RegistroUsuario by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroUsuario(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroUsuario{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroUsuario{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
