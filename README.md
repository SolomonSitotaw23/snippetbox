# Snippetbox

Snippetbox is a simple web application for storing and sharing code snippets.

## Features

- User authentication (signup, login, logout)
- Create, view, and list code snippets
- CSRF protection for forms
- Flash messages for user feedback
- Secure HTTP headers
- Password hashing with bcrypt
- Session management with MySQL backend
- Responsive UI with custom CSS
- Live reload during development (using Air)

## Built With

- [Go](https://golang.org/) (1.24+)
- [MySQL](https://www.mysql.com/) (database)
- [SCS](https://github.com/alexedwards/scs) (session management)
- [httprouter](https://github.com/julienschmidt/httprouter) (routing)
- [nosurf](https://github.com/justinas/nosurf) (CSRF protection)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) (password hashing)
- [Air](https://github.com/cosmtrek/air) (live reload for development)

## How to Run

1. **Clone the repository:**

   ```sh
   git clone https://github.com/solomonsitotaw23/snippetbox.git
   cd snippetbox
   ```

2. **Set up MySQL database:**

   - Create a database named `snippetbox`.
   - Create the required tables for `users` and `snippets`.

3. **Configure environment:**

   - Edit the DSN in [`cmd/web/main.go`](cmd/web/main.go) or pass it as a flag:
     ```
     -dsn="user:password@/snippetbox?parseTime=true"
     ```

4. **Run in development mode (with live reload):**

   ```sh
   make dev
   ```

   _(Requires [Air](https://github.com/cosmtrek/air) installed)_

5. **Build and run manually:**

   ```sh
   make build
   ./snippetbox
   ```

   Or:

   ```sh
   make run
   ```

6. **Access the app:**
   - Open [https://localhost:4000](https://localhost:4000) in your browser.

## Project Structure

- `cmd/web/` — main web application code
- `internal/models/` — database models
- `internal/validator/` — form validation
- `ui/html/` — HTML templates
- `ui/static/` — static assets (CSS, JS, images)
- `tls/` — TLS certificates for HTTPS

##
