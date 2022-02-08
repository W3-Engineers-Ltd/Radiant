CRUD of Object
==============

If the value of the primary key is already known, ``Read``, ``Insert``,
``Update``, ``Delete`` can be used to manipulate the object.

.. code:: go

   o := orm.NewOrm()
   user := new(User)
   user.Name = "slene"

   fmt.Println(o.Insert(user))

   user.Name = "Your"
   fmt.Println(o.Update(user))
   fmt.Println(o.Read(user))
   fmt.Println(o.Delete(user))

To query the object by conditions see `Query in
advance <query.html#all>`__

Read
----

.. code:: go

   o := orm.NewOrm()
   user := User{Id: 1}

   err := o.Read(&user)

   if err == orm.ErrNoRows {
       fmt.Println("No result found.")
   } else if err == orm.ErrMissPK {
       fmt.Println("No primary key found.")
   } else {
       fmt.Println(user.Id, user.Name)
   }

Read uses primary key by default. But it can use other fields as well:

.. code:: go

   user := User{Name: "slene"}
   err := o.Read(&user, "Name")
   ...

Other fields of the object are set to the default value according to the
field type.

For detailed single object query, see `One <query.html#one>`__

ReadOrCreate
------------

Try to read a row from the database, or insert one if it doesn’t exist.

At least one condition field must be supplied, multiple condition fields
are also supported.

.. code:: go

   o := orm.NewOrm()
   user := User{Name: "slene"}
   // Three return values：Is Created，Object Id，Error
   if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
       if created {
           fmt.Println("New Insert an object. Id:", id)
       } else {
           fmt.Println("Get an object. Id:", id)
       }
   }

Insert
------

The first return value is auto inc Id value.

.. code:: go

   o := orm.NewOrm()
   var user User
   user.Name = "slene"
   user.IsActive = true

   id, err := o.Insert(&user)
   if err == nil {
       fmt.Println(id)
   }

After creation, it will assign values for auto fields.

InsertMulti
-----------

Insert multiple objects in one api.

Like sql statement:

::

   insert into table (name, age) values("slene", 28),("astaxie", 30),("unknown", 20)

The 1st param is the number of records to insert in one bulk statement.
The 2nd param is models slice.

The return value is the number of successfully inserted rows.

.. code:: go

   users := []User{
       {Name: "slene"},
       {Name: "astaxie"},
       {Name: "unknown"},
       ...
   }
   successNums, err := o.InsertMulti(100, users)

When bulk is equal to 1, then models will be inserted one by one.

Update
------

The first return value is the number of affected rows.

.. code:: go

   o := orm.NewOrm()
   user := User{Id: 1}
   if o.Read(&user) == nil {
       user.Name = "MyName"
       if num, err := o.Update(&user); err == nil {
           fmt.Println(num)
       }
   }

Update updates all fields by default. You can update specified fields:

.. code:: go

   // Only update Name
   o.Update(&user, "Name")
   // Update multiple fields
   // o.Update(&user, "Field1", "Field2", ...)
   ...

For detailed object update, see `One <query.html#one>`__

Delete
------

The first return value is the number of affected rows.

.. code:: go

   o := orm.NewOrm()
   if num, err := o.Delete(&User{Id: 1}); err == nil {
       fmt.Println(num)
   }

Delete will also manipulate reverse relationships. E.g.: ``Post`` has a
foreign key to ``User``. If on_delete is set to ``cascade``, ``Post``
will be deleted while delete ``User``.

After deleting, it will clean up values for auto fields.

**Changed in 1.0.3** After deleting, it will **not** clean up values for
auto fields.

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