# personacmms
Personal computer maintenance management system
* An application to aid in tracking, recording, and reminding me about periodic maintenance for all the things I don't currently maintain properly around my house.

projects found within: 
### restapi: A REST api using gin/golang/postgres stack to manage the data
* TODO  
    * Implement proper error propagation up from store layer through app layer to httpapi layer
    * Implement proper app layer error handling in the httpapi layer (correct http status codes, etc.)
    * Implement integration testing for REST calls through all the layers
    * Implement app layer task frequency scanning / work order creation
    * create a dockerfile to containerize the app
    * add consistent logging
        * log levels (Err, Warn, Info, Debug)
        * log service response times
        * trace requests through layers


### webfe: a web based interface, probably will write it in REACT.. tbd
* TODO
    * Design UI using some design tool, maybe try out figma
    * Get started on the REACT app

### cli: a command line based interface, may not implement, tbd..
