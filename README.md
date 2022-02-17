# ikon_optimalDeviceConfiguration

There way to run this project is in command line

1.- There is necessary to run it in a normal way
    go run main.go
![alt text](https://github.com/GerardoZunigaItexico/ikon_optimalDeviceConfiguration/blob/main/readmeFiles/1_goRunMain.png)

2.- There is necessary to include the command "configure"
    go run main.go configure
    ![alt text](https://github.com/GerardoZunigaItexico/ikon_optimalDeviceConfiguration/blob/main/readmeFiles/2_goRunMainConfigure.png)
    
3.- We need to add the file input path as the help shows
    go run main.go configure --filePath {filePath}
    ![alt text](https://github.com/GerardoZunigaItexico/ikon_optimalDeviceConfiguration/blob/main/readmeFiles/3_goRunMainConfigureFilePath.png)
    
4.- Check the output File and content
    ![alt text](https://github.com/GerardoZunigaItexico/ikon_optimalDeviceConfiguration/blob/main/readmeFiles/4_goRunMainConfigureFilePath_1.png)
    ![alt text](https://github.com/GerardoZunigaItexico/ikon_optimalDeviceConfiguration/blob/main/readmeFiles/4_goRunMainConfigureFilePath_2.png)

5.- Optional.  There is a flag that indicates the delta configuration for device confis.
You can use this variable if you want to add Possible Ids that the sum is not exact for the device capacity
but the most close will be included
    go run main.go configure --filePath {filePath} --delta {int>0}
    ![alt text](https://github.com/GerardoZunigaItexico/ikon_optimalDeviceConfiguration/blob/main/readmeFiles/5_goRunMainConfigureFilePath_and_delta_1.png)
