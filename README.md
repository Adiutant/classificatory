# Это учебный проект в стадии pre-MVP
Его можно запустить без dev-окружения
# Описание 
Проект представляет собой кластер сервисов, задачей которых является ранжирование текстов по темам. Сервисы поддерживают авторизацию пользователей, учетные данные сохраняются в mock-БД. 
# Структура 
![1](https://github.com/Adiutant/classificatory/assets/17684112/d80eb66d-d3cf-4f8c-ad98-b6baee851d9e)

В кластере следующие сервисы:
  1. Gin сервер-маршрутизатор
  2. Python скрипт-классификатор
  3. Mock-БД

# Развертывание
Для развертывания используется оркестрация docker-compose. В папке make:
`docker-compose up`

# Использование

/authenticate
`{
    "username": "{username}",
    "password" : "{password}"  
}`
response -> 
`{
  "token" : "{JWT}"
}`
Запрос на получение JWT токена для учетных данных username , password

/request-payload POST HEADER Authorization {TOKEN}
`{
    "command" : "{Command}",
    "payload" : "{Payload}"
}`
response -> 
`{
   "command" : "{Command}",
    "payload" : "{Payload}"
}`

Запрос на выполнение Command с Payload, в headers указывается JWT токен TOKEN
Список команд для Command:
  1. Text  -  Ранжирование Payload текста
  2. Add - Добавление темы Payload
  3. Delete - Удаление темы Payload
  4. List - Список тем (Payload не используется)

