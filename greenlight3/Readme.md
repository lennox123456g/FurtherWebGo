we’re going to work through the start-to-finish build of an application called
Greenlight — a JSON API for retrieving and managing information about movies. You can
think of the core functionality as being a bit like the

TOOLS TO NOTE
1.HEY TOOL FOR LOADTESTING 
2.CURL TOOL FOOL FOR working with http request in the terminal


SECTION 1
Create a skeleton directory structure for the project and explain at a high-level how our
Go code and other assets will be organized.
Establish a HTTP server to listen for incoming HTTP requests.
Introduce a sensible pattern for managing configuration settings (via command-line
flags) and using dependency injection to make dependencies available to our handlers.
Use the httprouter package to help implement a standard RESTful structure for the API
endpoints

MODULE PATH an enabling modules 
go mod init 
we add our pthe
go mod init greenlight.lennoxmugumira12.net

LINK TO GO MODULES 
https://go.dev/wiki/Modules


The bin directory will contain our compiled application binaries, ready for deployment
to a production server.
The cmd/api directory will contain the application-specific code for our Greenlight API
application. This will include the code for running the server, reading and writing HTTP
requests, and managing authentication.
The internal directory will contain various ancillary packages used by our API. It will
contain the code for interacting with our database, doing data validation, sending emails
and so on. Basically, any code which isn’t application-specific and can potentially be
reused will live in here. Our Go code under cmd/api will import the packages in the
internal directory (but never the other way around).
The migrations directory will contain the SQL migration files for our database.
The remote directory will contain the configuration files and setup scripts for our
production server.
The go.mod file will declare our project dependencies, versions and module path.
The Makefile will contain recipes for automating common administrative tasks — like
auditing our Go code, building binaries, and executing database migrations.

Loogger

Putting it together: why config struct + logger work as a pair
•	The config struct holds settings — how the app should behave (which port, which environment, which DB)
•	The logger holds the means to report what the app is doing while it runs
A logger gives you structured, timestamped, traceable output instead of plain fmt.Println. Compare:
fmt.Println("server started")
•	vs.
•	go
•	infoLog.Println("server started
