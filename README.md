# personacmms
Personal computer maintenance management system
* An application to aid in tracking, recording, and reminding me about periodic maintenance for all the things I don't currently maintain properly around my house.

### Developer Environment Setup

- wsl setup

- vscode install
    - extensions
    - debugging
- postgresql install
    - apt
- postgresql setup
    - https://neon.tech/postgresql/postgresql-getting-started/install-postgresql-linux
    - db creation - personacmms-test, personacmms
    - environment variable configuration


### Build


### Test


### Deploy

projects found within: 
### restapi: A REST api to serve state and run background tasks on the data as business rules require
* Stack:
    * golang (gin for httpapi layer)
    * PostgreSQL

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
    * add integration with TODOist - can be used as an alternative UI (easier access less feature rich UI)
        * implement feature flag
        * implement monitoring for work order completion
        * implement api integration for creating TODOs 

* Future Features
    * Add asset tagging and filtering based on tag


### webui: a web based interface.
* Stack: 
    * CSS: TailwindCSS + DaisyUI + HeroIcons
    * JS Framework: SvelteKit

* TODO
    * Complete design in Figma
    * Build out design in SvelteKit app
    * Make UI responsive for mobile rendering

### cli: a command line based interface, may not implement, tbd..
