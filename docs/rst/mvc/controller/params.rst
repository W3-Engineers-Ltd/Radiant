Accept parameters
=================

Radiant will automatically parse data passed by user from GET, POST and
other methods. This data can be accessed using:

-  GetString(key string) string
-  GetStrings(key string) []string
-  GetInt(key string) (int, error)
-  GetInt8(key string) (int8, error)
-  GetInt16(key string) (int16, error)
-  GetInt32(key string) (int32, error)
-  GetInt64(key string) (int64, error)
-  GetUint8(key string) (uint8, error)
-  GetUint16(key string) (uint16, error)
-  GetUint32(key string) (uint32, error)
-  GetUint64(key string) (uint64, error)
-  GetBool(key string) (bool, error)
-  GetFloat(key string) (float64, error)

For example:

.. code:: go

   func (this *MainController) Post() {
       jsoninfo := this.GetString("jsoninfo")
       if jsoninfo == "" {
           this.Ctx.WriteString("jsoninfo is empty")
           return
       }
   }

More information about the request can be retrieved by accessing
``this.Ctx.Request``. For more details see
`Request <http://gowalker.org/net/http#Request>`__.

Parse to struct
---------------

Data submitted from a form may be assigned to a struct by mapping struct
fields to the form’s input elements and parsing all data into a struct.

Define struct:

.. code:: go

   type User struct {
       Id    int         `form:"-"`
       Name  interface{} `form:"username"`
       Age   int         `form:"age"`
       Email string
   }

Define form:

::

   <form id="user">
       name：<input name="username" type="text" />
       age：<input name="age" type="text" />
       email：<input name="Email" type="text" />
       <input type="submit" value="submit" />
   </form>

Parsing in Controller:

.. code:: go

   func (this *MainController) Post() {
       u := User{}
       if err := this.ParseForm(&u); err != nil {
           //handle error
       }
   }

Notes:

-  The same tag is used for the definition of structTag form and
   `renderform method <../view/view.html#renderform>`__.
-  If there is a form tag after the key while defining the struct, the
   value in the form which has the same name as that tag will be
   assigned. Otherwise, the value in the form which has the same name as
   that field name will be assigned. In the above example, Form values
   username and age will be assigned to Name and Age in user struct and
   Email will be assigned to Email in struct.
-  While calling the method ParseForm of the Controller the parameter
   passed in must be a pointer to a struct. Otherwise, the assignment
   will fail and will return a ``xx must be a struct pointer`` error.
-  Fields can be ignored by using lowercase for that field or by using
   ``-`` as the value of the tag.

Automatic Parameter Routing
---------------------------

Automatic parameter routing removes the need for boilerplate code like
``this.GetString(..)``, ``this.GetInt(..)`` etc. Instead https
parameters are injected directly as method parameters and the method
return values are rendered as http responses. This works in conjunction
with annotations to create a seamless integration.

How does it work?
~~~~~~~~~~~~~~~~~

Start by defining a regular controller method with a ``@router``
annotation and add parameters to the method signature

.. code:: go

   // @router /tasks
   func (c *TaskController) MyMethod(id int) {
   ...
   }

When an http request comes in that matches the defined routing Radiant
will scan the parameters in the method signature and try to find
matching http request paramters, where method parameter name is the http
request parameter name. Radiant will then convert them to the correct
parameter type and pass them to your method. By default Radiant will
look for parameters in the quey string (when using ``GET``) or form data
(when using ``POST``). If your routing definition contains parameters
Radiant will automatically search for them in the path:

.. code:: go

   // @router /task/:id
   func (c *TaskController) MyMethod(id int) {
   ...
   }

Annotations can also be used to indicate a parameter is passed in a
header or in the request body. Bego will search for it accordingly.

If a parameter is not found in the http request it will be passed to
your controller method as a zero value (i.e. 0 for int, false for bool
etc.). If a default value for that parameter has been defined in
annotations, Radiant will pass that default value if it is missing. To
differentiate between missing parameters and default values define the
parameter as a pointer, e.g.:

.. code:: go

   // @router /tasks
   func (c *TaskController) MyMethod(id *int) {
   ...
   }

If the parameter in the above case was missing, ``id`` would be null. If
the parameter exists and equals to zero, ``id`` would be 0. When using
annotations to create swagger documentation a parameter can be marked as
``required``. If the parameter is missing in the request a
``400 Bad Request`` error will be returned to the client:

.. code:: go

   // @Param   id     query   int true       "task id"
   // @router /tasks
   func (c *TaskController) MyMethod(id *int) {
   ...
   }

