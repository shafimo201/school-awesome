import { useState } from 'react'
import { Container, Paper, Box, Typography, TextField, Button, Alert } from '@mui/material'

const validateInput = ({ username, full_name, password }) => {
  const errors = {}
  if (!username || username.trim().length < 2) {
    errors.username = 'Enter a username or student ID.'
  }
  if (!full_name || full_name.length < 3) {
    errors.full_name = 'Full name must be at least 3 characters.'
  }
  if (!password || password.length < 8) {
    errors.password = 'Password must be at least 8 characters.'
  }
  return errors
}

export default function AdminPage({ token }) {
  const [student, setStudent] = useState({ username: '', full_name: '', password: '' })
  const [teacher, setTeacher] = useState({ username: '', full_name: '', password: '' })
  const [studentErrors, setStudentErrors] = useState({})
  const [teacherErrors, setTeacherErrors] = useState({})
  const [msg, setMsg] = useState(null)

  const submit = async (path, payload, setter, setErrors) => {
    setMsg(null)
    const errors = validateInput(payload)
    setErrors(errors)
    if (Object.keys(errors).length > 0) {
      setMsg({ type: 'error', text: 'Please fix the validation errors before submitting.' })
      return
    }

    try {
      const res = await fetch(path, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(payload),
      })
      const data = await res.json()
      if (!res.ok) throw new Error(data.message || 'request failed')
      setMsg({ type: 'success', text: `Created: ${data.username || data.email}` })
      setter({ username: '', full_name: '', password: '' })
      setErrors({})
    } catch (e) {
      setMsg({ type: 'error', text: e.message })
    }
  }

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Paper sx={{ p: 4 }}>
        <Typography variant="h5" sx={{ mb: 2 }}>Admin — Manage Users</Typography>
        {msg && <Alert severity={msg.type} sx={{ mb: 2 }}>{msg.text}</Alert>}

        <Box sx={{ mb: 4 }}>
          <Typography variant="h6">Create Student</Typography>
          <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2, mt: 2 }}>
            <TextField
              label="Username / Student ID"
              value={student.username}
              error={Boolean(studentErrors.username)}
              helperText={studentErrors.username}
              onChange={(e) => setStudent({ ...student, username: e.target.value })}
            />
            <TextField
              label="Full name"
              value={student.full_name}
              error={Boolean(studentErrors.full_name)}
              helperText={studentErrors.full_name}
              onChange={(e) => setStudent({ ...student, full_name: e.target.value })}
            />
            <TextField
              label="Password"
              type="password"
              value={student.password}
              error={Boolean(studentErrors.password)}
              helperText={studentErrors.password}
              onChange={(e) => setStudent({ ...student, password: e.target.value })}
            />
            <Button variant="contained" onClick={() => submit('/api/v1/admin/students', student, setStudent, setStudentErrors)}>
              Create
            </Button>
          </Box>
        </Box>

        <Box>
          <Typography variant="h6">Create Teacher</Typography>
          <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2, mt: 2 }}>
            <TextField
              label="Username / Student ID"
              value={teacher.username}
              error={Boolean(teacherErrors.username)}
              helperText={teacherErrors.username}
              onChange={(e) => setTeacher({ ...teacher, username: e.target.value })}
            />
            <TextField
              label="Full name"
              value={teacher.full_name}
              error={Boolean(teacherErrors.full_name)}
              helperText={teacherErrors.full_name}
              onChange={(e) => setTeacher({ ...teacher, full_name: e.target.value })}
            />
            <TextField
              label="Password"
              type="password"
              value={teacher.password}
              error={Boolean(teacherErrors.password)}
              helperText={teacherErrors.password}
              onChange={(e) => setTeacher({ ...teacher, password: e.target.value })}
            />
            <Button variant="contained" onClick={() => submit('/api/v1/admin/teachers', teacher, setTeacher, setTeacherErrors)}>
              Create
            </Button>
          </Box>
        </Box>
      </Paper>
    </Container>
  )
}
