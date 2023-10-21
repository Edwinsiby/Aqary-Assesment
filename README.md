# Aqary-Assessment

Used gin framework - all servers are running on port 8080

1. Rate-Limit  - post-localhost:8080/limit/check
2. Command     - only test is added -  go test ./pkg/test
3. Custom-Bind - post-localhost:8080/bind, need to enter the json data with respect to models.CustomInput
4. DB-transaction - post - localhost:8080/user/data, need to enter the json data with respect to models.StudentData
                    added roll back for transaction and if any err happens the previous transactions will discarded
5. SQL-query   - text
6. MinCost     - command line execution
