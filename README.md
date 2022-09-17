# snapp
Snapp trip ticket engine
 
 
# Build with
* Go
* Gin
* Gorm
* Postgres
* Redis

# How to start
First, named technologies in Build with section should been installed in your computer. </br>

Then you need to config database with project. </br>
So you can change your database dsn and redis configuration in `database.go` in initializers folder. </br>

After connecting postgres and redis to project you need to migrate tables in database. For that you should run `migrate.go` in migrate folder. </br>
You can run it by ``` go run migrate/migrate.go ```  command in project folder in terminal.</br>

Afterward, you need to import data in Cities, Suppliers, Agencies and Airlines table in database. </br>
Given data has been added besides the project folder in `data` folder. For importing csv data into database I highly recommend using pgadmin. </b>
For the default data you just need to import them in name Column of tables (ID column automatically would be added).</br>
</br>
At the end you just need to run `server.go`. </br>
Congrats your api is ready to use </br>

# How to use
Defenition of project is available in `project_defenition.pdf` . </br>
request and response formats are also available in `request-response` folder. </br>
For creating rules you can post your request to ```localhost:8080/createRule``` and for changePrice you can post your request to ```localhost:8080/changePrice``` .</br>
Be careful api has validation on data so only valid data could be applied.
