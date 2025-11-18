// React Frontend Integration for CryptoFortress
// This module provides the interface between the React frontend and backend services

import axios from 'axios';

// API client configuration
const apiClient = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL || 'http://localhost:8080',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('authToken');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor to handle errors
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Handle unauthorized access
      localStorage.removeItem('authToken');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Authentication API
export const authAPI = {
  login: (credentials) => apiClient.post('/auth/login', credentials),
  register: (userData) => apiClient.post('/auth/register', userData),
  refreshToken: (refreshToken) => apiClient.post('/auth/refresh', { refreshToken }),
  logout: () => apiClient.post('/auth/logout'),
};

// Encryption API
export const encryptionAPI = {
  encrypt: (data) => apiClient.post('/encryption/encrypt', data),
  decrypt: (data) => apiClient.post('/encryption/decrypt', data),
  generateKey: (params) => apiClient.post('/encryption/generate-key', params),
};

// Key Management API
export const keyManagementAPI = {
  listKeys: () => apiClient.get('/keys'),
  getKey: (keyId) => apiClient.get(`/keys/${keyId}`),
  createKey: (keyData) => apiClient.post('/keys', keyData),
  updateKey: (keyId, keyData) => apiClient.put(`/keys/${keyId}`, keyData),
  deleteKey: (keyId) => apiClient.delete(`/keys/${keyId}`),
  rotateKey: (keyId) => apiClient.post(`/keys/${keyId}/rotate`),
};

// Audit API
export const auditAPI = {
  getLogs: (params) => apiClient.get('/audit/logs', { params }),
  getLog: (logId) => apiClient.get(`/audit/logs/${logId}`),
  generateReport: (reportParams) => apiClient.post('/audit/reports', reportParams),
};

// File Management API
export const fileAPI = {
  upload: (file, metadata) => {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('metadata', JSON.stringify(metadata));
    return apiClient.post('/files/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
  },
  download: (fileId) => apiClient.get(`/files/download/${fileId}`, {
    responseType: 'blob',
  }),
  list: (params) => apiClient.get('/files', { params }),
  delete: (fileId) => apiClient.delete(`/files/${fileId}`),
  encrypt: (fileId, encryptionParams) => apiClient.post(`/files/${fileId}/encrypt`, encryptionParams),
  decrypt: (fileId, decryptionParams) => apiClient.post(`/files/${fileId}/decrypt`, decryptionParams),
};

// AI Layer Integration
export const aiAPI = {
  detectThreats: (data) => apiClient.post('/ai/threat-detection', data),
  analyzeKeyStrength: (keyData) => apiClient.post('/ai/key-analysis', keyData),
  getRecommendations: (requirements) => apiClient.post('/ai/recommendations', requirements),
  getRiskScore: (metrics) => apiClient.post('/ai/risk-assessment', metrics),
};

// User Management API
export const userAPI = {
  listUsers: () => apiClient.get('/users'),
  getUser: (userId) => apiClient.get(`/users/${userId}`),
  createUser: (userData) => apiClient.post('/users', userData),
  updateUser: (userId, userData) => apiClient.put(`/users/${userId}`, userData),
  deleteUser: (userId) => apiClient.delete(`/users/${userId}`),
};

// Export the base client for custom requests
export default apiClient;