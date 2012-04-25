package service

import (
	"github.com/timeredbull/tsuru/db"
	. "launchpad.net/gocheck"
	"launchpad.net/mgo/bson"
)

func (s *ServiceSuite) createServiceType() {
	s.serviceType = &ServiceType{Name: "Mysql", Charm: "mysql"}
	s.serviceType.Create()
}

func (s *ServiceSuite) TestAllServiceTypes(c *C) {
	st := ServiceType{Name: "Mysql", Charm: "mysql"}
	st2 := ServiceType{Name: "MongoDB", Charm: "mongodb"}
	st.Create()
	st2.Create()

	results := st.All()
	c.Assert(len(results), Equals, 2)
}

func (s *ServiceSuite) TestGetServiceType(c *C) {
	s.createServiceType()
	name := s.serviceType.Name
	charm := s.serviceType.Charm
	s.serviceType.Charm = ""
	s.serviceType.Name = ""
	s.serviceType.Get()

	c.Assert(s.serviceType.Name, Equals, name)
	c.Assert(s.serviceType.Charm, Equals, charm)
}

func (s *ServiceSuite) TestCreateServiceType(c *C) {
	s.createServiceType()
	query := bson.M{}
	result := ServiceType{}
	query["name"] = "Mysql"
	query["charm"] = "mysql"

	err := db.Session.ServiceTypes().Find(query).One(&result)
	c.Assert(err, IsNil)
	c.Assert(result.Name, Equals, "Mysql")
	c.Assert(result.Charm, Equals, "mysql")
}

func (s *ServiceSuite) TestDeleteServiceType(c *C) {
	s.createServiceType()
	s.serviceType.Delete()

	query := bson.M{}
	query["name"] = "Mysql"
	query["charm"] = "mysql"

	qtd, err := db.Session.ServiceTypes().Find(query).Count()
	c.Assert(err, IsNil)
	c.Assert(qtd, Equals, 0)
}
