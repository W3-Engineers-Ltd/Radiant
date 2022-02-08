Model Definition
================

Model names are used for database data conversion and `Database Schema
Generation <cmd.html#database-schema-generation>`__

Naming conventions
------------------

Table name conversion consists in translating camel case used for model
names to snake case for table names as follows:

::

   AuthUser -> auth_user
   Auth_User -> auth__user
   DB_AuthUser -> d_b__auth_user

In other words, all is converted to lower case and ``_`` is the
separator. Every uppercase add a separator before it, except the first
one.

Custom table name
-----------------

Using ``TableNameI`` interface:

.. code:: go

   type User struct {
       Id int
       Name string
   }

   func (u *User) TableName() string {
       return "auth_user"
   }

If you set `prefix <orm.html#registermodelwithprefix>`__ to ``prefix_``,
the table name will be ``prefix_auth_user``.

Custom index
------------

Using ``TableIndexI`` interface:

Add index to one or more fields:

.. code:: go

   type User struct {
       Id    int
       Name  string
       Email string
   }

   // multiple fields index
   func (u *User) TableIndex() [][]string {
       return [][]string{
           []string{"Id", "Name"},
       }
   }

   // multiple fields unique key
   func (u *User) TableUnique() [][]string {
       return [][]string{
           []string{"Name", "Email"},
       }
   }

Custom engine
-------------

Only supports MySQL database

The default engine is the default engine of the current database engine
of your mysql settings.

Using ``TableEngineI`` interface:

.. code:: go

   type User struct {
       Id    int
       Name  string
       Email string
   }

   // Set engine to INNODB
   func (u *User) TableEngine() string {
       return "INNODB"
   }

Set parameters
--------------

.. code:: go

   orm:"null;rel(fk)"

Use ``;`` as the separator of multiple settings. Use ``,`` as the
separator if a setting has multiple values.

Ignore field
^^^^^^^^^^^^

Use ``-`` to ignore field in the struct.

.. code:: go

   type User struct {
   ...
       AnyField string `orm:"-"`
   ...
   }

auto
^^^^

When Field type is int, int32, int64, uint, uint32 or uint64, you can
set it as auto increment.

-  If there is no primary key in the model definition, the field ``Id``
   with one of the types above will be considered as auto increment key

pk
^^

Set as primary key. Used for using other type field as primary key.

null
^^^^

Fields are ``NOT NULL`` by default. Set null to ``ALLOW NULL``.

.. code:: go

   Name string `orm:"null"`

index
^^^^^

Add index for one field

unique
^^^^^^

Add unique key for one field

.. code:: go

   Name string `orm:"unique"`

column
^^^^^^

Set column name in db table for field.

.. code:: go

   Name string `orm:"column(user_name)"`

size
^^^^

Default value for string field is varchar(255).

It will use varchar(size) after setting.

.. code:: go

   Title string `orm:"size(60)"`

digits / decimals
^^^^^^^^^^^^^^^^^

Set precision for float32 or float64.

.. code:: go

   Money float64 `orm:"digits(12);decimals(4)"`

Total 12 digits, 4 digits after point. For example: ``12345678.1234``

auto_now / auto_now_add
^^^^^^^^^^^^^^^^^^^^^^^

.. code:: go

   Created time.Time `orm:"auto_now_add;type(datetime)"`
   Updated time.Time `orm:"auto_now;type(datetime)"`

-  auto_now: every save will update time.
-  auto_now_add: set time at the first save

This setting won’t affect massive ``update``.

type
^^^^

If set type as date, the field’s db type is date.

.. code:: go

   Created time.Time `orm:"auto_now_add;type(date)"`

If set type as datetime, the field’s db type is datetime.

.. code:: go

   Created time.Time `orm:"auto_now_add;type(datetime)"`

Time Precision
^^^^^^^^^^^^^^

.. code:: go

   type User struct {
   ...
   Created time.Time `orm:"type(datetime);precision(4)"`
   ...
   }

default value
^^^^^^^^^^^^^

you could use it like:

.. code:: go


   import (
   "github.com/W3-Engineers-Ltd/Radiant/client/orm/filter/bean"
   "github.com/W3-Engineers-Ltd/Radiant/client/orm"
   )

   type DefaultValueTestEntity struct {
   Id            int
   Age           int `default:"12"`
   AgeInOldStyle int `orm:"default(13);radical()"`
   AgeIgnore     int
   }

   func XXX() {
       builder := bean.NewDefaultValueFilterChainBuilder(nil, true, true)
       orm.AddGlobalFilterChain(builder.FilterChain)
       o := orm.NewOrm()
       _, _ = o.Insert(&User{
           ID: 1,
           Name: "Tom",
       })
   }

``NewDefaultValueFilterChainBuilder``\ will create an instance of
``DefaultValueFilterChainBuilder`` In radiant v1.x, the default value
config looks like ``orm:default(xxxx)`` But the default value in 2.x is
``default:xxx``, so if you want to be compatible with v1.x, please pass
true as ``compatibleWithOldStyle``

