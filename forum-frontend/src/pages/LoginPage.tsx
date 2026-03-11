
import React from 'react'
import Login from '../components/auth/Login'

interface LoginPageProps {
  onLoginSuccess: (token: string) => void
  onSwitchToRegister: () => void
}

const LoginPage: React.FC<LoginPageProps> = (props) => {
  return <Login {...props} />
}

export default LoginPage