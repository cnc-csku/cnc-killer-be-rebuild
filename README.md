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


