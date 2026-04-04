# Email Domain — Implementation Plan

## Current State

| Path | Status |
|---|---|
| `internal/email/application/service/email.service.go` | Mock — hardcoded API key, no config injection |
| `internal/email/controller/dto/email_template.html` | Static placeholder HTML |
| No `SendGrid` config section in `pkg/settings/section.go` | Missing |
| No `OTP` email template (only plain-text `"OTP: "+otp`) | Incomplete |

---

## Target Structure

```
internal/email/
├── application/
│   └── service/
│       └── email.service.go       ← refactored: config-injected, HTML template
└── controller/
    └── dto/
        └── otp_template.html      ← proper OTP HTML template
```

---

## Implementation Steps

### 1. Add `SendGrid` Config Section

**`pkg/settings/section.go`**

Add `SendGrid SendGridSetting` to `Config` and define:

```go
type SendGridSetting struct {
    APIKey    string `mapstructure:"api_key"`
    FromEmail string `mapstructure:"from_email"`
    FromName  string `mapstructure:"from_name"`
}
```

---

### 2. Refactor `EmailService`

**`internal/email/application/service/email.service.go`**

- Inject `*settings.SendGridSetting` via constructor (no hardcoded key)
- Parse `otp_template.html` with `html/template` and render OTP into HTML body
- Keep `EmailServiceInterface` — add `SendWelcome(email, username string) error` for post-register flow

```go
type EmailServiceInterface interface {
    SendOTP(email, otp string) error
    SendWelcome(email, username string) error
}
```

Constructor:

```go
func NewEmailService(cfg *settings.SendGridSetting) EmailServiceInterface
```

---

### 3. OTP HTML Template

**`internal/email/controller/dto/otp_template.html`**

Replace static placeholder with a proper Go template:

```html
<h2>Your OTP Code</h2>
<p>Code: <b>{{.OTP}}</b></p>
<p>Expires in 2 minutes.</p>
```

---

### 4. Wire `EmailService` into `RegisterWorker`

**`internal/user/application/worker/register_worker.go`**

After successful user creation (step 6), dispatch or directly call `emailService.SendWelcome(payload.Email, payload.Username)`.

> Prefer direct call inside the worker — no new event type needed.

---

### 5. Update Wire Providers

**`internal/wire/wire.go` / `wire_gen.go`**

- Add `NewEmailService` provider
- Inject `*settings.SendGridSetting` from the loaded config
- Inject `EmailServiceInterface` into `RegisterWorker`

---

## Dependency Map

```
Config (SendGridSetting)
    └── NewEmailService → EmailServiceInterface
            └── RegisterWorker (inject alongside existing repos)
```

---

## Out of Scope

- Email delivery retries / dead-letter queue (Kafka `AuditEvent` can cover this later)
- Email domain DB persistence (OTP already lives in Redis via `OTPRepository`)
- Additional templates beyond OTP + Welcome
