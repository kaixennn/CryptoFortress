#!/usr/bin/env python3

import os

def check_directory_structure():
    """Check if the AI layer directory structure is correct"""
    print("Verifying Python AI Layer Structure")
    print("=" * 40)
    
    # Check main directories
    main_dirs = ["threat_detection", "security_recommender", "risk_assessment", "shared_libs"]
    for dir_name in main_dirs:
        if os.path.exists(dir_name):
            print(f"✓ {dir_name} directory exists")
        else:
            print(f"✗ {dir_name} directory missing")
    
    print()
    
    # Check service directories
    services = ["threat_detection", "security_recommender", "risk_assessment"]
    for service in services:
        print(f"Checking {service}:")
        
        # Check src directory
        src_dir = f"{service}/src"
        if os.path.exists(src_dir):
            print(f"  ✓ src directory exists")
        else:
            print(f"  ✗ src directory missing")
        
        # Check main.py
        main_file = f"{service}/src/main.py"
        if os.path.exists(main_file):
            print(f"  ✓ main.py exists")
        else:
            print(f"  ✗ main.py missing")
        
        # Check requirements.txt
        req_file = f"{service}/requirements.txt"
        if os.path.exists(req_file):
            print(f"  ✓ requirements.txt exists")
        else:
            print(f"  ✗ requirements.txt missing")
        
        # Check Dockerfile
        docker_file = f"{service}/Dockerfile"
        if os.path.exists(docker_file):
            print(f"  ✓ Dockerfile exists")
        else:
            print(f"  ✗ Dockerfile missing")
        
        print()

def check_required_files():
    """Check if required files exist"""
    print("Checking Required Files:")
    print("-" * 20)
    
    required_files = [
        "requirements.txt",
        "docker-compose.yml",
        "README.md",
        "test-ai-services.py",
        "verify-ai-structure.py"
    ]
    
    for file_name in required_files:
        if os.path.exists(file_name):
            print(f"✓ {file_name}")
        else:
            print(f"✗ {file_name}")

if __name__ == "__main__":
    check_directory_structure()
    print()
    check_required_files()
    print("\nVerification complete.")