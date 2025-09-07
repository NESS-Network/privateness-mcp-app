from flask import Flask, jsonify, request
import subprocess
import json
import os
from pathlib import Path

app = Flask(__name__)

PRIVATENESS_HOME = os.path.expanduser("~/.privateness-keys")

@app.route('/api/nodes', methods=['GET'])
def get_nodes():
    try:
        result = subprocess.run(
            ['./node', 'list', 'all'],
            cwd=os.getcwd(),
            capture_output=True,
            text=True
        )
        # Parse the output and return as JSON
        nodes = parse_node_list(result.stdout)
        return jsonify(nodes)
    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/api/nodes/update', methods=['POST'])
def update_nodes():
    try:
        source = request.json.get('source', 'blockchain')
        if source == 'blockchain':
            result = subprocess.run(
                ['./nodes-update', 'blockchain'],
                cwd=os.getcwd(),
                capture_output=True,
                text=True
            )
        else:
            result = subprocess.run(
                ['./nodes-update', 'node', source],
                cwd=os.getcwd(),
                capture_output=True,
                text=True
            )
        
        if result.returncode == 0:
            return jsonify({"status": "success", "message": "Node list updated"})
        else:
            return jsonify({"status": "error", "message": result.stderr}), 400
            
    except Exception as e:
        return jsonify({"error": str(e)}), 500

def parse_node_list(output):
    # Parse the text output into a list of node objects
    # This is a simplified example - adjust based on actual output format
    lines = output.strip().split('\n')
    nodes = []
    
    for line in lines:
        if line.strip():
            parts = line.split()
            if len(parts) >= 2:
                nodes.append({
                    "id": parts[0],
                    "name": parts[1],
                    "network": parts[2] if len(parts) > 2 else "unknown",
                    "status": "active"
                })
    
    return nodes

def ensure_executable(file_path):
    if os.path.exists(file_path):
        try:
            os.chmod(file_path, 0o755)
        except Exception as e:
            print(f"Warning: Could not set executable permissions for {file_path}: {e}")
    else:
        print(f"Warning: {file_path} not found. Some features may not work.")

if __name__ == '__main__':
    # Ensure we're in the correct directory
    script_dir = os.path.dirname(os.path.abspath(__file__))
    os.chdir(script_dir)
    
    # Make sure the tools are executable if they exist
    ensure_executable('./node')
    ensure_executable('./nodes-update')
    
    # Check for required tools
    if not os.path.exists('./node') or not os.path.exists('./nodes-update'):
        print("\nError: Required tools 'node' and/or 'nodes-update' not found in the web directory.")
        print("Please copy these executables from the PrivatenessTools directory and try again.\n")
    
    app.run(debug=True, port=5000)
