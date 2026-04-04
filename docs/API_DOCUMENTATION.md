# API Documentation

All endpoints are prefixed with `/v1`.

---

## Auth

### 1. Login
- **POST** `/v1/auth/login`
- **Auth**: No
- **Rate Limited**: Yes
- **Request**:
  ```json
  {
    "email": "your_email",
    "password": "your_password"
  }
  ```
- **Response**:
  ```json
  {
    "user_id": "0x019CD...",
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci...",
    "message": "login success"
  }
  ```

---

### 2. Logout
- **POST** `/v1/auth/logout`
- **Auth**: Yes (Bearer Access Token)
- **Request**: `{}`
- **Response**:
  ```json
  { "message": "Logout success" }
  ```

---

### 3. Refresh Token
Implements JTI rotation — old refresh token is blacklisted on use.

- **POST** `/v1/auth/refresh-token`
- **Auth**: No
- **Request**:
  ```json
  { "refresh_token": "eyJhbGci..." }
  ```
- **Response**:
  ```json
  {
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci..."
  }
  ```

---

## User

### 4. Register
- **POST** `/v1/user/register`
- **Auth**: No
- **Rate Limited**: Yes
- **Request**:
  ```json
  {
    "username": "your_username",
    "password": "your_password",
    "email": "your_email"
  }
  ```
- **Response**:
  ```json
  { "message": "Registration successful" }
  ```

---

### 5. Get User Profile
- **POST** `/v1/user/profile`
- **Auth**: Yes (Bearer Access Token)
- **Request**:
  ```json
  { "userID": "0x019CD..." }
  ```
- **Response**:
  ```json
  {
    "user_id": "0x019CD...",
    "profile_name": "John Doe",
    "mobile": "08123456789",
    "gender": 1,
    "birthday": "1990-06-15"
  }
  ```

---

### 6. Update User Profile
- **POST** `/v1/user/profile/update`
- **Auth**: Yes (Bearer Access Token)
- **Request**:
  ```json
  {
    "user_id": "0x019CD...",
    "data": {
      "profile_name": "John Doe",
      "mobile": "08123456789",
      "gender": 1,
      "birthday": "1990-06-15"
    }
  }
  ```
  > All fields inside `data` are optional. Only provided fields will be updated.
  > `birthday` must be in `YYYY-MM-DD` format.
- **Response**:
  ```json
  { "message": "Update successful" }
  ```

---

## Product

### 7. Get Product
- **GET** `/v1/order/get-product`
- **Auth**: No
- **Status**: 🚧 Work in progress

---

## Health

### 8. Health Check
- **GET** `/v1/health`
- **Auth**: No
- **Response**:
  ```json
  { "status": "ok" }
  ```
