from flask import Flask, request, jsonify
import numpy as np
from sklearn.ensemble import IsolationForest
from sklearn.preprocessing import StandardScaler
import joblib
import logging

app = Flask(__name__)

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Initialize models
isolation_forest = IsolationForest(contamination=0.1, random_state=42)
scaler = StandardScaler()

# In a real implementation, you would load trained models from disk
# For now, we'll initialize them without training data
logger.info("Threat Detection Engine initialized")

@app.route('/health', methods=['GET'])
def health_check():
    """Health check endpoint"""
    return jsonify({"status": "Threat Detection Engine is running", "service": "threat-detection"}), 200

@app.route('/detect/anomaly', methods=['POST'])
def detect_anomaly():
    """Detect anomalies in encryption patterns"""
    try:
        data = request.get_json()
        
        # Extract features from request
        features = np.array(data.get('features', []))
        
        if len(features) == 0:
            return jsonify({"error": "No features provided"}), 400
        
        # Reshape if needed
        if len(features.shape) == 1:
            features = features.reshape(1, -1)
        
        # Scale features
        scaled_features = scaler.fit_transform(features)
        
        # Detect anomalies
        anomaly_scores = isolation_forest.decision_function(scaled_features)
        predictions = isolation_forest.predict(scaled_features)
        
        # Convert predictions to readable format
        # -1 indicates anomaly, 1 indicates normal
        results = []
        for i, pred in enumerate(predictions):
            results.append({
                "index": i,
                "is_anomaly": bool(pred == -1),
                "anomaly_score": float(anomaly_scores[i])
            })
        
        return jsonify({
            "results": results,
            "total_samples": len(results),
            "anomalies_detected": sum(1 for r in results if r["is_anomaly"])
        }), 200
        
    except Exception as e:
        logger.error(f"Error in anomaly detection: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/behavioral/analyze', methods=['POST'])
def analyze_behavioral_patterns():
    """Analyze behavioral patterns in encryption usage"""
    try:
        data = request.get_json()
        
        # Extract behavioral data
        user_patterns = data.get('user_patterns', [])
        
        if not user_patterns:
            return jsonify({"error": "No user patterns provided"}), 400
        
        # In a real implementation, you would:
        # 1. Compare current patterns with historical baselines
        # 2. Detect deviations from normal behavior
        # 3. Calculate risk scores
        
        # Placeholder implementation
        risk_scores = []
        for pattern in user_patterns:
            # Simple risk calculation based on frequency and data volume
            frequency = pattern.get('frequency', 0)
            data_volume = pattern.get('data_volume', 0)
            
            # Calculate a simple risk score
            risk_score = min(1.0, (frequency * 0.01) + (data_volume * 0.0001))
            
            risk_scores.append({
                "user_id": pattern.get('user_id', 'unknown'),
                "risk_score": risk_score,
                "is_suspicious": risk_score > 0.7
            })
        
        return jsonify({
            "analysis": risk_scores,
            "total_users": len(risk_scores),
            "suspicious_activities": sum(1 for r in risk_scores if r["is_suspicious"])
        }), 200
        
    except Exception as e:
        logger.error(f"Error in behavioral analysis: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/predict/key-strength', methods=['POST'])
def predict_key_strength():
    """Predict key strength using ML models"""
    try:
        data = request.get_json()
        
        # Extract key parameters
        key_params = data.get('key_parameters', {})
        
        # In a real implementation, you would:
        # 1. Extract features from key parameters
        # 2. Use a trained model to predict strength
        # 3. Return strength assessment
        
        # Placeholder implementation
        key_length = key_params.get('length', 0)
        algorithm = key_params.get('algorithm', 'unknown')
        entropy = key_params.get('entropy', 0)
        
        # Simple strength calculation
        if key_length >= 256 and algorithm in ['AES-256-GCM', 'ChaCha20-Poly1305']:
            strength = 0.95
            rating = "Strong"
        elif key_length >= 128:
            strength = 0.75
            rating = "Good"
        else:
            strength = 0.3
            rating = "Weak"
        
        return jsonify({
            "strength_score": strength,
            "rating": rating,
            "recommendations": [
                "Use keys of at least 256 bits for symmetric encryption",
                "Consider quantum-resistant algorithms for long-term security"
            ] if strength < 0.8 else [
                "Key strength is adequate for current security requirements"
            ]
        }), 200
        
    except Exception as e:
        logger.error(f"Error in key strength prediction: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/vulnerability/assess', methods=['POST'])
def assess_vulnerabilities():
    """Automated vulnerability assessment"""
    try:
        data = request.get_json()
        
        # Extract system information
        system_info = data.get('system_info', {})
        
        # In a real implementation, you would:
        # 1. Check for known vulnerabilities in crypto libraries
        # 2. Assess configuration security
        # 3. Identify potential attack vectors
        
        # Placeholder implementation
        vulnerabilities = []
        
        # Check key rotation policy
        rotation_period = system_info.get('key_rotation_days', 365)
        if rotation_period > 180:
            vulnerabilities.append({
                "id": "KEY_ROTATION_001",
                "severity": "medium",
                "description": "Key rotation period is longer than recommended",
                "recommendation": "Rotate keys every 90 days or less"
            })
        
        # Check for weak algorithms
        algorithms = system_info.get('algorithms', [])
        weak_algorithms = ['DES', '3DES', 'RC4']
        for algo in algorithms:
            if algo in weak_algorithms:
                vulnerabilities.append({
                    "id": "CRYPTO_ALGO_001",
                    "severity": "high",
                    "description": f"Weak cryptographic algorithm detected: {algo}",
                    "recommendation": f"Replace {algo} with a stronger algorithm"
                })
        
        # Check for quantum-resistant preparation
        quantum_ready = system_info.get('quantum_resistant', False)
        if not quantum_ready:
            vulnerabilities.append({
                "id": "QUANTUM_001",
                "severity": "medium",
                "description": "System not prepared for quantum computing threats",
                "recommendation": "Implement quantum-resistant cryptography"
            })
        
        return jsonify({
            "vulnerabilities": vulnerabilities,
            "total_vulnerabilities": len(vulnerabilities),
            "critical": sum(1 for v in vulnerabilities if v["severity"] == "high"),
            "high": sum(1 for v in vulnerabilities if v["severity"] == "medium")
        }), 200
        
    except Exception as e:
        logger.error(f"Error in vulnerability assessment: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)