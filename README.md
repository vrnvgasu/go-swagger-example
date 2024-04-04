1. клиент отправляет картинки и принимает pdf
2. сервер принимает картинки и возвращает pdf

3. Генерация кода через swagger
##### Генерация кода на сервере на основе `server/api/api.yml`
```
swagger generate server -f ./api/api.yml --server-package=internal/handler -A "service-pdf-compose"
```

##### Генерация кода на клиенте на основе `server/api/api.yml`
```
swagger generate client -f ../server/api/api.yml -A "controller-service"
```

4. На сервере из сгенерированного кода можно изменять только server/internal/handler/configure_service_pdf_compose.go
5. Форма для загрузки картинок на клиенте http://localhost:8080
6. На клиенте пока обрабатывается только первый файл из трех 