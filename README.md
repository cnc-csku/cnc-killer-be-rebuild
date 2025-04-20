# cnc-killer-be-rebuild
for new cnc killer rebuild in go
## How to run project
first clone this project and install dependencies using 
```bash
go mod tidy 
```

then you have to create `.env` file in your local machine we have some example in `.env.example` you can custom it as you want

or if you want to try our project immediately you can try this
```bash
cp ./.env.example ./.env
```

after you have create `.env` you can try
```bash
docker compose up -d --build
```
if you want to close thi project you can try 
```bash
docker compose down
```

## Migrations
if your project have migrations you can migrate your database using 
```bash
make migrate-all
```
to migrate all of your migration


or if you have update your migration so you can do this 
```bash
make migrate-up # for migrate up 1 step
# and 
make migrate-down # for migrate down 1 step
```

if you want to create your own migrations you can try this
```bash
make migrate-create name="your migration name"
```

### Tips 
in our project use something calls  `Hexagonal architecture` this is architecture that help us to decouple our code as module

in this project contains 2 important directory 

1) `core` directory , use to store `port` in hexagonal architecture in this case is using to store `interface` in core contain's 3 main sub directory 
    
    - `repository` is use to store `secondary-port` or store interface that use for `core` activity in our application
    - `services` is use to store `primary-port` or store interface of `business logic` of our application
    - `models` is use to defined how to contract between port
2) `adaptors` directory, use to store implementation of interface and store "adaptors" to tranfer data through each level of port
