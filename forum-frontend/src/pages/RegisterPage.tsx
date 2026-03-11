import React from 'react'
import Register from '../components/auth/Register'

interface RegisterPageProps {
  onSwitchToLogin: () => void
}

const RegisterPage: React.FC<RegisterPageProps> = (props) => {
  return <Register {...props} />
}

export default RegisterPage