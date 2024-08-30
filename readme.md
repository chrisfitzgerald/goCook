# goCook

## 🍳 Overview

goCook is a modern web application template built with Go, leveraging the power of Echo framework and Google OAuth for authentication. It provides a solid foundation for building secure and scalable web applications.

## 🚀 Features

- **Google OAuth Integration**: Seamless login experience using Google accounts
- **JWT Authentication**: Secure session management with JSON Web Tokens
- **Responsive UI**: Clean and intuitive user interface
- **Customizable**: Easy to extend and adapt to your specific needs

## 🛠 Tech Stack

- **Backend**: Go with Echo framework
- **Authentication**: Google OAuth 2.0 and JWT
- **Frontend**: HTML, CSS
- **Environment**: dotenv for configuration management

## 🚦 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/goCook.git
   cd goCook
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up your environment variables:
   Create a `.env` file in the root directory and add the following:
   ```bash
   GOOGLE_CLIENT_ID=your_google_client_id
   GOOGLE_CLIENT_SECRET=your_google_client_secret
   JWT_SECRET_KEY=your_jwt_secret_key
   ```

4. Run the application:
   ```bash
   go run cmd/app/*.go
   ```

5. Open your browser and navigate to `http://localhost:8080`

## 🔒 Authentication Flow

1. User clicks "Login with Google"
2. User is redirected to Google for authentication
3. Upon successful authentication, a JWT is created and stored as a cookie
4. The JWT is used for subsequent authenticated requests

## 🎨 Customization

- Modify the HTML templates in the `template` directory to change the UI
- Adjust the CSS in `template/style.css` to match your desired look and feel
- Extend the backend functionality by adding new routes and handlers in `cmd/app/main.go`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License.