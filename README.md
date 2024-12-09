# E-Commerce Backend Project


## Deployment
The project is deployed and accessible at:
- **Frontend**: [https://shopping-cart-topaz-mu.vercel.app/](https://shopping-cart-topaz-mu.vercel.app/)
- **Backend**: [https://shopping-cart-production-a800.up.railway.app/](https://shopping-cart-production-a800.up.railway.app/)
  
## Overview
This project implements a simple e-commerce backend using the following tools:
- **ORM**: [GORM](https://github.com/jinzhu/gorm/)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin/)

The backend supports essential e-commerce functionalities such as user management, item listing, cart management, and order creation.

## Prerequisites
- **Go**: Installed on your system.
- **PostgreSQL**: Set up as the database.
- **Node.js**: For the React frontend.

## Setup Instructions

### Backend
1. Clone the repository:
   ```bash
   git clone https://github.com/vasanthakumar21/Shopping-cart.git
   cd backend
   ```

2. Run database migrations:
   ```bash
   go run migrate/migrate.go
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

The backend will run at `http://localhost:8080`.

### Frontend
1. Navigate to the `frontend` folder:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the React app:
   ```bash
   npm run dev
   ```

The frontend will run at `http://localhost:5173`.

## Features
1. **User Management**:
   - Sign up (`POST /users`): Create a new user account.
   - Login (`POST /users/login`): Authenticate an existing user and receive a token for further requests.
   - List Users (`GET /users`): Retrieve a list of all users.

2. **Item Management**:
   - Add Item (`POST /items`): Create a new item.
   - List Items (`GET /items`): Retrieve a list of all items.

3. **Cart Management**:
   - Add to Cart (`POST /carts`): Add items to a user's cart.
   - List Carts (`GET /carts`): Retrieve all carts.

4. **Order Management**:
   - Create Order (`POST /checkout`): Convert a cart into an order.
   - List Orders (`GET /orders`): Retrieve all orders.

