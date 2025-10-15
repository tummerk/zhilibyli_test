# **Blog на laravel**

# Routes

1. GET http://localhost:80/api/posts - возвращает все посты
2. POST http://localhost:80/api/posts - создает пост, поля title и body
3. PUT/PATCH http://localhost:80/api/posts{id} - обновление
4. GET http://localhost:80/api/posts/{id} - 1 пост
5. DELETE http://localhost:80/api/posts/{id} - удалить пост
6. GET http://localhost:80/api/posts/count - кол-во постов

# Запуск

1) composer install
2) cp .env.example .env
3) docker-compose up -d

# INFO

User`ов нету не знал надо или нет( всё сделано только для одного блога
т.к. без авторизации пришлось бы пользоваться всякими костылями по типу передачи user_id в headers или токенами

миграции стандартно php artisan migrate 

есть seeder 

проект собран в sail

RMB не использовался чтобы не нарушать solid
