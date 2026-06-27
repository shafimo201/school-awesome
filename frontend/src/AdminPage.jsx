import { useState } from 'react'
import { Container, Paper, Box, Typography, TextField, Button, Alert } from '@mui/material'

export default function AdminPage({ token }) {
  const [student, setStudent] = useState({ email: '', full_name: '', password: '' })
  const [teacher, setTeacher] = useState({ email: '', full_name: '', password: '' })
  const [msg, setMsg] = useState(null)

  const submit = async (path, payload, setter) => {
    setMsg(null)
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
      setMsg({ type: 'success', text: `Created: ${data.email}` })
      setter({ email: '', full_name: '', password: '' })
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
          <Box sx={{ display: 'flex', gap: 2, mt: 2 }}>
            <TextField label="Email" value={student.email} onChange={(e) => setStudent({ ...student, email: e.target.value })} />
            <TextField label="Full name" value={student.full_name} onChange={(e) => setStudent({ ...student, full_name: e.target.value })} />
            <TextField label="Password" type="password" value={student.password} onChange={(e) => setStudent({ ...student, password: e.target.value })} />
            <Button variant="contained" onClick={() => submit('/api/v1/admin/students', student, setStudent)}>Create</Button>
          </Box>
        </Box>

        <Box>
          <Typography variant="h6">Create Teacher</Typography>
          <Box sx={{ display: 'flex', gap: 2, mt: 2 }}>
            <TextField label="Email" value={teacher.email} onChange={(e) => setTeacher({ ...teacher, email: e.target.value })} />
            <TextField label="Full name" value={teacher.full_name} onChange={(e) => setTeacher({ ...teacher, full_name: e.target.value })} />
            <TextField label="Password" type="password" value={teacher.password} onChange={(e) => setTeacher({ ...teacher, password: e.target.value })} />
            <Button variant="contained" onClick={() => submit('/api/v1/admin/teachers', teacher, setTeacher)}>Create</Button>
          </Box>
        </Box>
      </Paper>
    </Container>
  )
}
