# Go-Vue Base

A modern, reusable full-stack template with Go + Gin + GORM backend and Vue.js frontend, featuring Auth0 authentication, contact form system, and admin panel.

## 🏗️ Project Structure

```
go-vue-base/
├── backend/                         # Go backend
│   ├── main.go                      # Application entry point
│   ├── internal/                    # Private application code
│   │   ├── config/                  # Configuration (DB, CORS)
│   │   ├── controllers/             # HTTP handlers (Auth0, Contact, Health)
│   │   ├── middleware/              # Auth0 JWT middleware
│   │   ├── models/                  # Data models (Contact requests)
│   │   ├── routes/                  # Route definitions
│   │   └── services/                # Business logic (Email service)
│   └── migrations/                  # Database migrations
└── frontend/                        # Vue.js frontend
    ├── src/
    │   ├── components/              # Reusable components
    │   ├── pages/                   # Page components
    │   ├── layouts/                 # Layout components
    │   └── router/                  # Vue Router configuration
    └── package.json
```

## 🚀 Quick Start

### First Time Setup (New Developers)
```bash
# Clone the repo and run setup once:
make setup
```

This will:
- ✅ Create `.env` file (you'll need to configure Auth0 settings)
- ✅ Start database and MailHog
- ✅ Run migrations
- ✅ Set up development environment

### Daily Development
```bash
# Start your development session:
make up

# When done for the day:
make stop
```

### Access Your Application
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080  
- **Health Check**: http://localhost:8080/api/health
- **MailHog UI**: http://localhost:8026 (for testing emails)

## ⚙️ Environment Configuration

Create a `.env` file in the root directory with these variables:

```bash
# Database Configuration
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_vue_base
DB_PORT=5432
DB_SSLMODE=disable

# Auth0 Configuration (Required)
AUTH0_DOMAIN=your-auth0-domain.auth0.com
AUTH0_CLIENT_ID=your-auth0-client-id
AUTH0_CLIENT_SECRET=your-auth0-client-secret
JWT_SECRET=your-jwt-secret-key
SITE_ID=go-vue-base

# Email Configuration
SMTP_HOST=mailhog
SMTP_PORT=1025
SMTP_FROM=noreply@go-vue-base.com
NOTIFICATION_EMAILS=admin@example.com

# For production, use SendGrid instead:
# SENDGRID_API_KEY=your-sendgrid-api-key

# Site Configuration
SITE_URL=http://localhost:3000
```

### Auth0 Setup
1. Create an Auth0 application
2. Enable "Resource Owner Password" grant type
3. Configure user app_metadata with `role` and `sites` fields
4. Set up Management API access for user lookups

## 🔧 Development

### Hot Reload
Both frontend and backend support hot reload:
- **Backend**: Uses Air for Go hot reload (watches `internal/` directory)
- **Frontend**: Uses Vite dev server with hot reload
- **Database**: PostgreSQL with persistent data

### Making Changes
- Edit Go files in `internal/` - backend will restart automatically
- Edit Vue files in `frontend/src/` - frontend will hot reload
- Database changes persist in Docker volume
- Use migrations for database schema changes

### Database Migrations
```bash
# Create a new migration
make db-create-migration name=create_users_table

# Run pending migrations
make db-migrate

# Rollback last migration
make db-rollback
```

### Services
- **postgres**: PostgreSQL database on port 5432
- **backend**: Go API server on port 8080
- **frontend**: Vue.js app on port 3000
- **mailhog**: Email testing server (SMTP: 1026, Web UI: 8026)

## 📚 API Architecture

### Clean Architecture Layers

1. **Controllers** (`internal/controllers/`)
   - Handle HTTP requests/responses
   - Input validation and serialization
   - Delegate business logic to services

2. **Services** (`internal/services/`)
   - Business logic and validation
   - Orchestrate data operations
   - Independent of HTTP layer

3. **Models** (`internal/models/`)
   - Data structures and database schema
   - GORM model definitions

4. **Config** (`internal/config/`)
   - Database connections
   - CORS configuration
   - Application setup

### API Endpoints

**Public Routes:**
- `GET /api/health` - Health check with database status
- `POST /api/auth/login` - Auth0 login
- `GET /api/auth/logout` - Auth0 logout
- `POST /api/contact-requests` - Submit contact form

**Protected Routes (require Auth0 JWT):**
- `GET /api/profile` - Get current user profile
- `GET /api/contact-requests` - List contact requests (admin)
- `GET /api/contact-requests/:id` - Get contact request details
- `DELETE /api/contact-requests/:id` - Delete contact request

## 🛠️ Tech Stack

- **Backend**: Go 1.22, Gin, GORM, PostgreSQL, JWT, Auth0
- **Frontend**: Vue.js 3, Tailwind CSS, Vite
- **Infrastructure**: Docker, Docker Compose, MailHog
- **Development**: Air (Go hot reload), Vite HMR
- **Authentication**: Auth0 with role-based access control
- **Email**: SendGrid (production) / MailHog (development)

## 🎯 Key Features

- **🔐 Auth0 Integration**: Modern authentication with role-based access control
- **📧 Contact Form System**: SendGrid email notifications with database storage
- **👨‍💼 Admin Panel**: Clean, responsive interface with ANHELM branding
- **🏢 Multi-tenant Support**: Site-specific access control via environment variables
- **🔄 Hot Reload**: Fast development cycle with automatic restarts
- **📱 Responsive Design**: Mobile-first approach for all interfaces
- **🐳 Docker Ready**: Complete containerized development environment
- **📊 Clean Architecture**: Separation of concerns with testable business logic
