import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Dashboard from './pages/monitoring/Dashboard';
import FileManagement from './pages/file-management/FileManagement';
import UserManagement from './pages/user-management/UserManagement';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import { AuthProvider } from './contexts/AuthContext';
import { EncryptionProvider } from './contexts/EncryptionContext';

const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2',
    },
    secondary: {
      main: '#e57373',
    },
    background: {
      default: '#f5f5f5',
    },
  },
});

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <AuthProvider>
        <EncryptionProvider>
          <Router>
            <Header />
            <Sidebar />
            <Routes>
              <Route path="/" element={<Dashboard />} />
              <Route path="/dashboard" element={<Dashboard />} />
              <Route path="/files" element={<FileManagement />} />
              <Route path="/users" element={<UserManagement />} />
            </Routes>
          </Router>
        </EncryptionProvider>
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;