Comment
^^^^^^^

Set comment value for field.

.. code:: go

   type User struct {
       ...
       Status int `orm:"default(1);description(this is status)"`
       ...
   }

Relationships
-------------

One to one
^^^^^^^^^^

**RelOneToOne**:

.. code:: go

   type User struct {
       ...
       Profile *Profile `orm:"null;rel(one);on_delete(set_null)"`
       ...
   }

The reverse relationship **RelReverseOne**:

.. code:: go

   type Profile struct {
       ...
       User *User `orm:"reverse(one)"`
       ...
   }

One to many
^^^^^^^^^^^

**RelForeignKey**:

.. code:: go

   type Post struct {
       ...
       User *User `orm:"rel(fk)"` // RelForeignKey relation
       ...
   }

The reverse relationship **RelReverseMany**:

.. code:: go

   type User struct {
       ...
       Posts []*Post `orm:"reverse(many)"` // reverse relationship of fk
       ...
   }

Many to many
^^^^^^^^^^^^

**RelManyToMany**:

.. code:: go

   type Post struct {
       ...
       Tags []*Tag `orm:"rel(m2m)"` // ManyToMany relation
       ...
   }

The reverse relationship **RelReverseMany**:

.. code:: go

   type Tag struct {
       ...
       Posts []*Post `orm:"reverse(many)"`
       ...
   }

In this example, by default the auto-generated table name is:
``post_tag``. The name of the struct in which we have ``orm:"rel(m2m)"``
defines the first half part, the name of the struct in which we have
``orm:"reverse(many)"`` defines the other half. It respects the naming
conversion convention we have seen in `Naming
conventions <#naming-conventions>`__

