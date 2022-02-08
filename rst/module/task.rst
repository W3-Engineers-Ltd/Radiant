Core task Module
================

Installation
------------

::

   go get github.com/W3-Engineers-Ltd/Radiant/task

Tasks
-----

Tasks work very similarly to cron jobs. Tasks are used to run a job
outside the normal request/response cycle. These can be adhoc or
scheduled to run regularly. Examples include: Reporting memory and
goroutine status, periodically triggering GC or cleaning up log files at
fixed intervals.

Creating a new Task
~~~~~~~~~~~~~~~~~~~

To initialize a task implement :

::

   tk1 := task.NewTask("tk1", "0 12 * * * *", func(ctx context.Context) error {
       fmt.Println("tk1")
       return nil
   })

The NewTask signature:

::

   NewTask(tname string, spec string, f TaskFunc) *Task    

-  ``tname``: Task name
-  ``spec``: Task format. See below for details.
-  ``f``: The function which will be run as the task.

To implement this task, add it to the global task list and start it.

::

   task.AddTask("tk1", tk1)
   task.StartTask()
   defer task.StopTask()

Testing the TaskFunc
~~~~~~~~~~~~~~~~~~~~

Use the code below to test if the TaskFunc is working correctly.

::

   err := tk.Run()
   if err != nil {
       t.Fatal(err)
   }

spec in detail
~~~~~~~~~~~~~~

``spec`` specifies when the new Task will be run. Its format is the same
as that of traditional crontab:

::

   // The first 6 parts are:
   //       second: 0-59
   //       minute: 0-59
   //       hour: 1-23
   //       day: 1-31
   //       month: 1-12
   //       weekdays: 0-6（0 is Sunday）

   // Some special sign:
   //       *: any time
   //       ,: separator. E.g.: 2,4 in the third part means run at 2 and 4 o'clock
   //　　    －: range. E.g.: 1-5 in the third part means run between 1 and 5 o'clock
   //       /n : run once every n time. E.g.: */1 in the third part means run once every an hour. Same as 1-23/1
   /////////////////////////////////////////////////////////
   //  0/30 * * * * *                        run every 30 seconds
   //  0 43 21 * * *                         run at 21:43
   //  0 15 05 * * *                         run at 05:15
   //  0 0 17 * * *                          run at 17:00
   //  0 0 17 * * 1                          run at 17:00 of every Monday
   //  0 0,10 17 * * 0,2,3                   run at 17:00 and 17:10 of every Sunday, Tuesday and Wednesday
   //  0 0-10 17 1 * *                       run once every minute from 17:00 to 7:10 on 1st day of every month
   //  0 0 0 1,15 * 1                        run at 0:00 on 1st and 15th of each month and every Monday
   //  0 42 4 1 * *                          run at 4:42 on 1st of every month
   //  0 0 21 * * 1-6                        run at 21:00 from Monday to Saturday
   //  0 0,10,20,30,40,50 * * * *            run every 10 minutes
   //  0 */10 * * * *                        run every 10 minutes
   //  0 * 1 * * *                           run every one minute from 1:00 to 1:59
   //  0 0 1 * * *                           run at 1:00
   //  0 0 */1 * * *                         run at :00 of every hour
   //  0 0 * * * *                           run at :00 of every hour
   //  0 2 8-20/3 * * *                      run at 8:02, 11:02, 14:02, 17:02 and 20:02
   //  0 30 5 1,15 * *                       run at 5:30 of 1st and 15th of every month

Debug module (Already moved to utils module)
--------------------------------------------

We always use print for debugging. But the default output is not good
enough for debugging. Radiant provides this debug module

-  Display() print result to console
-  GetDisplayString() return the string

It print key/value pairs. The following code:

::

   Display("v1", 1, "v2", 2, "v3", 3)

will output:

::

   2013/12/16 23:48:41 [Debug] at TestPrint() [/Users/astaxie/github/radiant/task/debug_test.go:13]

   [Variables]
   v1 = 1
   v2 = 2
   v3 = 3  

For pointer type:

::

   type mytype struct {
       next *mytype
       prev *mytype
   }   

   var v1 = new(mytype)
   var v2 = new(mytype)

   v1.prev = nil
   v1.next = v2

   v2.prev = v1
   v2.next = nil

   Display("v1", v1, "v2", v2)

The output result

::

   2013/12/16 23:48:41 [Debug] at TestPrintPoint() [/Users/astaxie/github/radiant/task/debug_test.go:26]

   [Variables]
   v1 = &task.mytype{
       next: &task.mytype{
           next: nil,
           prev: 0x210335420,
       },
       prev: nil,
   }
   v2 = &task.mytype{
       next: nil,
       prev: &task.mytype{
           next: 0x210335430,
           prev: nil,
       },
   }       

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