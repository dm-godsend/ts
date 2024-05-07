# ts
Timestamp converter console command

## Install
Run from project root
```sh
go install .
```

## Run
Run `ts` command anywhere
```sh
# convert int arg to utc date
ts 1718744340 # 2024-06-18 20:59:00 +0000 UTC

# convert utc date arg to start ts of day
ts 2024-06-18 # 1718668800

# convert utc datetime args to ts
ts 2024-06-18 20:59:00 # 1718744340
```