If Radiant can not convert the parameter to the requested type (i.e. if
a string is passed that can not be parsed as an integer) an error will
be returned to the client.

The following table shows which types are supported and how they are
parsed:

+-----------------+-----------------+---------------+-----------------+
| Data Type       | Location        | Example       | Comment         |
+=================+=================+===============+=================+
| int, int64,     | anywhere        | “1”,“-100”    | Uses            |
| uint etc.       |                 |               | ``strconv.Atoi( |
|                 |                 |               | value)``        |
+-----------------+-----------------+---------------+-----------------+
| float32,float64 | anywhere        | “1.5”, “-3.5” | Uses            |
|                 |                 |               | ``strconv.Parse |
|                 |                 |               | Float()``       |
+-----------------+-----------------+---------------+-----------------+
| bool            | anywhere        | “1”, “T”,     | Uses            |
|                 |                 | “false”       | ``strconv.Parse |
|                 |                 |               | Bool()``        |
+-----------------+-----------------+---------------+-----------------+
| time.Time       | anywhere        | “2017-01-01”  | Uses RFC3339 or |
|                 |                 | “2017-01-01T0 | short date      |
|                 |                 | 0:00:00Z”     | format          |
|                 |                 |               | (``"2006-01-02" |
|                 |                 |               | ``)             |
|                 |                 |               | when parsing    |
+-----------------+-----------------+---------------+-----------------+
| []string, []int | query           | “A,B,C”       | Any type is     |
| etc.            |                 | “1,2,3”       | supported as a  |
|                 |                 |               | slice. When it  |
|                 |                 |               | is located in   |
|                 |                 |               | the query       |
|                 |                 |               | string, it is   |
|                 |                 |               | parsed as a     |
|                 |                 |               | comma separated |
|                 |                 |               | list            |
+-----------------+-----------------+---------------+-----------------+
| []string, []int | body            | [“A”,“B”,“C”] | When slices are |
| etc.            |                 | [1,2,3]       | located in the  |
|                 |                 |               | request body    |
|                 |                 |               | they are parsed |
|                 |                 |               | as JSON arrays  |
+-----------------+-----------------+---------------+-----------------+
| []byte          | anywhere        | “ABC”         | byte[] is not   |
|                 |                 |               | treated as an   |
|                 |                 |               | array but as a  |
|                 |                 |               | string          |
+-----------------+-----------------+---------------+-----------------+
| \*int,          | anywhere        | Pointers will |                 |
| \*string,       |                 | receive null  |                 |
| \*float etc.    |                 | if the        |                 |
|                 |                 | parameter is  |                 |
|                 |                 | missing from  |                 |
|                 |                 | the request   |                 |
|                 |                 | otherwise, it |                 |
|                 |                 | will behave   |                 |
|                 |                 | the same as   |                 |
|                 |                 | defined in    |                 |
|                 |                 | the other     |                 |
|                 |                 | rows          |                 |
+-----------------+-----------------+---------------+-----------------+
| structs, all    | anywhere        | {“X”:“Y”}     | structs and     |
| others          |                 |               | other types     |
|                 |                 |               | (e.g. maps) are |
|                 |                 |               | always parsed   |
|                 |                 |               | as JSON using   |
|                 |                 |               | ``json.Unmarsha |
|                 |                 |               | l()``           |
+-----------------+-----------------+---------------+-----------------+

How are method return values handled?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Method return values are handled automatically in the same manner as
parameters. A method can have one or more return values and Radiant will
render all of them to the response. The best practice is to define one
result as a ‘regular’ type (i.e. a map, a struct or any other data type)
and another as an error data type:

.. code:: go

   // @Param   id     query   int true       "task id"
   // @router /tasks
   func (c *TaskController) MyMethod(id *int) (*MyModel, error) {
   ...
   }

In the code above the method can return three different results: - Only
``MyModel`` (nil ``error``) - Only ``error`` (nil ``MyModel``) - Both
``MyModel`` and ``error``

When a regular type is returned it is rendered directly as JSON, and
when an error is returned it is rendered as an http status code. Radiant
will handle all cases correctly and supports returning both response
body and http error if both values are non-nil.

A few helper types will return common http status codes easily. For
example, ``404 Not Found``, ``302 Redirect`` or other http status codes
like in the following example:

.. code:: go

   func (c *TaskController) MyMethod(id *int) (*MyModel, error) {
     if /* not found */ {
       return nil, context.NotFound
     } else if /* some error */ {
       return nil, context.StatusCode(401)
     } else {
       return &MyModel{}, nil
     }
   }

