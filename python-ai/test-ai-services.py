#!/usr/bin/env python3

import requests
import json
import time

# Service URLs
SERVICES = {
    "threat_detection": "http://localhost:5000",
    "security_recommender": "http://localhost:5001",
    "risk_assessment": "http://localhost:5002"
}

def test_service_health(service_name, base_url):
    """Test the health endpoint of a service"""
    try:
        response = requests.get(f"{base_url}/health", timeout=5)
        if response.status_code == 200:
            print(f"  ✓ {service_name} is running")
            return True
        else:
            print(f"  ✗ {service_name} returned status {response.status_code}")
            return False
    except requests.exceptions.RequestException as e:
        print(f"  ✗ {service_name} is not responding: {str(e)}")
        return False

def test_threat_detection():
    """Test threat detection service endpoints"""
    base_url = SERVICES["threat_detection"]
    
    # Test anomaly detection
    try:
        data = {
            "features": [[1, 2, 3, 4], [5, 6, 7, 8]]
        }
        response = requests.post(f"{base_url}/detect/anomaly", 
                               json=data, timeout=10)
        if response.status_code == 200:
            print("  ✓ Threat detection anomaly detection working")
        else:
            print(f"  ✗ Threat detection anomaly detection failed: {response.status_code}")
    except Exception as e:
        print(f"  ✗ Threat detection anomaly detection error: {str(e)}")
    
    # Test behavioral analysis
    try:
        data = {
            "user_patterns": [
                {"user_id": "user1", "frequency": 10, "data_volume": 1000},
                {"user_id": "user2", "frequency": 100, "data_volume": 50000}
            ]
        }
        response = requests.post(f"{base_url}/behavioral/analyze", 
                               json=data, timeout=10)
        if response.status_code == 200:
            print("  ✓ Threat detection behavioral analysis working")
        else:
            print(f"  ✗ Threat detection behavioral analysis failed: {response.status_code}")
    except Exception as e:
        print(f"  ✗ Threat detection behavioral analysis error: {str(e)}")

def test_security_recommender():
    """Test security recommender service endpoints"""
    base_url = SERVICES["security_recommender"]
    
    # Test algorithm recommendation
    try:
        data = {
            "sensitivity": "high",
            "data_type": "personal",
            "compliance": ["GDPR", "HIPAA"]
        }
        response = requests.post(f"{base_url}/recommend/algorithm", 
                               json=data, timeout=10)
        if response.status_code == 200:
            print("  ✓ Security recommender algorithm recommendation working")
        else:
            print(f"  ✗ Security recommender algorithm recommendation failed: {response.status_code}")
    except Exception as e:
        print(f"  ✗ Security recommender algorithm recommendation error: {str(e)}")
    
    # Test compliance mapping
    try:
        data = {
            "frameworks": ["GDPR", "HIPAA", "SOC2"]
        }
        response = requests.post(f"{base_url}/compliance/map", 
                               json=data, timeout=10)
        if response.status_code == 200:
            print("  ✓ Security recommender compliance mapping working")
        else:
            print(f"  ✗ Security recommender compliance mapping failed: {response.status_code}")
    except Exception as e:
        print(f"  ✗ Security recommender compliance mapping error: {str(e)}")

def test_risk_assessment():
    """Test risk assessment service endpoints"""
    base_url = SERVICES["risk_assessment"]
    
    # Test real-time risk scoring
    try:
        data = {
            "risk_factors": {
                "user_behavior_score": 0.8,
                "system_vulnerability_score": 0.6,
                "network_threat_level": 0.7,
                "data_sensitivity_score": 0.9
            }
        }
        response = requests.post(f"{base_url}/score/realtime", 
                               json=data, timeout=10)
        if response.status_code == 200:
            print("  ✓ Risk assessment real-time scoring working")
        else:
            print(f"  ✗ Risk assessment real-time scoring failed: {response.status_code}")
    except Exception as e:
        print(f"  ✗ Risk assessment real-time scoring error: {str(e)}")
    
    # Test threat intelligence integration
    try:
        data = {
            "threat_feeds": [
                {
                    "name": "test_feed",
                    "threats": [
                        {
                            "id": "T001",
                            "type": "malware",
                            "severity": "high",
                            "description": "Test threat",
                            "indicators": ["192.168.1.100"],
                            "confidence": 0.8
                        }
                    ]
                }
            ]
        }
        response = requests.post(f"{base_url}/threat/integrate", 
                               json=data, timeout=10)
        if response.status_code == 200:
            print("  ✓ Risk assessment threat intelligence integration working")
        else:
            print(f"  ✗ Risk assessment threat intelligence integration failed: {response.status_code}")
    except Exception as e:
        print(f"  ✗ Risk assessment threat intelligence integration error: {str(e)}")

def main():
    print("Testing CryptoFortress Python AI Services")
    print("=" * 50)
    
    # Test health endpoints
    print("\n1. Testing service health:")
    all_healthy = True
    for service_name, base_url in SERVICES.items():
        if not test_service_health(service_name, base_url):
            all_healthy = False
    
    if not all_healthy:
        print("\nSome services are not responding. Please ensure all services are running.")
        return
    
    print("\n2. Testing service functionality:")
    
    # Test each service
    print("\nThreat Detection Service:")
    test_threat_detection()
    
    print("\nSecurity Recommender Service:")
    test_security_recommender()
    
    print("\nRisk Assessment Service:")
    test_risk_assessment()
    
    print("\n" + "=" * 50)
    print("AI services testing completed.")

if __name__ == "__main__":
    main()