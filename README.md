
For my personal learning purpose. 
anybody can use as reference with your own risk. 

The toolse that i use 
- goose https://github.com/pressly/goose
- bun for model https://github.com/uptrace/bun
- bun router https://github.com/uptrace/bunrouter

for my personal bun have a good sample implementation, but for me its still difficult to understand all of it. 

this project i created some of it based on my personal experience. 

Folder structure: 
- app 
    - database -> for put model
        - migrations -> sqlite db and migrations are here
    - http -> for put routing and dep injection for server start here
    - view -> read only function puts here. 
    - service.go -> interface that need to be implemented is put here 
    - task.go -> implementing service.go, create new file and update service.go for new implementation
- cmd 
    - api -> for running the apps is here. 


## How to run
`go run cmd/api/main.go`

will try update when i have time.
Plan to use this as my personal template for my new golang project in future. 