# splitwise

## Requirements
Create a backend application that allows users to do the following tasks
1. Create a user account/log in
2. Search for other users in the database
3. Add an expense with one or more users where
   1. they can choose who paid
   2. who owes how much in
      1. percentages
      2. absolute amounts
   3. they can exclude a particular user from the expense altogether

## Future scope
1. Add reminders
2. Integrate with collect flows for P2P (Gpay, PayTM)
3. Add update history in expenses API
4. Create a group of users where adding a new expense will by default split the amount equally
5. Simplify payments in a group

## APIs
```
POST /users/
Request Body:
{
    "email": "harshad@example.com",
    "phoneNumber": "9923124184",
    "firstName": "Harshad",
    "lastName": "Manglani",
    "passwordHash": "hash"
}

Response:
Code 201 Created
Body:
{
    "userId": "U1234"
}

### POST /users/login
Request Body:
{
    "email": "harshad@example.com",
    "phoneNumber": "9923125184", // either email or phoneNumber
    "passwordHash": "hash"
}

Response:
Code 200 OK
Body:
{
    "userId": "U1234"
    "token": "token",
    "expiry": 1201212121,
    "refreshToken": "refreshToken"
}

### GET /users/{userId}
Request Headers:
Authorization: Bearer <JWT>
Response:
Code 200 OK
Body:
{
    "totalBalance": -1000.34,
    "expenses": [],
    "balances": []
}

### GET /users/search?searchTerms=Harshad&Manglani
Request Headers:
Authorization: Bearer <JWT>
Response:
Code 200 OK
Body:
{
    "results": [
        {
            "userId": "U1234",
            "firstName": "Harshad",
            "lastName": "Manglani"
        }
    ]
}

### POST /expenses
Request Headers:
Authorization: Bearer <JWT>
Request Body:
{
    "userId": "U1234",
    "createdBy": "U1234, // must be same as userId for this call, immutable
    "amount": 100.34,
    "splitMode": "EQUAL", // PERCENTAGE, AMOUNT
    "splitBetween": [
        "U1234": {
            "paidFull": true,
            "amountOwed": 55, // based on enum value of splitMode
            "percentageOwed": 10 // based on enum value of splitMode
        }
        "U2345": {
            "amountOwed": 22,
            "percentageOwed": 80
        }
        "U3456": {
            "amountOwed": 13,
            "percentageOwed": 7
        }
        "U4567": {
            "amountOwed": 10.34,
            "percentageOwed": 3
        }
    ],
    "simplifyDebts": true // false
}

Response:
Code 201 Created
Body:
{
    "expenseId": "EX1234",
    "balances": [
        {
            "from": "U2345",
            "to": "U1234",
            "amount": 22
        },
        {
            "from": "U3456",
            "to": "U1234",
            "amount": 13
        },
        {
            "from": "U4567",
            "to": "U1234",
            "amount": 10.34
        }
    ],
    "createdBy": "U1234",
    "createdAt": 121291212,
    "updatedAt": 121291212
}

### PATCH /expenses/{expenseId}
Request Headers:
Authorization: Bearer <JWT>
Request Body:
{
    "userId": "U2345",
    "amount": 100.34,
    "splitMode": "EQUAL", // PERCENTAGE, AMOUNT
    "splitBetween": [
        "U1234": {
            "amountOwed": 55, // based on enum value of splitMode
            "percentageOwed": 10 // based on enum value of splitMode
        }
        "U2345": {
            "paidFull": true,
            "amountOwed": 22,
            "percentageOwed": 80
        }
    ],
}

### GET /expenses/{expenseId}
Request Headers:
Authorization: Bearer <JWT>
Request Body:
{
    "expenseId": "EX1234",
    "amount": 100.34,
    "balances": [
        {
            "from": "U2345",
            "to": "U1234",
            "amount": 22
        },
        {
            "from": "U3456",
            "to": "U1234",
            "amount": 13
        },
        {
            "from": "U4567",
            "to": "U1234",
            "amount": 10.34
        }
    ],
    "createdBy": "U1234",
    "createdAt": 121291212,
    "updatedAt": 121291212
}

POST /expenses/{expenseId}/pay
{
   "from": "U3456",
   "to": "U1234",
   "amount": 13
}
```