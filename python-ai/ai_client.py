"""
Python AI Layer Integration for CryptoFortress
This module provides the interface between the Python AI services and the Go backend.
"""

import requests
import json
from typing import Dict, List, Any

class AILayerClient:
    """Client for interacting with Python AI services from Go backend"""
    
    def __init__(self, base_url: str = "http://localhost:5000"):
        self.base_url = base_url
        self.session = requests.Session()
        
    def detect_encryption_patterns(self, data: Dict[str, Any]) -> Dict[str, Any]:
        """
        Send data to threat detection service for pattern analysis
        
        Args:
            data: Dictionary containing encryption metadata
            
        Returns:
            Dictionary with threat analysis results
        """
        try:
            response = self.session.post(
                f"{self.base_url}/detect/anomaly",
                json=data,
                timeout=30
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            return {
                "error": f"Failed to detect patterns: {str(e)}",
                "status": "error"
            }
    
    def analyze_key_strength(self, key_data: Dict[str, Any]) -> Dict[str, Any]:
        """
        Analyze key strength using entropy analysis
        
        Args:
            key_data: Dictionary containing key metadata
            
        Returns:
            Dictionary with key strength analysis
        """
        try:
            response = self.session.post(
                f"{self.base_url}/analyze/key-strength",
                json=key_data,
                timeout=30
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            return {
                "error": f"Failed to analyze key strength: {str(e)}",
                "status": "error"
            }
    
    def get_security_recommendations(self, requirements: Dict[str, Any]) -> Dict[str, Any]:
        """
        Get security recommendations based on requirements
        
        Args:
            requirements: Dictionary containing security requirements
            
        Returns:
            Dictionary with security recommendations
        """
        try:
            response = self.session.post(
                f"{self.base_url}/recommend/algorithm",
                json=requirements,
                timeout=30
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            return {
                "error": f"Failed to get recommendations: {str(e)}",
                "status": "error"
            }
    
    def generate_security_report(self, metrics: Dict[str, Any]) -> Dict[str, Any]:
        """
        Generate automated security report
        
        Args:
            metrics: Dictionary containing security metrics
            
        Returns:
            Dictionary with security report
        """
        try:
            response = self.session.post(
                f"{self.base_url}/report/generate",
                json=metrics,
                timeout=30
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            return {
                "error": f"Failed to generate report: {str(e)}",
                "status": "error"
            }

# Example usage
if __name__ == "__main__":
    # Example of how Go backend would use this client
    ai_client = AILayerClient("http://localhost:5000")
    
    # Example threat detection request
    encryption_data = {
        "algorithm": "AES-256-GCM",
        "key_length": 256,
        "data_size": 1024,
        "timestamp": "2023-01-01T00:00:00Z",
        "user_id": "user123"
    }
    
    result = ai_client.detect_encryption_patterns(encryption_data)
    print("Threat Detection Result:", result)
    
    # Example key strength analysis
    key_data = {
        "key_hex": "0123456789abcdef" * 8,  # 256-bit key in hex
        "algorithm": "AES-256",
        "usage_count": 1000
    }
    
    result = ai_client.analyze_key_strength(key_data)
    print("Key Strength Analysis:", result)