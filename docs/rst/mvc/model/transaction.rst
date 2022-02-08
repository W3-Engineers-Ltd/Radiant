Transaction
===========

There are two ways to handle transaction in Radiant. One is closure:

.. code:: go

   // Radiant will manage the transaction's lifecycle
   // if the @param task return error, the transaction will be rollback
   // or the transaction will be committed
   err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
       // data
       user := new(User)
       user.Name = "test_transaction"

       // insert data
       // Using txOrm to execute SQL
       _, e := txOrm.Insert(user)
       // if e != nil the transaction will be rollback
       // or it will be committed
       return e
   })

In this way, the first parameter is ``task``, all DB operation should be
inside the task.

If the task return error, Radiant rollback the transaction.

We recommend you to use this way.

Another way is that users handle transaction manually:

.. code:: go

       o := orm.NewOrm()
       to, err := o.Begin()
       if err != nil {
           logs.Error("start the transaction failed")
           return
       }

       user := new(User)
       user.Name = "test_transaction"

       // do something with to. to is an instance of TxOrm

       // insert data
       // Using txOrm to execute SQL
       _, err = to.Insert(user)

       if err != nil {
           logs.Error("execute transaction's sql fail, rollback.", err)
           err = to.Rollback()
           if err != nil {
               logs.Error("roll back transaction failed", err)
           }
       } else {
           err = to.Commit()
           if err != nil {
               logs.Error("commit transaction failed.", err)
           }
       }

.. code:: go

   o := orm.NewOrm()
   to, err := o.Begin()

   // outside the txn
   o.Insert(xxx)

   // inside the txn
   to.Insert(xxx)


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