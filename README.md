## Overview

The repository of notification service

## Endpoints

Method | Path                             | Description                                   |                                                                         
---    |----------------------------------|------------------------------------------------
GET    | `/health`                        | Health page                                   |
GET    | `/metrics`                       | Страница с метриками                          |
GET    | `/v1/notifications/list`         | Получение всех уведомлений                    |
GET    | `/v1/notifications/get/{userId}` | Получение всех уведомлений по id пользователя |