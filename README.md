# ikon_optimalDeviceConfiguration

There way to run this project is in command line

1.- There is necessary to run it in a normal way
    go run main.go

2.- There is necessary to include the command "configure"
    go run main.go configure
    
3.- We need to add the file input path as the help shows
    go run main.go configure --filePath {filePath}
    
4.- Check the output File and content

5.- Optional.  There is a flag that indicates the delta configuration for device confis.
You can use this variable if you want to add Possible Ids that the sum is not exact for the device capacity
but the most close will be included
    go run main.go configure --filePath {filePath} --delta {int>0}