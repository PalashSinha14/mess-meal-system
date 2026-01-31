# Mess & Meal Experience – Meal Intent System

## Problem Chosen
Mess staff must estimate food quantities in advance, but students’ plans frequently change, leading to food waste or shortages.

This project focuses on a Meal Intent Declaration system that allows residents to declare whether they will eat a meal ahead of time, with a fixed cutoff to provide predictable numbers to the mess staff.

## Solution Overview
- Meals are created for a given date and time
- Students declare intent (eat / skip / guests)
- Declarations are allowed only before the cutoff time
- Mess staff can view final headcount per meal

## Assumptions
- Authentication is out of scope
- One intent per user per meal
- Guests are counted as additional servings
- No UI; API-only system

## Tech Stack
- Go
- net/http (or Gin)
- SQLite

## How to Run
```bash
go mod tidy
go run main.go

I deliberately kept the API surface minimal — just enough to support daily intent declaration and final headcount generation — because simplicity and reliability matter more than feature breadth in operational systems.