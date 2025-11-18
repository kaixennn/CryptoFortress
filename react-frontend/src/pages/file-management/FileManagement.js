import React, { useState } from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Grid,
  IconButton,
  List,
  ListItem,
  ListItemText,
  TextField,
  Typography,
  Snackbar,
  Alert
} from '@mui/material';
import {
  CloudUpload as UploadIcon,
  Lock as EncryptIcon,
  LockOpen as DecryptIcon,
  Delete as DeleteIcon,
  Download as DownloadIcon
} from '@mui/icons-material';
import { useEncryption } from '../../contexts/EncryptionContext';
import { encryptFile, decryptFile, uploadFile, downloadFile } from '../../services/fileService';

const FileManagement = () => {
  const { encryptFile: encryptContext, decryptFile: decryptContext } = useEncryption();
  const [files, setFiles] = useState([]);
  const [selectedFile, setSelectedFile] = useState(null);
  const [openDialog, setOpenDialog] = useState(false);
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });

  const handleFileUpload = async (event) => {
    const file = event.target.files[0];
    if (!file) return;

    setLoading(true);
    try {
      const uploadedFile = await uploadFile(file);
      setFiles(prev => [...prev, uploadedFile]);
      showSnackbar('File uploaded successfully', 'success');
    } catch (error) {
      showSnackbar('Error uploading file: ' + error.message, 'error');
    } finally {
      setLoading(false);
    }
  };

  const handleEncrypt = async (file) => {
    setLoading(true);
    try {
      const encryptedFile = await encryptFile(file.id, password);
      setFiles(prev => prev.map(f => f.id === file.id ? encryptedFile : f));
      showSnackbar('File encrypted successfully', 'success');
    } catch (error) {
      showSnackbar('Error encrypting file: ' + error.message, 'error');
    } finally {
      setLoading(false);
      setOpenDialog(false);
    }
  };

  const handleDecrypt = async (file) => {
    setLoading(true);
    try {
      const decryptedFile = await decryptFile(file.id, password);
      setFiles(prev => prev.map(f => f.id === file.id ? decryptedFile : f));
      showSnackbar('File decrypted successfully', 'success');
    } catch (error) {
      showSnackbar('Error decrypting file: ' + error.message, 'error');
    } finally {
      setLoading(false);
      setOpenDialog(false);
    }
  };

  const handleDownload = async (file) => {
    try {
      await downloadFile(file.id, file.name);
      showSnackbar('File downloaded successfully', 'success');
    } catch (error) {
      showSnackbar('Error downloading file: ' + error.message, 'error');
    }
  };

  const handleDelete = (fileId) => {
    setFiles(prev => prev.filter(f => f.id !== fileId));
    showSnackbar('File deleted successfully', 'success');
  };

  const showSnackbar = (message, severity) => {
    setSnackbar({ open: true, message, severity });
  };

  const closeSnackbar = () => {
    setSnackbar(prev => ({ ...prev, open: false }));
  };

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Typography variant="h4" gutterBottom>
        File Management
      </Typography>
      
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Box display="flex" justifyContent="space-between" alignItems="center">
                <Typography variant="h6">Upload Files</Typography>
                <input
                  accept="*/*"
                  style={{ display: 'none' }}
                  id="file-upload"
                  type="file"
                  onChange={handleFileUpload}
                />
                <label htmlFor="file-upload">
                  <Button
                    variant="contained"
                    component="span"
                    startIcon={<UploadIcon />}
                    disabled={loading}
                  >
                    Select Files
                  </Button>
                </label>
              </Box>
              
              {loading && (
                <Box display="flex" justifyContent="center" mt={2}>
                  <CircularProgress />
                </Box>
              )}
            </CardContent>
          </Card>
        </Grid>
        
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Your Files
              </Typography>
              
              {files.length === 0 ? (
                <Typography color="textSecondary" align="center">
                  No files uploaded yet
                </Typography>
              ) : (
                <List>
                  {files.map((file) => (
                    <ListItem
                      key={file.id}
                      secondaryAction={
                        <>
                          <IconButton 
                            edge="end" 
                            aria-label="encrypt"
                            onClick={() => {
                              setSelectedFile(file);
                              setOpenDialog(true);
                            }}
                          >
                            <EncryptIcon />
                          </IconButton>
                          <IconButton 
                            edge="end" 
                            aria-label="download"
                            onClick={() => handleDownload(file)}
                          >
                            <DownloadIcon />
                          </IconButton>
                          <IconButton 
                            edge="end" 
                            aria-label="delete"
                            onClick={() => handleDelete(file.id)}
                          >
                            <DeleteIcon />
                          </IconButton>
                        </>
                      }
                    >
                      <ListItemText
                        primary={file.name}
                        secondary={`Size: ${file.size} bytes | Status: ${file.status}`}
                      />
                    </ListItem>
                  ))}
                </List>
              )}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
      
      <Dialog open={openDialog} onClose={() => setOpenDialog(false)}>
        <DialogTitle>
          {selectedFile ? `Encrypt ${selectedFile.name}` : 'Encrypt File'}
        </DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label="Encryption Password"
            type="password"
            fullWidth
            variant="standard"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenDialog(false)}>Cancel</Button>
          <Button 
            onClick={() => handleEncrypt(selectedFile)}
            variant="contained"
          >
            Encrypt
          </Button>
          <Button 
            onClick={() => handleDecrypt(selectedFile)}
            variant="outlined"
          >
            Decrypt
          </Button>
        </DialogActions>
      </Dialog>
      
      <Snackbar open={snackbar.open} autoHideDuration={6000} onClose={closeSnackbar}>
        <Alert onClose={closeSnackbar} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Box>
  );
};

export default FileManagement;