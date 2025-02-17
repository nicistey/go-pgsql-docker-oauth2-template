# Описание проекта

**Этот проект создан для того, чтобы его можно было использовать как шаблон для начала разработки веб-приложений на Go. В проекте реализованы основные функции, такие как:**

- Связь с PostgreSQL
- Кэширование запросов с помощью Redis, что позволяет ускорить ответы API и снизить нагрузку на базу данных
- Аутентификация с использованием JWT
- OAuth авторизация через Google
- CRUD операции для пользователей и событий
- Middleware для обработки запросов
- Настройка CORS политики

**Настроенный шаблон клиента для подключения к серверу**

[Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt)

[Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt)

[Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt)

[Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt) [Frontend](https://github.com/nicistey/vue-event-manager-jwt)



## Содержание

- [Начало работы](#начало-работы)
  - [Клонирование репозитория](#клонирование-репозитория)
  - [Настройка окружения](#настройка-окружения)
  - [Получение ключей для .env](#получение-ключей-для-файла-env)
  - [Запуск в Docker](#запуск-в-docker)
  - [Дополнительные шаги](#дополнительные-шаги)
- [Примеры API запросов](#примеры-api-запросов)
- [Описание работы кода](#описание-работы-кода)
  - [Основные компоненты](#основные-компоненты)
  - [Пошаговое описание работы](#пошаговое-описание-работы)
- [Разработка в будущем](#разработка-в-будущем)

# Начало работы
## Клонирование репозитория
Для начала работы необходимо клонировать репозиторий:
```sh
git clone https://github.com/nicistey/go-pgsql-docker-oauth2-template
cd go-pgsql-docker-oauth2-template
```
## Настройка окружения
Откройте файл .env в директории Server\cmd и рассмотрите в не следующие переменные:
```.env
CONN_STR_DB= "postgres://***:***@postgres:5432/***"
GOOGLE_CLIENT_ID="***.apps.googleusercontent.com"
GOOGLE_CLIENT_SECRET="***"
GOOGLE_REDIRECT_URL="http://localhost:6080/auth/callback"
JWT_SECRET_KEY="***"
```
## Получение ключей для файла .env

  **1. CONN_STR_DB**: Строка подключения к базе данных PostgreSQL. Замените username, password, hostname, port и dbname на ваши данные.
<details> <summary>Получение CONN_STR_DB</summary>
Убедитесь, что у вас установлен PostgreSQL.
Сформируйте строку подключения в формате: postgres://username:password@hostname:port/dbname.
  
  **Что бы покдючиться через PgAdmin**
  
  Создайте сервер
  
  ![FirstPgsql](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/FirstPgsql.png)
  
  Подключитесь,  [после запуска проекта в докере](#запуск-в-docker)
  
  ![SecondPgsql](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/SecondPgsql.jpg)
  
</details>

  **2. GOOGLE_CLIENT_ID** и **GOOGLE_CLIENT_SECRET**: Для получения этих ключей необходимо зарегистрировать ваше приложение в Google API Console. Создайте OAuth 2.0 Client ID и получите необходимые ключи.
<details> <summary>Получение GOOGLE_CLIENT_ID и GOOGLE_CLIENT_SECRET </summary>

1) Перейдите на [Google API Console](https://console.cloud.google.com/cloud-resource-manager). 
  
2) Создайте новый проект или выберите существующий.
3) У проекта нажмите на три точки у проекта и перейдите в раздел Settings
4) Перейдите в Navigation menu (три полоски у лого Console Logo) и перейдите в раздел APIs & Services
5) Перейдите в раздел "OAuth consent screen".
6) Выберите тип External

![External](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/External.png)

7)  Заполните необходимые поля App name, Support email, Contact email addresses
8)  После всех действий перейдите в раздел "Credentials".
  
  ![Credentials](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/Credentials.png)

9)  Выберите Application type и тип Web application
  
  ![Application](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/Application.png)

10)  В разделе Authorized redirect URIs укажите http://localhost:6080/auth/callback
     (в случае когда будете запускать сервер вне докера, не забудьте поменять порт)
  
  ![redirect](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/Redirect.png)

11) Вы получите окно с необходимыми данными Client ID (GOOGLE_CLIENT_ID) и Client secret (GOOGLE_CLIENT_SECRET)
  
  ![Client](https://github.com/nicistey/images-for-projects/blob/main/go-pgsql-docker-oauth2-template-image/Client.jpg)
</details>

  **3. GOOGLE_REDIRECT_URL**: URL для перенаправления после успешной авторизации через Google. В нашем случае это  `http://localhost:6080/auth/callback`.

  **4. JWT_SECRET_KEY**: Секретный ключ для подписи JWT токенов. Вы можете сгенерировать его самостоятельно.
<details> <summary>  Получение JWT_SECRET_KEY</summary>
Сгенерируйте секретный ключ для подписи JWT токенов. Вы можете использовать любой генератор случайных строк или команду в терминале:

  ```bash
  openssl rand -base64 32
```
  Пример
  ` 1a+Aa/1a+lA1a1A1/aA1+a11a/AA+a1a+a1aA1+a1a1+a1a1+a/a1A1+a1a1a/a1a1a+a1a1+a/a1A1+a1a1a+a1a1/a1a1a+a/a1a1+a1a1+a/a1a1+a1a1+a/a1a1 `
</details>

## Запуск в Docker
Для запуска проекта в Docker выполните следующие команды:

  1.Постройте Docker образы:
```docker
docker-compose build
```
  2.Запустите контейнеры:
```
docker-compose up
```
## Дополнительные шаги
- Убедитесь, что у вас установлен Docker и Docker Compose.
- Проверьте, что порты 6080 и 5433 не заняты другими приложениями.
# Примеры API запросов
Получение всех пользователей: `GET /api/users`

Получение пользователя по ID: `GET /api/users/{IDus}`

Создание нового пользователя: `POST /api/users`

Обновление пользователя: POST `/api/users/{IDus}`

Удаление пользователя: DELETE `/api/users/{IDus}`

Получение всех событий: GET `/api/events`

Получение событий по ID пользователя: `GET /api/eventsByID`

Создание нового события: `POST /api/events`

Обновление события: `POST /api/events/{IDev}`

Удаление события: `DELETE /api/events/{IDev}`

# Описание работы кода
Этот сервер написан на языке Go и предназначен для работы с веб-приложением, используя PostgreSQL в качестве базы данных и JWT для аутентификации. Ниже приведено пошаговое описание работы сервера.

## Основные компоненты
**1. Файл конфигурации** `(config/config.go)`:
-  Загружает переменные окружения из файла .env.
-  Проверяет наличие всех необходимых переменных.
-  
**2. Основной файл** `(cmd/main.go)`:
-  Загружает конфигурацию.
-  Подключается к базе данных PostgreSQL.
-  Инициализирует маршрутизатор и обработчики.
-  Запускает сервер.
  
**3. Маршрутизатор и обработчики** `(pkg/api/api.go)`:
-  Определяет маршруты для различных API-эндпоинтов.
-  Включает middleware для обработки запросов и аутентификации.
  
**4. Middleware** `(pkg/api/middleware.go)`:
-  Логирует запросы.
-  Обрабатывает CORS-заголовки.
-  Проверяет JWT токены для защищенных маршрутов.
  
**5. Обработчики для пользователей** `(pkg/api/users.go)`:
-  Обрабатывают запросы для создания, получения, обновления и удаления пользователей.
  
**6. Обработчики для событий** `(pkg/api/events.go)`:
-  Обрабатывают запросы для создания, получения, обновления и удаления событий.
  
**7. Аутентификация через Google** `(pkg/api/authGoogle.go)`:
-  Обрабатывает OAuth2 авторизацию через Google.
-  Генерирует JWT токены для аутентифицированных пользователей.
  
## Пошаговое описание работы
**1. Запуск сервера**:
-  В `main.go` вызывается функция `config.LoadConfig()`, которая загружает конфигурацию из файла `.env`.
-  Затем вызывается `repository.New(cfg.DBConnString)`, чтобы подключиться к базе данных PostgreSQL.
-  Инициализируется маршрутизатор `mux.NewRouter()` и создается новый экземпляр API с помощью `api.New(mux.NewRouter(), db, cfg)`.
-  Определяются маршруты и middleware с помощью `api.Hadle(cfg)`.
-  Сервер запускается на порту 8090 с помощью `api.ListenAndServe("0.0.0.0:8090")`.
  
**2. Маршруты и обработчики**:
-  Маршруты определяются в `api.go` с помощью `api.r.HandleFunc`.
-  Например, маршрут для получения всех пользователей: `api.r.HandleFunc("/api/users", api.getAllUsers).Methods(http.MethodGet,http.MethodOptions)`.
  
**3. Middleware**:
-  Middleware определен в `middleware.go` и используется для обработки всех запросов.
-  Он логирует запросы, обрабатывает CORS-заголовки и проверяет JWT токены для защищенных маршрутов.
  
**4. Обработка запросов**:
-  Когда поступает запрос, маршрутизатор передает его соответствующему обработчику.
-  Например, запрос `GET /api/users` обрабатывается функцией `getAllUsers` в` users.go`.
-  Обработчик выполняет необходимые действия (например, запрос к базе данных) и возвращает ответ клиенту.
  
**5. Аутентификация через Google**:
-  Запрос на `/auth` перенаправляет пользователя на страницу авторизации Google.
-  После успешной авторизации Google перенаправляет пользователя на `/auth/callback`, где сервер обрабатывает ответ и генерирует JWT токен.

# Разработка в будущем
-  Добавить тесты для всех основных функций.
-  Реализовать логирование запросов и ошибок.
-  Добавить поддержку других баз данных.
-  Улучшить обработку ошибок и валидацию данных.
-  Этот проект предоставляет базовую структуру для разработки веб-приложений на Go и может быть расширен в соответствии с вашими требованиями
