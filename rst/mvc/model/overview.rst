Models － Radiant ORM
=====================

|Build Status| |Go Walker|

Radiant ORM is a powerful ORM framework written in Go. It is inspired by
Django ORM and SQLAlchemy.

This framework is still under development so compatibility is not
guaranteed.

**Supported Database:**

-  MySQL：\ `github.com/go-sql-driver/mysql <https://github.com/go-sql-driver/mysql>`__
-  PostgreSQL：\ `github.com/lib/pq <https://github.com/lib/pq>`__
-  Sqlite3：\ `github.com/mattn/go-sqlite3 <https://github.com/mattn/go-sqlite3>`__

All of the database drivers have passed the tests, but we still need
your feedback and bug reports.

**ORM Features:**

-  Supports all the types in Go.
-  CRUD is easy to use.
-  Auto join connection tables.
-  Compatible with crossing database queries.
-  Supports raw SQL query and mapping.
-  Strict and well-covered test cases ensure the ORM’s stability.

You can learn more in this documentation.

**Install ORM:**

::

   go get github.com/W3-Engineers-Ltd/Radiant/client/orm

Quickstart
----------

Demo
~~~~

.. code:: go

   package main

   import (
       "fmt"
       "github.com/W3-Engineers-Ltd/Radiant/client/orm"
       _ "github.com/go-sql-driver/mysql" // import your required driver
   )

   // Model Struct
   type User struct {
       Id   int
       Name string `orm:"size(100)"`
   }

   func init() {
       // register model
       orm.RegisterModel(new(User))

       // set default database
       orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
   }

   func main() {
       o := orm.NewOrm()

       user := User{Name: "slene"}

       // insert
       id, err := o.Insert(&user)
       fmt.Printf("ID: %d, ERR: %v\n", id, err)

       // update
       user.Name = "astaxie"
       num, err := o.Update(&user)
       fmt.Printf("NUM: %d, ERR: %v\n", num, err)

       // read one
       u := User{Id: user.Id}
       err = o.Read(&u)
       fmt.Printf("ERR: %v\n", err)

       // delete
       num, err = o.Delete(&u)
       fmt.Printf("NUM: %d, ERR: %v\n", num, err)
   }

Relation Query
~~~~~~~~~~~~~~

.. code:: go

   type Post struct {
       Id    int    `orm:"auto"`
       Title string `orm:"size(100)"`
       User  *User  `orm:"rel(fk)"`
   }

   var posts []*Post
   qs := o.QueryTable("post")
   num, err := qs.Filter("User__Name", "slene").All(&posts)

Raw SQL query
~~~~~~~~~~~~~

You can always use raw SQL to query and mapping.

.. code:: go

   var maps []Params
   num, err := o.Raw("SELECT id FROM user WHERE name = ?", "slene").Values(&maps)
   if num > 0 {
       fmt.Println(maps[0]["id"])
   }

Transactions
~~~~~~~~~~~~

.. code:: go

   o.Begin()
   ...
   user := User{Name: "slene"}
   id, err := o.Insert(&user)
   if err == nil {
       o.Commit()
   } else {
       o.Rollback()
   }

Debugging query log
~~~~~~~~~~~~~~~~~~~

In development environment, you can enable debug mode by:

.. code:: go

   func main() {
       orm.Debug = true
   ...

It will output every query statement including execution, preparation
and transactions.

For example:

.. code:: go

   [ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] -   [INSERT INTO `user` (`name`) VALUES (?)] - `slene`
   ...

Notes: It is not recommended to enable debug mode in a production
environment.

Index
-----

1.  `Orm Usage <orm.html>`__

    -  `Set up database <orm.html#set-up-database>`__

       -  `Register Driver <orm.html#registerdatabase>`__
       -  `Variables Config <orm.html#setmaxidleconns>`__
       -  `Timezone Config <orm.html#timezone-config>`__

    -  `Registering Model <orm.html#registering-model>`__
    -  `ORM API Usage <orm.html#orm-api-usage>`__
    -  `Print Out SQL Query in Debugging
       Mode <orm.html#print-out-sql-query-in-debugging-mode>`__

2.  `CRUD of Object <object.html>`__
3.  `Advanced Queries <query.html>`__

    -  `expr <query.html#expr>`__
    -  `Operators <query.html#operators>`__
    -  `Advanced query API <query.html#advanced-query-api>`__
    -  `Relational Queries <query.html#relational-query>`__
    -  `Load Related Fields <query.html#load-related-field>`__
    -  `Handling ManyToMany
       Relation <query.html#handling-manytomany-relation>`__

4.  `Use Raw SQL <rawsql.html>`__
5.  `Transactions <transaction.html>`__
6.  `Model Definition <models.html>`__

    -  `Custom Table Names <models.html#custom-table-name>`__
    -  `Custom engine <models.html#custom-engine>`__
    -  `Set Parameters <models.html#set-parameters>`__
    -  `Relationship <models.html#relationship>`__
    -  `Model Fields Mapping with Database
       Type <models.html#model-fields-mapping-with-database-type>`__

7.  `Command Line <cmd.html>`__

    -  `Table Auto generating <cmd.html#table-auto-generating>`__
    -  `Print SQL Statements <cmd.html#print-sql-statements>`__

8.  `Test ORM <test.html>`__
9.  `Custom Fields <custom_fields.html>`__
10. `FAQ <faq.html>`__

.. |Build Status| image:: https://drone.io/github.com/W3-Engineers-Ltd/Radiant/status.png
   :target: https://drone.io/github.com/W3-Engineers-Ltd/Radiant/latest
.. |Go Walker| image:: http://gowalker.org/api/v1/badge
   :target: http://gowalker.org/github.com/W3-Engineers-Ltd/Radiant/client/orm

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