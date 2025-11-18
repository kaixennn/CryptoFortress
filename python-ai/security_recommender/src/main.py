from flask import Flask, request, jsonify
import logging

app = Flask(__name__)

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

logger.info("Security Recommender Engine initialized")

@app.route('/health', methods=['GET'])
def health_check():
    """Health check endpoint"""
    return jsonify({"status": "Security Recommender Engine is running", "service": "security-recommender"}), 200

@app.route('/recommend/algorithm', methods=['POST'])
def recommend_algorithm():
    """Recommend encryption algorithms based on data sensitivity"""
    try:
        data = request.get_json()
        
        # Extract data sensitivity information
        sensitivity = data.get('sensitivity', 'medium')
        data_type = data.get('data_type', 'general')
        compliance_requirements = data.get('compliance', [])
        performance_constraints = data.get('performance_constraints', {})
        
        # Algorithm recommendations based on sensitivity
        recommendations = []
        
        if sensitivity == 'high':
            if 'HIPAA' in compliance_requirements or 'GDPR' in compliance_requirements:
                recommendations.append({
                    "algorithm": "AES-256-GCM",
                    "reason": "Meets HIPAA/GDPR requirements for sensitive data",
                    "strength": "Very Strong"
                })
                recommendations.append({
                    "algorithm": "Kyber-1024",
                    "reason": "Quantum-resistant algorithm for long-term security",
                    "strength": "Future-Proof"
                })
            else:
                recommendations.append({
                    "algorithm": "ChaCha20-Poly1305",
                    "reason": "High performance with strong security",
                    "strength": "Strong"
                })
        elif sensitivity == 'medium':
            recommendations.append({
                "algorithm": "AES-192-GCM",
                "reason": "Balanced security and performance",
                "strength": "Strong"
            })
        else:  # low sensitivity
            recommendations.append({
                "algorithm": "AES-128-GCM",
                "reason": "Adequate security with optimal performance",
                "strength": "Good"
            })
        
        # Add quantum-resistant recommendation if not already included
        if not any('Kyber' in rec['algorithm'] for rec in recommendations):
            recommendations.append({
                "algorithm": "Kyber-768",
                "reason": "Quantum-resistant algorithm for forward security",
                "strength": "Future-Proof"
            })
        
        return jsonify({
            "recommendations": recommendations,
            "data_sensitivity": sensitivity,
            "data_type": data_type
        }), 200
        
    except Exception as e:
        logger.error(f"Error in algorithm recommendation: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/compliance/map', methods=['POST'])
def map_compliance_requirements():
    """Map compliance requirements to security controls"""
    try:
        data = request.get_json()
        
        # Extract compliance frameworks
        frameworks = data.get('frameworks', [])
        
        # Map requirements to security controls
        compliance_map = {}
        
        for framework in frameworks:
            if framework == 'GDPR':
                compliance_map[framework] = {
                    "data_encryption": {
                        "requirement": "Article 32 - Encryption of personal data",
                        "controls": [
                            "AES-256 encryption for data at rest",
                            "TLS 1.3 for data in transit",
                            "Key rotation every 90 days"
                        ]
                    },
                    "data_breach_notification": {
                        "requirement": "Article 33 - Notification of personal data breach",
                        "controls": [
                            "Real-time breach detection",
                            "Automated incident response",
                            "72-hour notification procedure"
                        ]
                    },
                    "right_to_be_forgotten": {
                        "requirement": "Article 17 - Right to erasure",
                        "controls": [
                            "Secure data deletion",
                            "Cryptographic erasure techniques",
                            "Audit trail for deletions"
                        ]
                    }
                }
            elif framework == 'HIPAA':
                compliance_map[framework] = {
                    "data_encryption": {
                        "requirement": "164.312(a)(2)(iv) - Encryption",
                        "controls": [
                            "AES-256 encryption for ePHI",
                            "End-to-end encryption for transmission",
                            "HSM-based key management"
                        ]
                    },
                    "access_control": {
                        "requirement": "164.312(a)(1) - Access control",
                        "controls": [
                            "Role-based access control",
                            "Multi-factor authentication",
                            "Audit logs for all access"
                        ]
                    }
                }
            elif framework == 'SOC2':
                compliance_map[framework] = {
                    "security": {
                        "requirement": "Common Criteria - Security",
                        "controls": [
                            "Comprehensive access controls",
                            "Intrusion detection systems",
                            "Regular security assessments"
                        ]
                    },
                    "availability": {
                        "requirement": "Common Criteria - Availability",
                        "controls": [
                            "High availability architecture",
                            "Disaster recovery procedures",
                            "Performance monitoring"
                        ]
                    }
                }
        
        return jsonify({
            "compliance_mapping": compliance_map,
            "frameworks_covered": list(compliance_map.keys())
        }), 200
        
    except Exception as e:
        logger.error(f"Error in compliance mapping: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/tradeoff/analyze', methods=['POST'])
def analyze_tradeoffs():
    """Analyze performance vs. security tradeoffs"""
    try:
        data = request.get_json()
        
        # Extract requirements
        requirements = data.get('requirements', {})
        target_performance = requirements.get('performance_target', 'balanced')
        security_priority = requirements.get('security_priority', 'high')
        
        # Analyze tradeoffs
        analysis = {
            "performance_impact": {},
            "security_benefits": {},
            "recommendations": []
        }
        
        if target_performance == 'high':
            analysis["performance_impact"] = {
                "latency": "Low (<10ms)",
                "throughput": "High (>10,000 ops/sec)",
                "resource_usage": "Optimized"
            }
            analysis["security_benefits"] = {
                "encryption": "AES-128-GCM or ChaCha20-Poly1305",
                "key_management": "Standard rotation",
                "authentication": "Single-factor"
            }
            analysis["recommendations"].append({
                "priority": "performance",
                "action": "Use lightweight algorithms for maximum throughput",
                "risk": "Reduced security margin"
            })
        elif target_performance == 'balanced':
            analysis["performance_impact"] = {
                "latency": "Moderate (10-50ms)",
                "throughput": "Good (1,000-10,000 ops/sec)",
                "resource_usage": "Balanced"
            }
            analysis["security_benefits"] = {
                "encryption": "AES-192-GCM",
                "key_management": "Regular rotation with HSM",
                "authentication": "MFA recommended"
            }
            analysis["recommendations"].append({
                "priority": "balanced",
                "action": "Use standard security practices with reasonable performance",
                "risk": "Acceptable for most applications"
            })
        else:  # low performance priority (max security)
            analysis["performance_impact"] = {
                "latency": "Higher (>50ms)",
                "throughput": "Lower (<1,000 ops/sec)",
                "resource_usage": "High"
            }
            analysis["security_benefits"] = {
                "encryption": "AES-256-GCM + Kyber-1024",
                "key_management": "HSM with Shamir's Secret Sharing",
                "authentication": "MFA + Biometrics"
            }
            analysis["recommendations"].append({
                "priority": "security",
                "action": "Implement maximum security measures",
                "risk": "Performance impact but highest protection"
            })
        
        # Adjust based on security priority
        if security_priority == 'maximum':
            analysis["recommendations"].append({
                "priority": "security",
                "action": "Consider quantum-resistant algorithms for long-term security",
                "risk": "Additional computational overhead"
            })
        
        return jsonify(analysis), 200
        
    except Exception as e:
        logger.error(f"Error in tradeoff analysis: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

@app.route('/policy/generate', methods=['POST'])
def generate_security_policy():
    """Generate automated security policies"""
    try:
        data = request.get_json()
        
        # Extract policy requirements
        organization = data.get('organization', 'Unknown')
        industry = data.get('industry', 'General')
        compliance_frameworks = data.get('compliance_frameworks', [])
        
        # Generate policy sections
        policy = {
            "organization": organization,
            "industry": industry,
            "effective_date": "2025-01-01",
            "review_date": "2026-01-01",
            "sections": {}
        }
        
        # Data encryption policy
        policy["sections"]["data_encryption"] = {
            "title": "Data Encryption Policy",
            "description": "All sensitive data must be encrypted both at rest and in transit",
            "requirements": [
                "Use AES-256-GCM for data at rest",
                "Use TLS 1.3 for data in transit",
                "Implement key rotation every 90 days",
                "Store encryption keys in HSM"
            ]
        }
        
        # Access control policy
        policy["sections"]["access_control"] = {
            "title": "Access Control Policy",
            "description": "Implement role-based access control with multi-factor authentication",
            "requirements": [
                "Enforce least privilege principle",
                "Require MFA for all users",
                "Regular access reviews",
                "Automated deprovisioning"
            ]
        }
        
        # Incident response policy
        policy["sections"]["incident_response"] = {
            "title": "Incident Response Policy",
            "description": "Establish procedures for detecting and responding to security incidents",
            "requirements": [
                "24/7 monitoring of security events",
                "72-hour breach notification for GDPR",
                "Document all security incidents",
                "Regular incident response drills"
            ]
        }
        
        # Add compliance-specific sections
        if 'GDPR' in compliance_frameworks:
            policy["sections"]["gdpr_compliance"] = {
                "title": "GDPR Compliance",
                "requirements": [
                    "Implement data protection by design",
                    "Appoint Data Protection Officer",
                    "Conduct Data Protection Impact Assessments",
                    "Ensure right to data portability"
                ]
            }
        
        if 'HIPAA' in compliance_frameworks:
            policy["sections"]["hipaa_compliance"] = {
                "title": "HIPAA Compliance",
                "requirements": [
                    "Encrypt all ePHI",
                    "Implement audit controls",
                    "Train workforce on HIPAA requirements",
                    "Business Associate Agreements"
                ]
            }
        
        return jsonify(policy), 200
        
    except Exception as e:
        logger.error(f"Error in policy generation: {str(e)}")
        return jsonify({"error": "Internal server error"}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)