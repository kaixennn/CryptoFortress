#!/bin/bash

# Test script to verify all CryptoFortress services are running

echo "Testing CryptoFortress Services"

# Define service URLs and ports
SERVICES=(
  "Authentication Service:http://localhost:8080/health"
  "Encryption Service:http://localhost:8081/health"
  "Key Management Service:http://localhost:8082/health"
  "Audit & Compliance Service:http://localhost:8083/health"
)

# Function to test a service
test_service() {
  local service_name=$1
  local service_url=$2
  
  echo "Testing $service_name..."
  
  # Use curl to test the service with a 5-second timeout
  response=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 "$service_url")
  
  if [ "$response" = "200" ]; then
    echo "  ✓ $service_name is running"
  else
    echo "  ✗ $service_name is not responding (HTTP $response)"
  fi
}

# Test each service
for service in "${SERVICES[@]}"; do
  # Split the service name and URL
  IFS=':' read -r service_name service_url <<< "$service"
  test_service "$service_name" "$service_url"
done

echo ""
echo "Database connectivity test:"
# Test database connectivity
if command -v pg_isready &> /dev/null; then
  if pg_isready -h localhost -p 5432 -U user -d cryptofortress &> /dev/null; then
    echo "  ✓ PostgreSQL database is accessible"
  else
    echo "  ✗ PostgreSQL database is not accessible"
  fi
else
  echo "  ? pg_isready not available, skipping database test"
fi

echo ""
echo "Test completed."