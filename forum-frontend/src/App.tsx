import React, { useState, useEffect } from 'react'
import { BrowserRouter, Routes, Route, Navigate, useNavigate } from 'react-router-dom'

import LoginPage from './pages/LoginPage'
import RegisterPage from './pages/RegisterPage'
import TopicsPage from './pages/TopicsPage'
import TopicPage from './pages/TopicPage'
import PostPage from './pages/PostPage'

const Header: React.FC<{ onLogout: () => void }> = ({ onLogout }) => (
  <header style={{ padding: '10px 20px', borderBottom: '1px solid #eee', display: 'flex', justifyContent: 'flex-end' }}>
    <button
      onClick={onLogout}
      style={{
        padding: '6px 16px',
        background: '#ff4444',
        color: 'white',
        border: 'none',
        borderRadius: '4px',
        cursor: 'pointer',
      }}
    >
      Logout
    </button>
  </header>
)

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const token = localStorage.getItem('token')
  if (!token) return <Navigate to="/login" replace />
  return <>{children}</>
}

const PublicRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const token = localStorage.getItem('token')
  if (token) return <Navigate to="/topics" replace />
  return <>{children}</>
}

const AppContent: React.FC = () => {
  const navigate = useNavigate()
  const [token, setToken] = useState<string | null>(localStorage.getItem('token'))

  useEffect(() => {
    const checkToken = () => setToken(localStorage.getItem('token'))
    window.addEventListener('storage', checkToken)
    return () => window.removeEventListener('storage', checkToken)
  }, [])

  const handleLoginSuccess = (token: string) => {
    localStorage.setItem('token', token)
    setToken(token)
    navigate('/topics')
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    setToken(null)
    navigate('/login')
  }

  return (
    <div className="app">
      <Routes>
        <Route
          path="/login"
          element={
            <PublicRoute>
              <LoginPage onLoginSuccess={handleLoginSuccess} onSwitchToRegister={() => navigate('/register')} />
            </PublicRoute>
          }
        />
        <Route
          path="/register"
          element={
            <PublicRoute>
              <RegisterPage onSwitchToLogin={() => navigate('/login')} />
            </PublicRoute>
          }
        />

        <Route
          path="/topics"
          element={
            <ProtectedRoute>
              <Header onLogout={handleLogout} />
              <TopicsPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/topics/:id"
          element={
            <ProtectedRoute>
              <Header onLogout={handleLogout} />
              <TopicPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/posts/:postId"
          element={
            <ProtectedRoute>
              <Header onLogout={handleLogout} />
              <PostPage />
            </ProtectedRoute>
          }
        />

        <Route
          path="/"
          element={token ? <Navigate to="/topics" replace /> : <Navigate to="/login" replace />}
        />
        <Route
          path="*"
          element={token ? <Navigate to="/topics" replace /> : <Navigate to="/login" replace />}
        />
      </Routes>
    </div>
  )
}

const App: React.FC = () => (
  <BrowserRouter>
    <AppContent />
  </BrowserRouter>
)

export default App