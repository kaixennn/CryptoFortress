# CryptoFortress Python AI Layer

The Python AI Layer provides advanced security intelligence capabilities for the CryptoFortress encryption suite.

## Components

### 1. Threat Detection Engine (Port 5000)
- Anomaly detection using Isolation Forests and Autoencoders
- Behavioral analysis of encryption patterns
- Predictive key strength analysis with ML models
- Automated vulnerability assessment

### 2. Security Recommender (Port 5001)
- Algorithm selection based on data sensitivity
- Compliance requirement mapping
- Performance vs. security tradeoff analysis
- Automated security policy generation

### 3. Risk Assessment Module (Port 5002)
- Real-time risk scoring using ensemble methods
- Threat intelligence feed integration
- Adaptive security policies based on risk levels
- Automated incident response recommendations

## Prerequisites

- Python 3.9+
- Docker and Docker Compose (for containerized deployment)

## Installation

1. Install Python dependencies:
```bash
pip install -r requirements.txt
```

2. Or use Docker (recommended):
```bash
docker-compose up --build
```

## API Endpoints

### Threat Detection Engine
- `GET /health` - Health check
- `POST /detect/anomaly` - Detect anomalies in encryption patterns
- `POST /behavioral/analyze` - Analyze behavioral patterns
- `POST /predict/key-strength` - Predict key strength
- `POST /vulnerability/assess` - Automated vulnerability assessment

### Security Recommender
- `GET /health` - Health check
- `POST /recommend/algorithm` - Recommend encryption algorithms
- `POST /compliance/map` - Map compliance requirements
- `POST /tradeoff/analyze` - Analyze performance vs. security tradeoffs
- `POST /policy/generate` - Generate automated security policies

### Risk Assessment Module
- `GET /health` - Health check
- `POST /score/realtime` - Calculate real-time risk scores
- `POST /threat/integrate` - Integrate threat intelligence feeds
- `POST /policy/adaptive` - Generate adaptive security policies
- `POST /incident/recommend` - Recommend incident response actions

## Running the Services

### Option 1: Using Docker Compose (Recommended)
```bash
# Navigate to the python-ai directory
cd python-ai

# Build and start all services
docker-compose up --build

# To run in detached mode
docker-compose up --build -d

# To stop all services
docker-compose down
```

### Option 2: Running Services Individually
```bash
# Navigate to each service directory and run
cd threat_detection
python src/main.py

cd security_recommender
python src/main.py

cd risk_assessment
python src/main.py
```

## Testing the Services

Run the test script to verify all services are working:
```bash
python test-ai-services.py
```

## Service Ports

- Threat Detection Engine: http://localhost:5000
- Security Recommender: http://localhost:5001
- Risk Assessment Module: http://localhost:5002

## Development

### Adding New Dependencies
```bash
# Add new dependency
pip install some-new-package

# Update requirements.txt
pip freeze > requirements.txt
```

### Testing
```bash
# Run tests (when implemented)
python -m pytest tests/
```