rel_table / rel_through
'''''''''''''''''''''''

This setting is for ``orm:"rel(m2m)"`` field:

::

   rel_table       Set the auto-generated m2m connecting table name
   rel_through     If you want to use custom m2m connecting table, set name by using this setting.
                 Format: `project_path/current_package.ModelName`
                 For example: `app/models.PostTagRel` PostTagRel table needs to have a relationship to Post table and Tag table.

If rel_table is set, rel_through is ignored.

You can set these as follows:

``orm:"rel(m2m);rel_table(the_table_name)"``

``orm:"rel(m2m);rel_through(project_path/current_package.ModelName)"``

on_delete
^^^^^^^^^

Set how to deal with field if related relationship is deleted:

::

   cascade        cascade delete (default)
   set_null       set to NULL. Need to set null = true
   set_default    set to default value. Need to set default value.
   do_nothing     do nothing. ignore.

.. code:: go

   type User struct {
       ...
       Profile *Profile `orm:"null;rel(one);on_delete(set_null)"`
       ...
   }
   type Profile struct {
       ...
       User *User `orm:"reverse(one)"`
       ...
   }

   // Set User.Profile to NULL while deleting Profile

Examples of on_delete
^^^^^^^^^^^^^^^^^^^^^

.. code:: go

   type User struct {
       Id int
       Name string
   }

   type Post struct {
       Id int
       Title string
       User *User `orm:"rel(fk)"`
   }

Assume Post -> User is ManyToOne relationship by foreign key.

::

   o.Filter("Id", 1).Delete()

This will delete User with Id 1 and all his Posts.

If you don’t want to delete the Posts, you need to set ``set_null``

.. code:: go

   type Post struct {
       Id int
       Title string
       User *User `orm:"rel(fk);null;on_delete(set_null)"`
   }

In this case, only set related Post.user_id to NULL while deleting.

Usually for performance purposes, it doesn’t matter to have redundant
data. The massive deletion is the real problem

.. code:: go

   type Post struct {
       Id int
       Title string
       User *User `orm:"rel(fk);null;on_delete(do_nothing)"`
   }

So just don’t change Post (ignore it) while deleting User.

Model fields mapping with database type
---------------------------------------

Here is the recommended database type mapping. It’s also the standard
for table generation.

All the fields are **NOT NULL** by default.

MySQL
^^^^^

============================================
===============================
go                                           mysql
============================================
===============================
int, int32 - set as auto or name is ``Id``   integer AUTO_INCREMENT
int64 - set as auto or name is\ ``Id``       bigint AUTO_INCREMENT
uint, uint32 - set as auto or name is ``Id`` integer unsigned AUTO_INCREMENT
uint64 - set as auto or name is ``Id``       bigint unsigned AUTO_INCREMENT
bool                                         bool
string - default size 255                    varchar(size)
string - set type(char)                      char(size)
string - set type(text)                      longtext
time.Time - set type as date                 date
time.Time                                    datetime
byte                                         tinyint unsigned
rune                                         integer
int                                          integer
int8                                         tinyint
int16                                        smallint
int32                                        integer
int64                                        bigint
uint                                         integer unsigned
uint8                                        tinyint unsigned
uint16                                       smallint unsigned
uint32                                       integer unsigned
uint64                                       bigint unsigned
float32                                      double precision
float64                                      double precision
float64 - set digits and decimals            numeric(digits, decimals)
============================================
===============================

Sqlite3
^^^^^^^

+-----------------------------------+-----------------------------------+
| go                                | sqlite3                           |
+===================================+===================================+
| int, int32, int64, uint, uint32,  | integer AUTOINCREMENT             |
| uint64 - set as auto or name is   |                                   |
| ``Id``                            |                                   |
+-----------------------------------+-----------------------------------+
| bool                              | bool                              |
+-----------------------------------+-----------------------------------+
| string - default size 255         | varchar(size)                     |
+-----------------------------------+-----------------------------------+
| string - set type(char)           | character(size)                   |
+-----------------------------------+-----------------------------------+
| string - set type(text)           | text                              |
+-----------------------------------+-----------------------------------+
| time.Time - set type as date      | date                              |
+-----------------------------------+-----------------------------------+
| time.Time                         | datetime                          |
+-----------------------------------+-----------------------------------+
| byte                              | tinyint unsigned                  |
+-----------------------------------+-----------------------------------+
| rune                              | integer                           |
+-----------------------------------+-----------------------------------+
| int                               | integer                           |
+-----------------------------------+-----------------------------------+
| int8                              | tinyint                           |
+-----------------------------------+-----------------------------------+
| int16                             | smallint                          |
+-----------------------------------+-----------------------------------+
| int32                             | integer                           |
+-----------------------------------+-----------------------------------+
| int64                             | bigint                            |
+-----------------------------------+-----------------------------------+
| uint                              | integer unsigned                  |
+-----------------------------------+-----------------------------------+
| uint8                             | tinyint unsigned                  |
+-----------------------------------+-----------------------------------+
| uint16                            | smallint unsigned                 |
+-----------------------------------+-----------------------------------+
| uint32                            | integer unsigned                  |
+-----------------------------------+-----------------------------------+
| uint64                            | bigint unsigned                   |
+-----------------------------------+-----------------------------------+
| float32                           | real                              |
+-----------------------------------+-----------------------------------+
| float64                           | real                              |
+-----------------------------------+-----------------------------------+
| float64 - set digits and decimals | decimal                           |
+-----------------------------------+-----------------------------------+

PostgreSQL
^^^^^^^^^^

+-----------------------------------+-----------------------------------+
| go                                | postgres                          |
+===================================+===================================+
| int, int32, int64, uint, uint32,  | serial                            |
| uint64 - set as auto or name is   |                                   |
| ``Id``                            |                                   |
+-----------------------------------+-----------------------------------+
| bool                              | bool                              |
+-----------------------------------+-----------------------------------+
| string - if not set size default  | varchar(size)                     |
| text                              |                                   |
+-----------------------------------+-----------------------------------+
| string - set type(char)           | char(size)                        |
+-----------------------------------+-----------------------------------+
| string - set type(text)           | text                              |
+-----------------------------------+-----------------------------------+
| string - set type(json)           | json                              |
+-----------------------------------+-----------------------------------+
| string - set type(jsonb)          | jsonb                             |
+-----------------------------------+-----------------------------------+
| time.Time - set type as date      | date                              |
+-----------------------------------+-----------------------------------+
| time.Time                         | timestamp with time zone          |
+-----------------------------------+-----------------------------------+
| byte                              | smallint CHECK(“column” >= 0 AND  |
|                                   | “column” <= 255)                  |
+-----------------------------------+-----------------------------------+
| rune                              | integer                           |
+-----------------------------------+-----------------------------------+
| int                               | integer                           |
+-----------------------------------+-----------------------------------+
| int8                              | smallint CHECK(“column” >= -127   |
|                                   | AND “column” <= 128)              |
+-----------------------------------+-----------------------------------+
| int16                             | smallint                          |
+-----------------------------------+-----------------------------------+
| int32                             | integer                           |
+-----------------------------------+-----------------------------------+
| int64                             | bigint                            |
+-----------------------------------+-----------------------------------+
| uint                              | bigint CHECK(“column” >= 0)       |
+-----------------------------------+-----------------------------------+
| uint8                             | smallint CHECK(“column” >= 0 AND  |
|                                   | “column” <= 255)                  |
+-----------------------------------+-----------------------------------+
| uint16                            | integer CHECK(“column” >= 0)      |
+-----------------------------------+-----------------------------------+
| uint32                            | bigint CHECK(“column” >= 0)       |
+-----------------------------------+-----------------------------------+
| uint64                            | bigint CHECK(“column” >= 0)       |
+-----------------------------------+-----------------------------------+
| float32                           | double precision                  |
+-----------------------------------+-----------------------------------+
| float64                           | double precision                  |
+-----------------------------------+-----------------------------------+
| float64 - set digits and decimals | numeric(digits, decimals)         |
+-----------------------------------+-----------------------------------+

Relational fields
-----------------

It’s field type depends on related primary key.

-  RelForeignKey
-  RelOneToOne
-  RelManyToMany
-  RelReverseOne
-  RelReverseMany


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