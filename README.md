# goScoreboard

- Stores users in a leaderboard using redis
- nginx used as a load balancer to distribute traffic

``` / go / fiber / Docker / nginx /```

## Local setup

1. Clone the repo into your local
   ```
   git clone https://github.com/bcaglaraydin/go-scoreboard
   ```

3. Create an ```.env``` file and place it in root directory </br>
 Sample ```.env``` file:
    ```
    APP_PORT=":3000"
    REDIS_PORT=":6379"
    REDIS_HOST="redis"
    ```
4. Make sure you have docker-compose and docker installed on your machine

5. Run the following command (N is the number of servers)
   ```
   docker-compose up --scale api=N
   ```
## Endpoints

### Create User

#### Request

`POST http://localhost:3000/user/create`

#### Sample Request Body

   ```
{
        "user_id": "389e4264-e88b-12d3-a456-426614174000",
        "display_name": "username",
        "points": 1000,
        "rank": 0,
        "country": "tr"
    }
```

### Update User Score

#### Request

`POST http://localhost:3000/score/submit/{user_guid}`

#### Sample Request Body

   ```
{
    "score_worth": 10,
    "user_id": "399e4264-e88b-12d3-a456-426614174000",
    "timestamp": 8573040751
}
```

### Get User Profile

#### Request

`GET http://localhost:3000/user/profile/{user_guid}`

#### Sample Response Body

   ```
{
        "user_id": "389e4264-e88b-12d3-a456-426614174000",
        "display_name": "hey",
        "points": 1000,
        "rank": 1,
        "country": "gr"
}
```

### Get Leaderboard

#### Request

`GET http://localhost:3000/leaderboard`

#### Sample Response Body

   ```
[
    {
        "user_id": "389e4464-e88b-12d3-a456-426614174000",
        "display_name": "user",
        "points": 5000,
        "rank": 1,
        "country": "tr"
    },
    {
        "user_id": "389e4264-e88b-12d3-a456-426614174000",
        "display_name": "user2",
        "points": 1000,
        "rank": 2,
        "country": "gr"
    },
    {
        "user_id": "399e4264-e88b-12d3-a456-426614174000",
        "display_name": "user3,
        "points": 2,
        "rank": 3,
        "country": "tr"
    }
]
```

### Get Leaderboard Using Country Filter

#### Request

`GET http://localhost:3000/leaderboard/{country_iso_code}`

#### Sample Response Body

   ```
[
    {
        "user_id": "389e4464-e88b-12d3-a456-426614174000",
        "display_name": "user",
        "points": 5000,
        "rank": 1,
        "country": "tr"
    },
    {
        "user_id": "399e4264-e88b-12d3-a456-426614174000",
        "display_name": "user3,
        "points": 2,
        "rank": 3,
        "country": "tr"
    }
]
```


## Contact

Berdan Çağlar Aydın - https://www.linkedin.com/in/bcaglaraydin/ - berdancaglaraydin@gmail.com
