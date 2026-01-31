# Mess & Meal Experience – Meal Intent System

## Overview

A residential mess serves meals daily to a large group of residents. However, students’ plans frequently change — they may skip meals or decide to eat at the last minute — while mess staff must decide food quantities in advance.

This mismatch often leads to food waste on some days and shortages on others, as decisions are based on guesswork or late information.

This project implements a **Meal Intent Declaration System** that captures residents’ meal confirmations ahead of time, enabling mess staff to plan food quantities more accurately while still allowing flexibility for residents.

---

## Problem Chosen

The core problem addressed is the lack of reliable, per-meal attendance information before food preparation begins.

Instead of solving the entire mess management problem, this system focuses on a **single, high-impact slice**:
- Allowing residents to explicitly confirm meals
- Aggregating confirmations into a clear planning signal for staff

---

## Solution Summary

- Users can sign up and log in
- Mess staff can create meals for a given date
- Residents can view available meals
- Residents can confirm or update their meal status
- Each confirmation is tied to a specific user and meal
- Mess staff can view a summary of confirmed meals to decide quantities

Authentication is handled using JWT, and protected routes require a valid token.

---

## Key Features

- User authentication (signup & login)
- Per-user meal confirmation
- Idempotent meal confirmation updates
- Staff-facing meal summary endpoint
- SQLite-backed persistence
- Minimal, REST-based backend design

---

## Assumptions

To keep the system focused and realistic, the following assumptions were made:

- Users are pre-registered via signup (no admin user management)
- Authentication is token-based and simplified
- No role-based access enforcement (staff vs student is assumed)
- No UI; this is an API-only backend
- No payment, billing, or penalty logic
- No cutoff-time enforcement (can be added as an extension)

---

## Tech Stack

- Language: Go
- Framework: Gin
- Database: SQLite (modernc.org/sqlite)
- Authentication: JWT
- Password hashing: bcrypt

---

## Project Structure
mess-meal-system/
├── main.go
├── db/
│ └── db.go
├── middlewares/
│ └── auth.go
├── models/
│ ├── user.go
│ ├── meal.go
│ └── meal_confirmation.go
├── routes/
│ ├── routes.go
│ ├── users.go
│ ├── meals.go
│ ├── confirm.go
│ └── summary.go
├── utils/
│ ├── hash.go
│ └── jwt.go
├── go.mod
├── go.sum
└── README.md


---

## How to Run

### Prerequisites
- Go 1.21 or higher

### Steps
```bash
go mod tidy
go run main.go

The server will start on:

http://localhost:8080

API Endpoints
🔐 Authentication
Sign up
POST /signup


Request

{
  "email": "user@example.com",
  "password": "password123"
}

Login
POST /login


Response

{
  "message": "Login successful!",
  "token": "<jwt-token>"
}

👤 User (Authenticated)
Get logged-in user details
GET /mydetails


Headers

Authorization: <jwt-token>

🍽️ Meals (Public)
Get all meals
GET /meals

Get a specific meal
GET /meal/:id

✅ Meal Confirmation (Authenticated – Residents)
Confirm or update meal status
POST /meal/confirm


Headers

Authorization: <jwt-token>


Request

{
  "meal_id": 1,
  "status": "YES"
}

Get my meal confirmations
GET /meal/confirm

🧑‍🍳 Staff APIs (Authenticated)
Create a meal
POST /meals


Request

{
  "name": "Lunch",
  "date": "2026-02-01"
}

Get meal summary
GET /staff/meal/:id/summary


Response

{
  "meal_id": "1",
  "total_confirmed": 42
}

Design Rationale

The system focuses on explicit intent declaration rather than prediction
API surface is intentionally minimal to reduce complexity
Authentication exists only to associate actions with users
The summary endpoint provides a single, reliable planning metric for staff
The design can be extended to support cutoff times, roles, and forecasting
Possible Extensions
Role-based access control (student vs staff)
Cutoff time enforcement for confirmations
Guest count support
Notifications before cutoff
Historical analytics and forecasting
I deliberately kept the API surface minimal — just enough to support daily intent declaration and final headcount generation — because simplicity and reliability matter more than feature breadth in operational systems.

Author
Palash Sinha
