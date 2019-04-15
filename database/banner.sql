-- структура таблицы баннеров
CREATE TABLE banners
(
    id SERIAL,
    title VARCHAR(50) NOT NULL ,
    brand VARCHAR(50) NOT NULL,
    size VARCHAR(50) NOT NULL,
    active BOOLEAN,
    created_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY(id)
)