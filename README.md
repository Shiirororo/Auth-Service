# Auth Service (Go)
<DEPRECATED>
Authentication and user information service following a layered architecture.

## Architecture Flow
`User -> Router -> Middleware -> Service -> Repository -> Database`

## API Documentation
Detailed endpoint instructions can be found in [API_DOCUMENTATION.md](API_DOCUMENTATION.md).

## Key Features
- **Session-based Blacklist**: Logout revokes the entire session ID.
- **JWT Token Rotation**: Secure refresh token flow with JTI blocking.
- **Layered Design**: Adheres to SOLID principles for maintainability.


Update:
FROM THIS 3/10 commit foward, Authentication part will be seperated into a different repository.


wow