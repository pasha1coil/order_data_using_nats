# FIOAPI
API по определения гендера, пола и национальности по указанному фио
<details>
  <summary>Содержание</summary>
  <ol>
    <li><a href="#установка-и-запуск">Установка и запуск</a></li>
    <li><a href="#реализовано">Реализовано</a></li>
    <li><a href="#примеры запросов">Возникшие вопросы</a></li>
  </ol>
</details>

## Установка и запуск

Клонировать проект.
Далее через `makefile`:
- `make build` - собрать проект
- `docker-compose up` - запустить проект
- `make test` - запустить тесты
- `migrate -path ./db -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up` - инициализировть бд, подставить свои данные, также изсенить в .env

## Реализовано

1. Метод добавления пользователей списком json.
2. Метод получения данных о пользователях по возрасту, также больше меньше выставленного возраста со знаками +,-.
3. Метод получения данных о пользователях по гендеру.
4. Метод получения данных о пользователях по национальности.
5. Метод получения данных о добавленных пользователях из кеша Redis по id.
6. Метод обновления данных пользователя по двум параметрам (в примерах подробнее).
7. Метод удаления пользователя.

## Примеры запросов

Для взаимодействия с сервером есть 7 способов:
1. (POST) /main/adduser
2. (GET) /main/getage
3. (GET) /main/getgender
4. (GET) /main/getnationality
5. (GET) /main/getuserdata
6. (PUT) /main/user
7. (DELETE) /main/user 

В качестве ответа возвращается JSON с данными.

Примеры запросов и ответов:
1. http://localhost:8080/main/adduser POST:
  - INPUT
  ```json
    {
      "fio":[
          {
              "name":"Alisa",
              "surname":"Parker"
          },
          {
              "name":"Alex",
              "surname":"Parker"
          },
          {
              "name":"Oleg",
              "surname":"Parker"
          },
          {
              "name":"Yulia",
              "surname":"Pribilskaya",
              "patronymic":"Parker"
          }
      ]
  }
  ```
  - OUTPUT
  ```json
  {
    "Message": [    //возвращает id созданных записей
        "1",
        "2",
        "3",
        "4"
    ]
  }
  ```

2. http://localhost:8080/main/getage GET:
  - INPUT
  ```json
  {
    "age":45,    //znak можно не инициализировать -> записать просто {"age":45} - тогда ответ будет удовлетворять только тем пользователям
    "znak":"+"   //у которых возраст равен 45, если "znak":"-" -> ответ меньше и равно 45
  }
  ```
  - OUTPUT
  ```json
  {
    "Status": [
        {
            "id": 2,
            "name": "Alex",
            "surname": "Parker",
            "patronymic": "",
            "gender": "male",
            "age": 46,
            "Nationality": [
                {
                    "Country": {
                        "Country_id": "CZ"
                    }
                },
                {
                    "Country": {
                        "Country_id": "UA"
                    }
                },
                {
                    "Country": {
                        "Country_id": "RO"
                    }
                },
                {
                    "Country": {
                        "Country_id": "RU"
                    }
                },
                {
                    "Country": {
                        "Country_id": "IL"
                    }
                }
            ]
        },
        {
            "id": 3,
            "name": "Oleg",
            "surname": "Parker",
            "patronymic": "",
            "gender": "male",
            "age": 53,
            "Nationality": [
                {
                    "Country": {
                        "Country_id": "UA"
                    }
                },
                {
                    "Country": {
                        "Country_id": "RU"
                    }
                },
                {
                    "Country": {
                        "Country_id": "BY"
                    }
                },
                {
                    "Country": {
                        "Country_id": "MD"
                    }
                },
                {
                    "Country": {
                        "Country_id": "IL"
                    }
                }
            ]
        }
    ]
  }
  ```

3. http://localhost:8080/main/getgender GET:
  - INPUT
  ```json
  {
    "gender":"male" // или "gender":"female"
  }
  ```
  - OUTPUT
  ```json
  {
    "Status": [
        {
            "id": 2,
            "name": "Alex",
            "surname": "Parker",
            "patronymic": "",
            "gender": "male",
            "age": 46,
            "Nationality": [
                {
                    "Country": {
                        "Country_id": "CZ"
                    }
                },
                {
                    "Country": {
                        "Country_id": "UA"
                    }
                },
                {
                    "Country": {
                        "Country_id": "RO"
                    }
                },
                {
                    "Country": {
                        "Country_id": "RU"
                    }
                },
                {
                    "Country": {
                        "Country_id": "IL"
                    }
                }
            ]
        },
        {
            "id": 3,
            "name": "Oleg",
            "surname": "Parker",
            "patronymic": "",
            "gender": "male",
            "age": 53,
            "Nationality": [
                {
                    "Country": {
                        "Country_id": "UA"
                    }
                },
                {
                    "Country": {
                        "Country_id": "RU"
                    }
                },
                {
                    "Country": {
                        "Country_id": "BY"
                    }
                },
                {
                    "Country": {
                        "Country_id": "MD"
                    }
                },
                {
                    "Country": {
                        "Country_id": "IL"
                    }
                }
            ]
        }
    ]
  }
  ```

4. http://localhost:8080/main/getnationality GET -> принимает на вход список UserID и Countryid:
  - INPUT
  ```json
  {
    "UserID":[1,2,3],
    "Countryid":"BY"
  }
  ```
  - OUTPUT
  ```json
  {
    "Status": [
        {
            "UserID": 3,
            "Countryid": "BY",
            "UserData": [
                {
                    "id": 3,
                    "name": "Oleg",
                    "surname": "Parker",
                    "patronymic": "",
                    "age": 53,
                    "gender": "male"
                }
            ]
        }
    ]
  }
  ```

5. http://localhost:8080/main/getuserdata GET:
  - INPUT
  ```json
  {
    "id":1
  }
  ```
  - OUTPUT
  ```json
    ["Alisa","Parker","","female",42,["RU","UA","CN","BA","TH"]]
  ```

6. http://localhost:8080/main/user PUT:
  - INPUT
  ```json
  {
    "id":1,
    "name":"Irina" //можно вводить ТОЛЬКО один из этих - "name" - string,"surname" - string,"patronymic" - string,
  }               // "gender"- string,"age"- int параметров, либо измениться только первый параметр
  ```
  - OUTPUT:
  ```json
  {
    "Status": "The update was successful"
  }
  ```

7. http://localhost:8080/main/user DELETE:
  - INPUT
  ```json
  {
    "id":1
  }
  ```
  - OUTPUT
  ```json
  {
    "Status": "Successful removal"
  }
  ```