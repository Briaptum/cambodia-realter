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

### Production Deployment

1. **Setup production environment:**
   ```bash
   cp env.production.example .env
   # Edit .env with your production values (database, Auth0, SendGrid, etc.)
   ```

2. **Create external network (first time only):**
   ```bash
   docker network create anhelm-network
   ```

3. **Deploy with production compose:**
   ```bash
   docker-compose -f docker-compose.prod.yml up --build -d
   ```

4. **Access your application:**
   - Frontend: http://localhost:4002 (or your configured port)
   - Backend API: http://localhost:8005 (or your configured port)

## 🚀 Auto-Deployment Setup

### **GitHub Actions Deployment**

The repository includes auto-deployment to production on every push to `master` branch.

#### **Server Setup:**
1. **Clone repository on server:**
   ```bash
   cd /var/www/
   git clone https://github.com/yourusername/go-vue-base.git go-vue-base.com
   cd go-vue-base.com
   ```

2. **Setup production environment:**
   ```bash
   cp env.production.example .env
   # Edit .env with your production values
   ```

3. **Create external Docker network:**
   ```bash
   docker network create anhelm-network
   ```

#### **GitHub Secrets Configuration:**
In your GitHub repository settings → Secrets and variables → Actions, add:

- **`DEPLOY_HOST`** - Your server IP address (e.g., `123.456.789.0`)
- **`DEPLOY_USER`** - SSH username (usually `root`)
- **`DEPLOY_SSH_KEY`** - Your private SSH key content

#### **SSH Key Setup:**
```bash
# On your local machine, generate SSH key if needed
ssh-keygen -t rsa -b 4096 -C "deploy@go-vue-base"

# Copy public key to server
ssh-copy-id root@your-server-ip

# Add private key content to GitHub Secrets as DEPLOY_SSH_KEY
cat ~/.ssh/id_rsa  # Copy this entire content
```

#### **Deployment Process:**
1. **Push to master** → Triggers automatic deployment
2. **Manual deploy** → Use GitHub Actions tab → "Deploy to Production" → "Run workflow"

#### **What happens during deployment:**
- ✅ **Git pull** latest code from master
- ✅ **Database migrations** run automatically
- ✅ **Docker rebuild** with latest code
- ✅ **Container restart** with zero-downtime
- ✅ **Health checks** ensure successful deployment

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

## 🔧 Customization Checklist

When copying this template to create a new site, update these files and settings:

### **1. Environment Configuration**
- [ ] **`.env`** - Copy from `env.example` or `env.production.example`
- [ ] **Database settings** - Update `DATABASE_URL`, `DB_NAME`, etc.
- [ ] **Auth0 configuration** - Set your `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID`, `AUTH0_CLIENT_SECRET`
- [ ] **Site identification** - Change `SITE_ID` to your project name
- [ ] **Email settings** - Configure SendGrid `SENDGRID_API_KEY`, `NOTIFICATION_EMAILS`
- [ ] **Port configuration** - Set unique `BACKEND_PORT` and `FRONTEND_PORT` for production

### **2. Branding & Content**
- [ ] **`frontend/index.html`** - Update `<title>`, meta descriptions, and Open Graph tags
- [ ] **`frontend/public/manifest.json`** - Change app `name`, `short_name`, and `description`
- [ ] **`frontend/public/favicon.png`** - Replace with your logo/favicon (128x128 recommended)
- [ ] **Site content** - Update homepage, about page, and contact page content
- [ ] **Company name** - Replace "ANHELM" branding throughout the codebase

### **3. SEO & Robots**
- [ ] **`frontend/public/robots.txt`** - Update sitemap URL to your domain
- [ ] **`frontend/public/sitemap.xml`** - Replace `yourdomain.com` with your actual domain
- [ ] **Sitemap dates** - Update `<lastmod>` dates to current date
- [ ] **Add pages** - Include any additional pages in sitemap

### **4. Database & Migrations**
- [ ] **Database name** - Update database references if needed
- [ ] **Run migrations** - Execute `make db-migrate` on new database
- [ ] **Seed data** - Add any initial data your site needs

### **5. Docker & Deployment**
- [ ] **Port conflicts** - Ensure `BACKEND_PORT` and `FRONTEND_PORT` don't conflict with other sites
- [ ] **Network name** - Update `anhelm-network` in `docker-compose.prod.yml` if needed
- [ ] **Domain configuration** - Update any domain-specific settings

### **6. Auth0 Setup**
- [ ] **Create Auth0 application** - Set up new application in Auth0 dashboard
- [ ] **Configure callbacks** - Set allowed callback URLs for your domain
- [ ] **User roles** - Set up admin/customer roles and permissions
- [ ] **App metadata** - Configure user metadata structure

### **7. Email Configuration**
- [ ] **SendGrid setup** - Create SendGrid account and get API key
- [ ] **Email templates** - Customize email content in `backend/internal/services/email_service.go`
- [ ] **From address** - Set appropriate `SMTP_FROM` email address
- [ ] **Notification recipients** - Update `NOTIFICATION_EMAILS` list

### **8. Git & Version Control**
- [ ] **Remote repository** - Update git remote to your new repository
- [ ] **README** - Update this README with project-specific information
- [ ] **License** - Add appropriate license file
- [ ] **Initial commit** - Commit your customized version

### **Quick Find & Replace**
Use these search terms to find places that need customization:
- `yourdomain.com` - Replace with your actual domain
- `ANHELM` - Replace with your company/brand name
- `go-vue-base` - Replace with your project name
- `Go-Vue Base Template` - Replace with your project title
