### Start
* Open a terminal window, and switch to the project root
    * $> go run main.go
    * Expectations
        * Golang is installed on the system
        * GO "mux" package is installed
            * $> go install mux
### Test the endpoint "lza/bag" (HTTP POST) 
* Open another terminal window
    * $> curl -X POST http://127.0.0.1:8080/lza/bag -F "bagit_file=@my_bag.zip" -vvv
    * Expectations
        * A file my_bag.zip in the working directory
        * Server is already runnning
### View API
* Open another terminal window, and switch to the project root
    * $> ./start.sh
* open: http://localhost
    * select: File > Clear editor
    * Copy the spec to the editor 
        * spec: project_root > api > swagger.ymal
    

    
