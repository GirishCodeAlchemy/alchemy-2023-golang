import http.server
import socketserver

PORT = 8000
HOST = "0.0.0.0"

class SimpleHTTPRequestHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hello, World!')

with socketserver.TCPServer((HOST, PORT), SimpleHTTPRequestHandler) as httpd:
    print("Server started on", HOST, "port", PORT)
    httpd.serve_forever()
