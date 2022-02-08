Raw SQL to query
================

-  Using Raw SQL to query doesn’t require an ORM definition
-  Multiple databases support ``?`` as placeholders and auto convert.
-  The params of query support Model Struct, Slice and Array

.. code:: go

   ids := []int{1, 2, 3}
   p.Raw("SELECT name FROM user WHERE id IN (?, ?, ?)", ids)

Create a **RawSeter**

.. code:: go

   o := NewOrm()
   var r RawSeter
   r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")

.. code:: go


   // RawSeter raw query seter
   // create From Ormer.Raw
   // for example:
   //  sql := fmt.Sprintf("SELECT %sid%s,%sname%s FROM %suser%s WHERE id = ?",Q,Q,Q,Q,Q,Q)
   //  rs := Ormer.Raw(sql, 1)
   type RawSeter interface {
       // execute sql and get result
       Exec() (sql.Result, error)
       // query data and map to container
       // for example:
       //  var name string
       //  var id int
       //  rs.QueryRow(&id,&name) // id==2 name=="slene"
       QueryRow(containers ...interface{}) error

       // query data rows and map to container
       //  var ids []int
       //  var names []int
       //  query = fmt.Sprintf("SELECT 'id','name' FROM %suser%s", Q, Q)
       //  num, err = dORM.Raw(query).QueryRows(&ids,&names) // ids=>{1,2},names=>{"nobody","slene"}
       QueryRows(containers ...interface{}) (int64, error)
       SetArgs(...interface{}) RawSeter
       // query data to []map[string]interface
       // see QuerySeter's Values
       Values(container *[]Params, cols ...string) (int64, error)
       // query data to [][]interface
       // see QuerySeter's ValuesList
       ValuesList(container *[]ParamsList, cols ...string) (int64, error)
       // query data to []interface
       // see QuerySeter's ValuesFlat
       ValuesFlat(container *ParamsList, cols ...string) (int64, error)
       // query all rows into map[string]interface with specify key and value column name.
       // keyCol = "name", valueCol = "value"
       // table data
       // name  | value
       // total | 100
       // found | 200
       // to map[string]interface{}{
       //  "total": 100,
       //  "found": 200,
       // }
       RowsToMap(result *Params, keyCol, valueCol string) (int64, error)
       // query all rows into struct with specify key and value column name.
       // keyCol = "name", valueCol = "value"
       // table data
       // name  | value
       // total | 100
       // found | 200
       // to struct {
       //  Total int
       //  Found int
       // }
       RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error)

       // return prepared raw statement for used in times.
       // for example:
       //  pre, err := dORM.Raw("INSERT INTO tag (name) VALUES (?)").Prepare()
       //  r, err := pre.Exec("name1") // INSERT INTO tag (name) VALUES (`name1`)
       Prepare() (RawPreparer, error)
   }

Exec
^^^^

Run sql query and return
`sql.Result <http://gowalker.org/database/sql#Result>`__ object

.. code:: go

   res, err := o.Raw("UPDATE user SET name = ?", "your").Exec()
   if err == nil {
       num, _ := res.RowsAffected()
       fmt.Println("mysql row affected nums: ", num)
   }

QueryRow
^^^^^^^^

QueryRow and QueryRows support high-level sql mapper.

Supports struct:

.. code:: go

   type User struct {
       Id   int
       Name string
   }

   var user User
   err := o.Raw("SELECT id, name FROM user WHERE id = ?", 1).QueryRow(&user)

..

   from Radiant 1.1.0 remove multiple struct support `ISSUE
   384 <https://github.com/W3-Engineers-Ltd/Radiant/issues/384>`__

QueryRows
^^^^^^^^^

QueryRows supports the same mapping rules as QueryRow but all of them
are slice.

.. code:: go

   type User struct {
       Id   int
       Name string
   }

   var users []User
   num, err := o.Raw("SELECT id, name FROM user WHERE id = ?", 1).QueryRows(&users)
   if err == nil {
       fmt.Println("user nums: ", num)
   }

..

   from Radiant 1.1.0 remove multiple struct support `ISSUE
   384 <https://github.com/W3-Engineers-Ltd/Radiant/issues/384>`__

SetArgs
^^^^^^^

Changing args param in Raw(sql, args…) can return a new RawSeter.

It can reuse the same SQL query but different params.

.. code:: go

   res, err := r.SetArgs("arg1", "arg2").Exec()
   res, err := r.SetArgs("arg1", "arg2").Exec()
   ...

Values / ValuesList / ValuesFlat
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The resultSet values returned by Raw SQL query are ``string``. NULL
field will return empty string \`\`

   from Radiant 1.1.0 Values, ValuesList, ValuesFlat. The returned
   fields can be specified. Generally you don’t need to specify. Because
   the field names are already defined in your SQL.

Values
^^^^^^

The key => value pairs of resultSet:

.. code:: go

   var maps []orm.Params
   num, err := o.Raw("SELECT user_name FROM user WHERE status = ?", 1).Values(&maps)
   if err == nil && num > 0 {
       fmt.Println(maps[0]["user_name"]) // slene
   }

ValuesList
^^^^^^^^^^

slice of resultSet

.. code:: go

   var lists []orm.ParamsList
   num, err := o.Raw("SELECT user_name FROM user WHERE status = ?", 1).ValuesList(&lists)
   if err == nil && num > 0 {
       fmt.Println(lists[0][0]) // slene
   }

ValuesFlat
^^^^^^^^^^

Return slice of a single field:

.. code:: go

   var list orm.ParamsList
   num, err := o.Raw("SELECT id FROM user WHERE id < ?", 10).ValuesFlat(&list)
   if err == nil && num > 0 {
       fmt.Println(list) // []{"1","2","3",...}
   }

RowsToMap
^^^^^^^^^

SQL query results

===== =====
name  value
===== =====
total 100
found 200
===== =====

map rows results to map

.. code:: go

   res := make(orm.Params)
   nums, err := o.Raw("SELECT name, value FROM options_table").RowsToMap(&res, "name", "value")
   // res is a map[string]interface{}{
   //  "total": 100,
   //  "found": 200,
   // }

RowsToStruct
^^^^^^^^^^^^

SQL query results

===== =====
name  value
===== =====
total 100
found 200
===== =====

map rows results to struct

.. code:: go

   type Options struct {
       Total int
       Found int
   }

   res := new(Options)
   nums, err := o.Raw("SELECT name, value FROM options_table").RowsToStruct(res, "name", "value")
   fmt.Println(res.Total) // 100
   fmt.Println(res.Found) // 200

..

   support name conversion: snake -> camel, eg: SELECT user_name … to
   your struct field UserName.

Prepare
^^^^^^^

Prepare once and exec multiple times to improve the speed of batch
execution.

.. code:: go

   p, err := o.Raw("UPDATE user SET name = ? WHERE name = ?").Prepare()
   res, err := p.Exec("testing", "slene")
   res, err  = p.Exec("testing", "astaxie")
   ...
   ...
   p.Close() // Don't forget to close the prepare.

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