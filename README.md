## CVWO Assignment - Forum Web Application

Name: Paing Htoo

# Project Description

This project is a simple forum-style web application where users can create topics, write posts, and comment on posts.
The application is built with a React + TypeScript frontend and a Go backend API.

Users can:

- Create discussion topics
- Add posts under topics
- Comment on posts
- Edit or delete their content

# Tech stack

Frontend

- React
- TypeScript
- Vite
- Axios

Backend

- Go (Golang)
- REST API

Database

- MySQL

  
# Database Setup (MySQL)

1. Create a MySQL database

CREATE DATABASE forum;

# Project Structure

CVWO_Assignment has forum-frontend for react and type script application and forum-backend for Go REST API server
and README.md file to understand the project structure and tech stack and instructions to set up .

# Setup Instructions

Follow the steps below to run the project locally.

1. Clone the Repository

git clone https://github.com/feb-dylan/CVWO_Assignment.git

cd CVWO_Assignment

# Backend Setup (Go)

Step 1: Navigate to backend folder

cd forum-backend

Step 2: Install Go dependencies

go mod tidy

step 3: create .env file in forum-backend folder and set DB_DSN , PORT and JWT_SECRET .

Step 4: Run the backend server

go run main.go

The backend server should start on:

http://localhost:8080

# Frontend Setup (React + TypeScript)

Step 1: Open a new terminal

Step 2: Navigate to frontend folder

cd forum-frontend

Step 3: Install dependencies

npm install

Step 4: Start the development server

npm run dev

The frontend should start on:

http://localhost:5173

# Features Implemented

Topics

- Create topic
- View topics
- Edit topic
- Delete topic

Posts

- Create post
- View post
- Edit post
- Delete post

Comments

- Add comment
- View comments
- Edit comment
- Delete comment
- Reply features

  
Authentication

- username + password authentication using Jwt


# AI Usage Declaration

AI Tool Used:
ChatGPT (OpenAI) , deepseek

How AI Was Used:

- To understand Go syntax and backend structure
- To help debug programming errors
- To assist in writing documentation (README)
- To clarify concepts related to React and API integration

AI was used only as a learning assistant.
All code implementation and final decisions were written and verified by the author.

# Future Improvements

- Add likes or reactions to posts
- Improve UI design
- Deploy the application online
