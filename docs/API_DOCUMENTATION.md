# Auth Service API Documentation

All endpoints are prefixed with `/v1`.

## 1. Login
Authenticates a user and returns a session-bound token pair.

- **URL**: `/v1/auth/login`
- **Method**: `POST`
- **Auth required**: NO
- **Request Body**:
  ```json
  {
    "username": "your_username",
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

## 2. Logout
Revokes the current session. Any token (AT or RT) belonging to this session will no longer be valid.

- **URL**: `/v1/auth/logout`
- **Method**: `POST`
- **Auth required**: YES (Bearer Access Token)
- **Request Body**: Empty or `{}`
- **Response**:
  ```json
  {
    "message": "Logout success"
  }
  ```

---

## 3. Refresh Token
Exchanges an old Refresh Token for a new pair. Implements **JTI Rotation** (old RT is blacklisted).

- **URL**: `/v1/auth/refresh-token`
- **Method**: `POST`
- **Auth required**: NO (Tokens are passed in body)
- **Request Body**:
  ```json
  {
    "refresh_token": "eyJhbGci..."
  }
  ```
- **Response**:
  ```json
  {
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci..."
  }
  ```

---

## 4. Get Info (Example Protected Route)
Demonstrates session-based authentication enforcement.

- **URL**: `/v1/auth/get_info`
- **Method**: `GET`
- **Auth required**: YES (Bearer Access Token)
- **Response**:
  ```json
  {
    "claims": {
      "user_id": "...",
      "session_id": "...",
      "type": "access",
      "exp": "..."
  }
  ```

---

## 5. Get User Profile
Fetches the user profile for the specified user ID.

### Testing on Postman:
1. First, **Login** or **Register** to get a valid `access_token` and `user_id`.
2. Create a new request in Postman with the following details:
   - **Method**: `POST`
   - **URL**: `http://localhost:8080/v1/user/profile` (adjust port if necessary)
3. Go to the **Headers** tab and add:
   - `Authorization`: `Bearer <your_access_token>`
4. Go to the **Body** tab, select **raw** and format as **JSON**, then provide:
  ```json
  {
    "access_token": "<your_access_token>",
    "userID": "<your_user_id>"
  }
  ```
5. Send the request to receive the populated user profile.

- **Response Details**:
  ```json
  {
    "user_id": "...",
    "profile_name": "...",
    "mobile": "...",
    "gender": 1,
    "birthday": "..."
  }
  ```
