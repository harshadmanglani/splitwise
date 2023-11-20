# splitwise

## Requirements
Create a backend application that allows users to do the following tasks
1. Create a user account
2. Log in and get a JWT
3. Search for other users in the database
4. Add an expense with one or more users where
   1. they can choose who paid
   2. who owes how much in
      1. percentages
      2. absolute amounts
   3. they can exclude a particular user from the expense altogether
5. Create an expense group of users
6. Simplify payments in a group


## Future scope
1. Add reminders
2. Add multi-currency support
3. Integrate with collect flows for P2P (Gpay, PayTM)
4. Add update history in expenses API

## Database schema



## APIs
`POST /users`
```
Request:
{
    "username": "harshad.manglani",
    "passHash": "password",
    "email": "harshad.manglani@example.com",
    "name": "Harshad Manglani",
    "phone": "9923225282"
}

Response:
201 Created

Error responses:
409 Conflict
{
    "message": "USERNAME_ALREADY_EXISTS" // EMAIL_ALREADY_EXISTS, PHONE_ALREADY_EXISTS
}
```

`POST /users/login`
```
Request:
{
    "username": "harshad.manglani",
    "passHash": "password"
}

Response:
200 OK
{
    "data": {
        "access_token": "eyJhbGciOjAsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiYWNrZW5kIiwic3ViIjoiODQ2MjZiZjktZDVkNS00MDFmLTkzMDQtMDRiOWFmNWE0YTQ5IiwiZXhwIjoiMjAyMy0xMS0xMlQxMzoxMDoyNS41NjcwOTErMDU6MzAiLCJpYXQiOiIyMDIzLTExLTEyVDEzOjA4OjI1LjU2ODcxOSswNTozMCIsImN1c3RvbSI6bnVsbH0=.z9KIf30INSFACl0FsVNYphCLebRPapCMO/6qbXZ1sJI=",
        "user": {
            "id": 9,
            "createdAt": "2023-11-12T13:04:34.372747+05:30",
            "updatedAt": "2023-11-12T13:04:34.372747+05:30",
            "uuid": "84626bf9-d5d5-401f-9304-04b9af5a4a49",
            "username": "harshad.manglani",
            "email": "harshad.manglani@example.com",
            "name": "Harshad Manglani",
            "phone": "9923125184"
        }
    }
}

Error responses:
404 Not Found
{
    "message": "INVALID_USER_OR_PASSWORD"
}
```
`POST /expenses`
```
```

### Goals
1. Write to code to optimize for garbage collection
2. Understand GoRoutines and implement them to handle scale and check for race conditions

### Open Points
1. How do I implement user authentication with JWTs in the header and annotations?