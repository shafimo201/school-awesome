import { useEffect, useState } from 'react'
import { Routes, Route, Navigate } from 'react-router-dom'
import {
  Container,
  Typography,
  Paper,
  CircularProgress,
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Button,
} from '@mui/material'
import LoginPage from './LoginPage'
import AdminPage from './AdminPage'

function App() {
  const [token, setToken] = useState(localStorage.getItem('access_token'))
  const [profile, setProfile] = useState(null)
  const [users, setUsers] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    if (!token) {
      setLoading(false)
      setProfile(null)
      return
    }

    setLoading(true)
    fetch('/api/v1/me', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => {
        if (!res.ok) {
          throw new Error('Failed to fetch profile')
        }
        return res.json()
      })
      .then((data) => {
        setProfile(data)
        setError(null)
      })
      .catch((err) => {
        setToken(null)
        localStorage.removeItem('access_token')
        setError(err.message)
      })
      .finally(() => setLoading(false))
  }, [token])

  useEffect(() => {
    if (!token || !profile) {
      return
    }

    setLoading(true)
    fetch('/api/v1/users', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => {
        if (!res.ok) {
          throw new Error('Failed to fetch users')
        }
        return res.json()
      })
      .then((data) => {
        setUsers(data.data || [])
        setError(null)
      })
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false))
  }, [token, profile])

  const handleLogin = (newToken) => {
    localStorage.setItem('access_token', newToken)
    setToken(newToken)
  }

  const handleLogout = () => {
    localStorage.removeItem('access_token')
    setToken(null)
    setUsers([])
    setError(null)
  }

  const dashboard = (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Paper sx={{ p: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
          <Typography variant="h4">School Awesome Dashboard</Typography>
          <Button variant="outlined" onClick={handleLogout}>
            Logout
          </Button>
        </Box>

        {profile?.role_id === 'admin' && (
          <Box sx={{ mb: 3 }}>
            <Button variant="contained" href="/admin">Admin</Button>
          </Box>
        )}

        {loading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', py: 6 }}>
            <CircularProgress />
          </Box>
        ) : error ? (
          <Typography color="error">{error}</Typography>
        ) : (
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell>Email</TableCell>
                <TableCell>Full Name</TableCell>
                <TableCell>Role</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {users.map((user) => (
                <TableRow key={user.id}>
                  <TableCell>{user.id}</TableCell>
                  <TableCell>{user.email}</TableCell>
                  <TableCell>{user.full_name}</TableCell>
                  <TableCell>{user.role_id}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        )}
      </Paper>
    </Container>
  )

  return (
    <Routes>
      <Route path="/" element={token ? dashboard : <Navigate to="/login" replace />} />
      <Route path="/login" element={token ? <Navigate to="/" replace /> : <LoginPage onLogin={handleLogin} />} />
      <Route
        path="/admin"
        element={token && profile?.role_id === 'admin' ? <AdminPage token={token} /> : <Navigate to="/" replace />}
      />
      <Route path="*" element={<Navigate to={token ? '/' : '/login'} replace />} />
    </Routes>
  )
}

export default App
