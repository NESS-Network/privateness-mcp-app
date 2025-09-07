# Privateness Node Manager Web UI

A web-based interface for managing Privateness network nodes.

## Prerequisites

- Python 3.8+
- [uv](https://github.com/astral-sh/uv) package manager
- PrivatenessTools installed and configured
- Node.js (for development, optional)

## Setup

1. First, install `uv` if you haven't already:

   ```bash
   curl -LsSf https://astral.sh/uv/install.sh | sh
   ```

2. Install Python dependencies using `uv`:

   ```bash
   uv pip install -r requirements.txt
   ```

3. Make sure the PrivatenessTools are in your PATH or in the web directory:
   - `node`
   - `nodes-update`
   - Other PrivatenessTools executables

4. Make the tools executable (if needed):

   ```bash
   chmod +x node nodes-update
   ```

## Running the Web Interface

1. Navigate to the web directory:

   ```bash
   cd /path/to/privateness-mcp-app/web
   ```

2. Start the Flask server:

   ```bash
   python app.py
   ```

3. Open your browser and navigate to:

   ```bash
   http://localhost:5000
   ```

## Features

- View list of nodes
- Update node list from blockchain
- View node details
- Basic node management

## Security Notes

- The web interface should only be run on trusted networks
- By default, the server listens on all interfaces (0.0.0.0)
- For production use, consider adding authentication and HTTPS

## Development

To modify the frontend:

1. Edit files in the `templates` directory
2. The interface uses Tailwind CSS for styling
3. No build step is required - changes are reflected on page refresh