How annotations work in conjuction with method parameters?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Automatic Parameter Routing works best together with ``@Param``
annotations. The following features are supported with annotations: - If
a parameter is marked as required, Radiant will return an error if the
parameter is not present in the http request:

.. code:: go

   // @Param   brand_id    query   int true       "brand id"

(the ``true`` option in the annotation above indicates that brand_id is
a required parameter) - If a parameter has a default value and it does
not exist in the http request, Radiant will pass that default value to
the method:

.. code:: go

   // @Param   brand_id    query   int false  5  "brand id"

(the ``5`` in the annotation above indicates that this is the default
value for that parameter) - The location parameter in the annotation
indicates where radiant will search for that parameter in the request
(i.e. query, header, body etc.)

.. code:: go

   // @Param   brand_id    path    int     true  "brand id"
   // @Param   category    query   string  false "category" 
   // @Param   token   header  string  false "auth token"
   // @Param   task    body    {models.Task} false "the task object"

-  If a parameter name in the http request is different from the method
   parameter name, you can “redirect” the parameter using the ``=>``
   notation. This is useful, for example, When a header name is
   ``X-Token`` and the method parameter is named ``x_token``:

.. code:: go

   // @Param   X-Token=>x_token    header  string  false "auth token"

-  A parameter swagger data type can be inferred from the method to make
   maintainance easier. Use the ``auto`` data type and Radiant will
   generate the correct swagger documentation:

.. code:: go

   // @Param   id     query   auto true       "task id"
   // @router /tasks
   func (c *TaskController) MyMethod(id int) (*MyModel, error) {
   ...
   }

Retrieving data from request body
---------------------------------

In API application development always use ``JSON`` or ``XML`` as the
data type. To retrieve the data from the request body:

1. Set ``copyrequestbody = true`` in configuration file.
2. Then in the Controller you can

.. code:: go

   func (this *ObjectController) Post() {
       var ob models.Object
       json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
       objectid := models.AddOne(ob)
       this.Data["json"] = map[string]interface{}{"ObjectId": objectid }
       this.ServeJSON()
   }

Uploading files
---------------

To upload files with Radiant set attribute
``enctype="multipart/form-data"`` in your form.

Usually an uploaded file is stored in the system memory, but if the file
size is larger than the memory size limitation in the configuration
file, the file will be stored in a temporary file. The default memory
size is 64M but can be changed using (bit shift):

::

   radiant.MaxMemory = 1<<22

Or it can be set manualy in the configuration file (bit shift):

::

   maxmemory = 1<<22

In v2.x, there is another parameter ``MaxUploadSize`` used to limit the
max size of uploading files.

If you upload multiple files in one request, it limits the sum size of
those files.

Usually, ``web.BConfig.MaxMemory`` should be less than
``web.BConfig.MaxUploadSize``:

1. if file size < ``MaxMemory``, handling file in memory;
2. ``MaxMemory`` < file size < ``MaxUploadSize``, handling file by using
   temporary directory.
3. file size > ``MaxUploadSize``, return 413;

Radiant provides three functions to handle file uploads:

-  GetFile(key string) (multipart.File, \*multipart.FileHeader, error)

This method is used to read the file name ``the_file`` from form and
return the information. The uploaded file can then be processed based on
this information, such as filter or save the file.

-  GetFiles(key string) ([]*multipart.FileHeader, error)

This method returns all the multi-upload files:

.. code:: go

   func (m *MainController) Post() {
       // 'files' is the name of the multipart form input
       files, err := m.GetFiles("files")
       if err != nil {
           logger.Error(err.Error())
       }
       ... do something with files

-  SaveToFile(fromfile, tofile string) error

This method implements the saving function based on the method
``GetFile``

Here is an example of saving a file:

.. code:: go

   func (this *MainController) Post() {
       this.SaveToFile("the_file","/var/www/uploads/uploaded_file.txt")
   }

Data Bind
---------

Data bind lets the user bind the request data to a variable, the request
url as follows:

::

   ?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie

.. code:: go

   var id int
   ctx.Input.Bind(&id, "id")  // id ==123

   var isok bool
   ctx.Input.Bind(&isok, "isok")  // isok ==true

   var ft float64
   ctx.Input.Bind(&ft, "ft")  // ft ==1.2

   ol := make([]int, 0, 2)
   ctx.Input.Bind(&ol, "ol")  // ol ==[1 2]

   ul := make([]string, 0, 2)
   ctx.Input.Bind(&ul, "ul")  // ul ==[str array]

   user struct{Name}
   ctx.Input.Bind(&user, "user")  // user =={Name:"astaxie"}

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