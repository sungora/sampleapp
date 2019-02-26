# sample
Пример приложения (web, api, workflow)

#### get app 
Скачиваем приложения
    
    git clone git@github.com:sungora/sample.git

#### Зависимости и сборка

    make help       Справка по командам
    make dep        Получение и обновление необходимых зависимостей
    make build      Сборка проекта
    make run        Сборка и запуск проекта 

#### developer

Конфигурационные файлы вашего модуля-приложения
должны начинаться с названия самого моудля (sample) к кторому они относятся.
И находится в папке config 

Бланки конфигураций находятся в папке docs/config_blank

Структура проекта:

    /api        документация на api
    /bin        испольняемые файлы проекта (.gitignore)
    /build      сборка и деплой проекта
    /cmd        точки входа в приложения (main, golang)
    /config     конфигурации приложений проекта (.gitignore)
    /docs       общая документация проекта
    /log(s)     логи работы приложения (.gitignore)
    /tools      служебный инстурментарий
    /internal   исходный код проекта (golang)
      /modname  модуль приложения
      ...
    /vendor     сторонние библиотеки (зависимости, golang)
    /www        статика и клиентская часть (html, css, js, img)
     
Построение роутига api:

<pre>
/api/v1
     /users
         /{userID}
             /orders
                 /{orderID}
                 ...
</pre>
**Рекомендация:**

Структура пакетов обработчиков доложна совпадать с роутингом api


