# Генератор кода RESTful сервисов для Go по спецификации OpenAPI v3.

## Инструкция по использованию

В репозитории проекта создать Makefile:

```
run_generator:
    docker run -it -e SWAGGER_FILE=/$(SERVICE_NAME)/api/spec.yaml \
					-e SERVICE_NAME=$(SERVICE_NAME) \
					-e ROOT_PATH=/$(SERVICE_NAME) \
					--mount type=bind,src="$(PWD)",target=/$(SERVICE_NAME) panchesco13/gogen:latest /generate

init:
	go mod init $(SERVICE_NAME)

tidy:
	go mod tidy

lint:
	gofmt -w -s ./

clear:
	rm -rf cmd && rm -rf configs && rm -rf docs && rm -rf internal && rm -rf go.mod && rm -rf go.sum

build:
	go build -C cmd/app -o bin

generate: run_generator init tidy lint
```
1. Создать в корне репозитория директорию api, положить в нее спецификацию и переименовать на **spec.yaml**

2. Задать системную переменную **SERVICE_NAME** - наименование модуля проекта

3. Выполнить команду **make generate**

4. Задать системные переменные для подключения к БД (см. папку **configs**)
