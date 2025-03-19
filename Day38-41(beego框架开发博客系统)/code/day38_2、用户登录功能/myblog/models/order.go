package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id           int       `orm:"column(id);auto"`
	BrandId      uint      `orm:"column(brand_id)" description:"品牌id"`
	BrandCode    string    `orm:"column(brand_code);size(20)" description:"品牌编码"`
	BranchCode   string    `orm:"column(branch_code);size(20)" description:"门店code"`
	BranchName   string    `orm:"column(branch_name);size(255)" description:"门店名称"`
	OrderNo      string    `orm:"column(order_no);size(20)" description:"订单号"`
	UserId       uint      `orm:"column(user_id)" description:"用户id"`
	MemberId     uint      `orm:"column(member_id)" description:"会员id"`
	PeopleNum    uint      `orm:"column(people_num)" description:"用餐人数"`
	Name         string    `orm:"column(name);size(20)" description:"名字"`
	Sex          uint8     `orm:"column(sex)" description:"性别0未设置 1男 2女"`
	Seat         int8      `orm:"column(seat)" description:"订座位置 0不显示 1内场 2外摆"`
	DiningTime   time.Time `orm:"column(dining_time);type(datetime)" description:"就餐时间"`
	Status       uint8     `orm:"column(status)" description:"状态 1待确认 2已确认 3已取消 4已就餐"`
	Mobile       string    `orm:"column(mobile);size(20)" description:"手机号"`
	Remark       string    `orm:"column(remark);null" description:"备注"`
	WaiterRemark string    `orm:"column(waiter_remark);null" description:"店员备注"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);null"`
}

func (t *Order) TableName() string {
	return "order"
}

func init() {
	orm.RegisterModel(new(Order))
}

// AddOrder insert a new Order into database and returns
// last inserted Id on success.
func AddOrder(m *Order) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderById retrieves Order by Id. Returns error if
// Id doesn't exist
func GetOrderById(id int) (v *Order, err error) {
	o := orm.NewOrm()
	v = &Order{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrder retrieves all Order matches certain condition. Returns empty list if
// no records exist
func GetAllOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Order))
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

	var l []Order
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

// UpdateOrder updates Order by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderById(m *Order) (err error) {
	o := orm.NewOrm()
	v := Order{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrder deletes Order by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrder(id int) (err error) {
	o := orm.NewOrm()
	v := Order{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Order{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
