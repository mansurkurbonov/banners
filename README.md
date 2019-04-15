# banners
Простой rest api на GO

1) Созать таблицу banners(./database/banner.sql), СУБД Postgres
2) настроить конфигурации(./configs/dev.json) сервера и бд
3) запустить проект. go run main go
4) слушать порт указанный в конфиге

Апи-роуты
1) сохранения баннера
POST http://127.0.0.1:port/api/v1/banner

Headers
Content-Type - application/json
Authorization - apiKey указанный в конфиг файле

Body {
  "title": "LFC",
  "brand": "LIVERPOOL",
  "size": "300x400",
  "active": true
}

Ответ {
    "code": 200,
    "message": "OK",
    "payload": null
}

2) удаления баннера
DELETE http://127.0.0.1:port/api/v1/banner/:id


Headers
Content-Type - application/json
Authorization - apiKey указанный в конфиг файле

Ответ {
    "code": 200,
    "message": "OK",
    "payload": null
}

3) поиск баннеров
GET http://127.0.0.1:post/api/v1/banner?title=LFC

Headers
Content-Type - application/json

Ответ {
    "code": 200,
    "message": "OK",
    "payload": [
        {
            "id": 4,
            "title": "Wow1",
            "brand": "LIVERPOOL",
            "size": "300x400",
         }
     ]
 }
