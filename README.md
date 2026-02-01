# URL Shortener

Сервис для сокращения длинных URL и получения оригинального адреса по короткой ссылке. Пет-проект на Go с послойной архитектурой.

## Требования

- Go 1.25+

## Запуск

```bash
go run ./cmd/url-shortener
```

Сервер слушает порт **7000**.

## API

Все запросы и ответы — JSON. Обязательный заголовок: `Content-Type: application/json`.

### Сократить URL

**GET** `/shorten`

**Тело запроса:**
```json
{
  "original_url": "https://example.com/very-long-path"
}
```

**Ответ 200:**
```json
{
  "shortened_url": "короткая ссылка"
}
```

**Ошибки:** 400 — пустой или невалидный URL, конфликт короткой ссылки; 500 — внутренняя ошибка.

### Получить оригинальный URL

**POST** `/redirect`

**Тело запроса:**
```json
{
  "shortened_url": "короткая ссылка"
}
```

**Ответ 200:**
```json
{
  "original_url": "https://example.com/very-long-path"
}
```

**Ошибки:** 400 — пустой или невалидный short URL; 404 — короткая ссылка не найдена; 500 — внутренняя ошибка.

## Примеры (curl)

```bash
# Сократить URL
curl -X GET http://localhost:7000/shorten \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://example.com/long-url"}'

# Получить оригинал
curl -X POST http://localhost:7000/redirect \
  -H "Content-Type: application/json" \
  -d '{"shortened_url": "short-url"}'
```

## Структура проекта

```
cmd/url-shortener/     # Точка входа
internal/
  delivery/http/       # HTTP-хендлеры, роутер
  domain/             # Сущности и ошибки домена
  infrastructure/     # In-memory хранилище
  usecase/            # Бизнес-логика (shorten, redirect)
pkg/http/             # HTTP-сервер
```

Данные хранятся в памяти и теряются при перезапуске.

## Тесты

```bash
go test ./...
```
