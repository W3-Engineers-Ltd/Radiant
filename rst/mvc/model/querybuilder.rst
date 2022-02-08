Query Builder
=============

**QueryBuilder** provides an API for convenient and fluent construction
of SQL queries. It consists of a set of methods enabling developers to
easily construct SQL queries without compromising readability.

It serves as an alternative to ORM. ORM is more for simple CRUD
operations, whereas QueryBuilder is for complex queries with subqueries
and multi-joins.

Usage example:

.. code:: go

   // User is a wrapper for result row in this example
   type User struct {
       Name string
       Age  int
   }
   var users []User

   // Get a QueryBuilder object. Takes DB driver name as parameter
   // Second return value is error, ignored here
   qb, _ := orm.NewQueryBuilder("mysql")

   // Construct query object
   qb.Select("user.name",
       "profile.age").
       From("user").
       InnerJoin("profile").On("user.id_user = profile.fk_user").
       Where("age > ?").
       OrderBy("name").Desc().
       Limit(10).Offset(0)

   // export raw query string from QueryBuilder object
   sql := qb.String()

   // execute the raw query string
   o := orm.NewOrm()
   o.Raw(sql, 20).QueryRows(&users)

Full API interface:

.. code:: go

   type QueryBuilder interface {
       Select(fields ...string) QueryBuilder
       ForUpdate() QueryBuilder
       From(tables ...string) QueryBuilder
       InnerJoin(table string) QueryBuilder
       LeftJoin(table string) QueryBuilder
       RightJoin(table string) QueryBuilder
       On(cond string) QueryBuilder
       Where(cond string) QueryBuilder
       And(cond string) QueryBuilder
       Or(cond string) QueryBuilder
       In(vals ...string) QueryBuilder
       OrderBy(fields ...string) QueryBuilder
       Asc() QueryBuilder
       Desc() QueryBuilder
       Limit(limit int) QueryBuilder
       Offset(offset int) QueryBuilder
       GroupBy(fields ...string) QueryBuilder
       Having(cond string) QueryBuilder
       Update(tables ...string) QueryBuilder
       Set(kv ...string) QueryBuilder
       Delete(tables ...string) QueryBuilder
       InsertInto(table string, fields ...string) QueryBuilder
       Values(vals ...string) QueryBuilder
       Subquery(sub string, alias string) string
       String() string
   }

Now we support ``Postgress``, ``MySQL`` and ``TiDB``\ 。

.. toctree::
   :maxdepth: 4
   :caption: Contents:

   rst/quickstart
   rst/quickstart/README

.. toctree::
   :maxdepth: 4
   :caption: Quickstart Introduction:


   rst/quickstart/new
   rst/quickstart/router
   rst/quickstart/controller
   rst/quickstart/model
   rst/quickstart/view

.. toctree::
   :maxdepth: 4
   :caption: Introduction:

   rst/intro/Introduction
   rst/advantage/README
   rst/install/install
   rst/install/radical

.. toctree::
   :maxdepth: 4
   :caption: MVC Introduction:

   rst/mvc/README
.. toctree::
   :maxdepth: 4
   :caption: Controller:

   rst/mvc/controller/config
   rst/mvc/controller/controller
   rst/mvc/controller/error
   rst/mvc/controller/filter
   rst/mvc/controller/flash
   rst/mvc/controller/jsonxml
   rst/mvc/controller/params
   rst/mvc/controller/router
   rst/mvc/controller/session
   rst/mvc/controller/urlbuilding
   rst/mvc/controller/validation
   rst/mvc/controller/xsrf
.. toctree::
   :maxdepth: 4
   :caption: Models:

   rst/mvc/model/overview
   rst/mvc/model/orm
   rst/mvc/model/object
   rst/mvc/model/query
   rst/mvc/model/rawsql
   rst/mvc/model/querybuilder
   rst/mvc/model/transaction
   rst/mvc/model/models
   rst/mvc/model/commandline
   rst/mvc/model/test
   rst/mvc/model/custome_fields
   rst/mvc/model/faq

.. toctree::
   :maxdepth: 4
   :caption: Views:

   rst/mvc/view/view
   rst/mvc/view/template
   rst/mvc/view/static
   rst/mvc/view/page
   rst/mvc/view/global_variables

.. toctree::
   :maxdepth: 4
   :caption: Modules:

   rst/module/README
   rst/module/session
   rst/module/cache
   rst/module/logs
   rst/module/httplib
   rst/module/context
   rst/module/task
   rst/module/config

.. toctree::
   :maxdepth: 4
   :caption: Advanced Radiant:

   rst/advantage/README
   rst/advantage/monitor
   rst/advantage/docs

.. toctree::
   :maxdepth: 4
   :caption: Deployment:

   rst/deploy/README
   rst/deploy/radiant
   rst/deploy/supervisor
   rst/deploy/systemctl
   rst/deploy/nginx
   rst/deploy/apache

.. toctree::
   :maxdepth: 4
   :caption: FAQ:

   rst/faq/FAQ