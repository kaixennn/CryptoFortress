from flask import Flask, request, jsonify
import numpy as np
from sklearn.ensemble import RandomForestClassifier
from sklearn.preprocessing import StandardScaler
import logging

app = Flask(__name__)

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Initialize models
risk_model = RandomForestClassifier(n_estimators=100, random_state=42)
scaler = StandardScaler()

# In a real implementation, you would load trained models from disk
logger.info("Risk Assessment Module initialized")

@app.route('/health', methods=['GET'])
def health_check():
    """Health check endpoint"""
    return jsonify({"status": "Risk Assessment Module is running", "service": "risk-assessment"}), 200

@app.route('/score/realtime', methods=['POST'])
def calculate_realtime_risk():
    """Calculate real-time risk scores using ensemble methods"""
    try:
        data = request.get_json()
        
        # Extract risk factors
        risk_factors = data.get('risk_factors', {})
        
        # Convert to feature vector
        features = np.array([
            risk_factors.get('user_behavior_score', 0.5),
            risk_factors.get('system_vulnerability_score', 0.5),
            risk_factors.get('network_threat_level', 0.5),
            risk_factors.get('data_sensitivity_score', 0.5),
            risk_factors.get('access_frequency', 0.5),
            risk_factors.get('geographical_risk', 0.5),
            risk_factors.get('time_based_risk', 0.5),
            risk_factors.get('device_trust_score', 0.5)
        ]).reshape(1, -1)
        
        # In a real implementation, you would:
        # 1. Scale the features
        # 2. Use a trained model to predict risk
        # 3. Return risk assessment
        
        # Placeholder implementation with simple weighted scoring
        weights = np.array([0.15, 0.2, 0.15, 0.1, 0.1, 0.1, 0.1, 0.1])
        risk_score = np.dot(features[0], weights)[0]
        
        # Ensure risk score is between 0 and 1
        risk_score = max(0, min(1, risk_score))
        
        # Determine risk level
        if risk_score > 0.7:
            risk_level = "HIGH"
            action_required = True
        elif risk_score > 0.4:
            risk_level = "MEDIUM"
            action_required = True
        else:
            risk_level = "LOW"
            action_required = False
        
        return jsonify({
            "risk_score": float(risk_score),
            "risk_level": risk_level,
            "action_required": action_required,
            "factors": {
                "user_behavior": risk_factors.get('user_behavior_score', 0.5),
                "system_vulnerability": risk_factors.get('system_vulnerability_score', 0.5),
                "network_threat": risk_factors.get('network_threat_level', 0.5),
                "data_sensitivity": risk_factors.get('data_sensitivity_score', 0.5)
            },
            "recommendations": generate_risk_recommendations(risk_level, risk_factors)
        }), 200
        
    except Exception as e:
        logger.error(f"Error in real-time risk calculation: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

def generate_risk_recommendations(risk_level, risk_factors):
    """Generate recommendations based on risk level and factors"""
    recommendations = []
    
    if risk_level == "HIGH":
        recommendations.append({
            "priority": "immediate",
            "action": "Restrict user access pending investigation",
            "justification": "High risk score detected"
        })
        recommendations.append({
            "priority": "high",
            "action": "Initiate security incident response procedure",
            "justification": "Potential security breach detected"
        })
    elif risk_level == "MEDIUM":
        recommendations.append({
            "priority": "medium",
            "action": "Increase monitoring frequency for this user/system",
            "justification": "Moderate risk level detected"
        })
        recommendations.append({
            "priority": "medium",
            "action": "Review recent access patterns",
            "justification": "Unusual activity detected"
        })
    else:  # LOW
        recommendations.append({
            "priority": "low",
            "action": "Continue standard monitoring",
            "justification": "Risk level within acceptable parameters"
        })
    
    # Additional recommendations based on specific factors
    if risk_factors.get('network_threat_level', 0.5) > 0.7:
        recommendations.append({
            "priority": "high",
            "action": "Review network security controls",
            "justification": "High network threat level detected"
        })
    
    if risk_factors.get('system_vulnerability_score', 0.5) > 0.7:
        recommendations.append({
            "priority": "high",
            "action": "Patch vulnerable systems immediately",
            "justification": "Critical system vulnerabilities detected"
        })
    
    return recommendations

@app.route('/threat/integrate', methods=['POST'])
def integrate_threat_intelligence():
    """Integrate threat intelligence feeds"""
    try:
        data = request.get_json()
        
        # Extract threat intelligence data
        threat_feeds = data.get('threat_feeds', [])
        
        # Process and integrate threat intelligence
        integrated_threats = []
        
        for feed in threat_feeds:
            feed_name = feed.get('name', 'Unknown')
            threats = feed.get('threats', [])
            
            for threat in threats:
                integrated_threats.append({
                    "feed_source": feed_name,
                    "threat_id": threat.get('id'),
                    "threat_type": threat.get('type'),
                    "severity": threat.get('severity', 'medium'),
                    "description": threat.get('description'),
                    "indicators": threat.get('indicators', []),
                    "first_seen": threat.get('first_seen'),
                    "last_seen": threat.get('last_seen'),
                    "confidence": threat.get('confidence', 0.5)
                })
        
        # In a real implementation, you would:
        # 1. Correlate threats with internal data
        # 2. Update threat models
        # 3. Generate alerts for matching indicators
        
        return jsonify({
            "integrated_threats": integrated_threats,
            "total_threats": len(integrated_threats),
            "high_confidence_threats": sum(1 for t in integrated_threats if t["confidence"] > 0.8),
            "critical_threats": sum(1 for t in integrated_threats if t["severity"] == "critical")
        }), 200
        
    except Exception as e:
        logger.error(f"Error in threat intelligence integration: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/policy/adaptive', methods=['POST'])
def generate_adaptive_policies():
    """Generate adaptive security policies based on risk levels"""
    try:
        data = request.get_json()
        
        # Extract current risk assessment
        risk_assessment = data.get('risk_assessment', {})
        current_risk_level = risk_assessment.get('risk_level', 'LOW')
        
        # Generate adaptive policies
        policies = {}
        
        if current_risk_level == "HIGH":
            policies["access_control"] = {
                "policy": "restrictive",
                "description": "Enhanced access controls due to high risk",
                "rules": [
                    "Mandatory multi-factor authentication",
                    "Just-in-time access approval",
                    "Session timeout reduced to 15 minutes",
                    "Geographical access restrictions"
                ]
            }
            policies["monitoring"] = {
                "policy": "intensive",
                "description": "Increased monitoring due to high risk",
                "rules": [
                    "Real-time alerting for all activities",
                    "Full packet capture for network traffic",
                    "Enhanced logging for all systems",
                    "24/7 security operations center"
                ]
            }
        elif current_risk_level == "MEDIUM":
            policies["access_control"] = {
                "policy": "standard_plus",
                "description": "Enhanced standard controls due to moderate risk",
                "rules": [
                    "Multi-factor authentication required",
                    "Regular access reviews",
                    "Session timeout of 1 hour",
                    "Basic geographical restrictions"
                ]
            }
            policies["monitoring"] = {
                "policy": "enhanced",
                "description": "Enhanced monitoring due to moderate risk",
                "rules": [
                    "Alerting for suspicious activities",
                    "Sampling of network traffic",
                    "Standard logging for all systems",
                    "Business hours security monitoring"
                ]
            }
        else:  # LOW
            policies["access_control"] = {
                "policy": "standard",
                "description": "Standard access controls for normal risk",
                "rules": [
                    "Single-factor authentication for internal users",
                    "Quarterly access reviews",
                    "Session timeout of 8 hours",
                    "No geographical restrictions"
                ]
            }
            policies["monitoring"] = {
                "policy": "standard",
                "description": "Standard monitoring for normal risk",
                "rules": [
                    "Alerting for high-risk activities",
                    "Periodic log reviews",
                    "Standard logging for all systems",
                    "Business hours security monitoring"
                ]
            }
        
        return jsonify({
            "adaptive_policies": policies,
            "current_risk_level": current_risk_level,
            "policy_update_time": "2025-01-01T00:00:00Z"
        }), 200
        
    except Exception as e:
        logger.error(f"Error in adaptive policy generation: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/incident/recommend', methods=['POST'])
def recommend_incident_response():
    """Recommend incident response actions"""
    try:
        data = request.get_json()
        
        # Extract incident details
        incident = data.get('incident', {})
        incident_type = incident.get('type', 'unknown')
        severity = incident.get('severity', 'low')
        affected_systems = incident.get('affected_systems', [])
        
        # Generate response recommendations
        recommendations = []
        
        # General recommendations
        recommendations.append({
            "step": 1,
            "action": "Contain the incident",
            "priority": "immediate",
            "description": "Prevent further damage and limit the scope of the incident"
        })
        
        recommendations.append({
            "step": 2,
            "action": "Eradicate the threat",
            "priority": "high",
            "description": "Remove the cause of the incident and restore systems to a secure state"
        })
        
        recommendations.append({
            "step": 3,
            "action": "Recover and restore",
            "priority": "medium",
            "description": "Restore systems and services while ensuring security"
        })
        
        recommendations.append({
            "step": 4,
            "action": "Post-incident review",
            "priority": "low",
            "description": "Analyze the incident and improve security measures"
        })
        
        # Incident-specific recommendations
        if incident_type == "data_breach":
            recommendations.insert(1, {
                "step": 1.1,
                "action": "Notify affected parties",
                "priority": "immediate",
                "description": "Begin notification process as required by law"
            })
            recommendations.insert(1, {
                "step": 1.2,
                "action": "Preserve evidence",
                "priority": "immediate",
                "description": "Ensure forensic evidence is not lost"
            })
        elif incident_type == "malware":
            recommendations.insert(1, {
                "step": 1.1,
                "action": "Isolate infected systems",
                "priority": "immediate",
                "description": "Disconnect affected systems from the network"
            })
        
        # Severity-based escalation
        if severity == "critical":
            recommendations.append({
                "step": 5,
                "action": "Executive notification",
                "priority": "immediate",
                "description": "Notify senior management and board of directors"
            })
        elif severity == "high":
            recommendations.append({
                "step": 5,
                "action": "Management notification",
                "priority": "high",
                "description": "Notify relevant management personnel"
            })
        
        return jsonify({
            "incident_response_plan": recommendations,
            "incident_type": incident_type,
            "severity": severity,
            "estimated_resolution_time": estimate_resolution_time(severity, incident_type)
        }), 200
        
    except Exception as e:
        logger.error(f"Error in incident response recommendation: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

def estimate_resolution_time(severity, incident_type):
    """Estimate incident resolution time"""
    # Simple estimation logic
    base_hours = {
        "low": 4,
        "medium": 24,
        "high": 72,
        "critical": 168  # 1 week
    }
    
    multiplier = {
        "data_breach": 1.5,
        "malware": 1.2,
        "intrusion": 2.0,
        "denial_of_service": 1.1,
        "unknown": 1.0
    }
    
    base = base_hours.get(severity, 24)
    mult = multiplier.get(incident_type, 1.0)
    
    return f"{int(base * mult)} hours"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5002, debug=True)