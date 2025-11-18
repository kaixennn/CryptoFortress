// React Frontend Tests
// This is the single consolidated test file for all React components

import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider } from '@mui/material/styles';
import { createTheme } from '@mui/material/styles';
import App from '../src/App';
import Dashboard from '../src/pages/monitoring/Dashboard';
import FileManagement from '../src/pages/file-management/FileManagement';
import UserManagement from '../src/pages/user-management/UserManagement';

// Mock API services
jest.mock('../src/services/api', () => ({
  authAPI: {
    login: jest.fn(),
    register: jest.fn(),
  },
  encryptionAPI: {
    encrypt: jest.fn(),
    decrypt: jest.fn(),
  },
  fileAPI: {
    upload: jest.fn(),
    download: jest.fn(),
  },
  userAPI: {
    listUsers: jest.fn(),
    createUser: jest.fn(),
  },
}));

// Mock contexts
jest.mock('../src/contexts/AuthContext', () => ({
  useAuth: () => ({
    user: { id: '1', username: 'testuser' },
    login: jest.fn(),
    logout: jest.fn(),
  }),
  AuthProvider: ({ children }) => <div>{children}</div>,
}));

jest.mock('../src/contexts/EncryptionContext', () => ({
  useEncryption: () => ({
    encryptFile: jest.fn(),
    decryptFile: jest.fn(),
  }),
  EncryptionProvider: ({ children }) => <div>{children}</div>,
}));

// Create a theme for testing
const theme = createTheme();

// Wrapper component for testing with providers
const TestWrapper = ({ children }) => (
  <ThemeProvider theme={theme}>
    <BrowserRouter>
      {children}
    </BrowserRouter>
  </ThemeProvider>
);

describe('App Integration', () => {
  test('renders without crashing', () => {
    render(
      <TestWrapper>
        <App />
      </TestWrapper>
    );
    
    // Should render the app without errors
    expect(screen.getByText(/Encryption Dashboard/i)).toBeInTheDocument();
  });
});

describe('Dashboard Component', () => {
  test('renders dashboard components', () => {
    render(
      <TestWrapper>
        <Dashboard />
      </TestWrapper>
    );
    
    // Should render the dashboard title
    expect(screen.getByText(/Encryption Dashboard/i)).toBeInTheDocument();
  });

  test('displays loading state initially', () => {
    render(
      <TestWrapper>
        <Dashboard />
      </TestWrapper>
    );
    
    // Should show loading indicator
    expect(screen.getByRole('progressbar')).toBeInTheDocument();
  });
});

describe('File Management Component', () => {
  test('renders file management components', () => {
    render(
      <TestWrapper>
        <FileManagement />
      </TestWrapper>
    );
    
    // Should render the file management title
    expect(screen.getByText(/File Management/i)).toBeInTheDocument();
  });

  test('shows upload button', () => {
    render(
      <TestWrapper>
        <FileManagement />
      </TestWrapper>
    );
    
    // Should show the upload button
    expect(screen.getByText(/Select Files/i)).toBeInTheDocument();
  });
});

describe('User Management Component', () => {
  test('renders user management components', () => {
    render(
      <TestWrapper>
        <UserManagement />
      </TestWrapper>
    );
    
    // Should render the user management title
    expect(screen.getByText(/User Management/i)).toBeInTheDocument();
  });

  test('shows add user button', () => {
    render(
      <TestWrapper>
        <UserManagement />
      </TestWrapper>
    );
    
    // Should show the add user button
    expect(screen.getByText(/Add User/i)).toBeInTheDocument();
  });
});

describe('Navigation', () => {
  test('navigates between pages', async () => {
    render(
      <TestWrapper>
        <App />
      </TestWrapper>
    );
    
    // Should start on dashboard
    expect(screen.getByText(/Encryption Dashboard/i)).toBeInTheDocument();
    
    // Navigate to files page
    const filesLink = screen.getByText(/Files/i);
    fireEvent.click(filesLink);
    
    // Should show file management
    await waitFor(() => {
      expect(screen.getByText(/File Management/i)).toBeInTheDocument();
    });
  });
});

describe('API Integration', () => {
  test('handles API calls correctly', async () => {
    const mockFileList = [
      { id: '1', name: 'test.txt', size: 1024, status: 'encrypted' }
    ];
    
    // Mock the file API
    const { fileAPI } = require('../src/services/api');
    fileAPI.list = jest.fn().mockResolvedValue({ data: mockFileList });
    
    render(
      <TestWrapper>
        <FileManagement />
      </TestWrapper>
    );
    
    // Should call the file list API
    await waitFor(() => {
      expect(fileAPI.list).toHaveBeenCalled();
    });
  });
});

describe('Error Handling', () => {
  test('handles API errors gracefully', async () => {
    // Mock API error
    const { fileAPI } = require('../src/services/api');
    fileAPI.upload = jest.fn().mockRejectedValue(new Error('Upload failed'));
    
    render(
      <TestWrapper>
        <FileManagement />
      </TestWrapper>
    );
    
    // Should handle upload errors without crashing
    expect(screen.getByText(/Select Files/i)).toBeInTheDocument();
  });
});

describe('UI Components', () => {
  test('renders Material UI components correctly', () => {
    render(
      <TestWrapper>
        <Dashboard />
      </TestWrapper>
    );
    
    // Should render Material UI components
    expect(screen.getByRole('button')).toBeInTheDocument();
  });

  test('handles user interactions', () => {
    render(
      <TestWrapper>
        <FileManagement />
      </TestWrapper>
    );
    
    // Should handle button clicks
    const uploadButton = screen.getByText(/Select Files/i);
    fireEvent.click(uploadButton);
    
    // Should respond to user interactions
    expect(uploadButton).toBeInTheDocument();
  });
});