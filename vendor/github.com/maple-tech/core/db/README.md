# Core - Database (SQL)

Maintains the connection to the PostgreSQL database using the `jmoiron/sqlx` package. Besides connection pool management, it also provides some helpful functions for transactions and error checking that wraps up most of the provided functionality from the `sqlx` and standard `sql`  packages. 

## Usage

Like most packages, the connection must be initialized with the configurations using `db.Initialize()`. During the initialization the connection will be made, and tested as well using a ping for validity. After this point, other services are free to use the connection as they need. You can retrieve the raw connection using `db.GetConnection()` if needed.

To help out with queries, utility functions are provided (they are just wrapped from `sqlx`). Such as: `db.Get()`, `db.Select()` and `db.Exec()`. 

You can also wrap transactions for connection optimization using the `db.WrapTX()` function. Wrapping the transaction will automatically do error recover from panics or otherwise, which will `rollback` the transaction, or if no error is returned then `commit` to the database. **It is recommended to use the WrapTX function whenever more then 2 queries happen in a single process**

Error checking can be helped with the `db.CheckError()` and `db.IsEmptyError()` functions. The `CheckError` function will check the returned error against the code provided to see if they match. This is useful since it seems every SQL package relies on 3 others with their own error types. The `IsEmptyError` function returns `true` if the error is because no rows where returned (empty set).